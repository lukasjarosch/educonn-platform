package jwt_handler

import (
	"github.com/micro/go-micro/metadata"
	"strings"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"crypto/rsa"
	"io/ioutil"
	"strconv"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/config"
	"time"
	"github.com/lukasjarosch/educonn-platform/user/proto"
)

const (
	BasicSchema  = "Basic "
	BearerSchema = "Bearer "
)

type CustomClaims struct {
	User *proto.UserDetails
	jwt.StandardClaims
}


type JwtTokenHandler struct {
	verifyKey *rsa.PublicKey
	signKey *rsa.PrivateKey
}


// NewJwtTokenHandler creates a new TokenHandler. It is possible to only pass the publicKeyPath and set the privateKeyPath
// to an empty string. This way the class can also be used just for decoding in other services.
func NewJwtTokenHandler(publicKeyPath string, privateKeyPath string) (*JwtTokenHandler, error) {
	signKey := new(rsa.PrivateKey)
	signKey = nil

	// optional private key for signing
	if privateKeyPath != "" {
		signBytes, err := ioutil.ReadFile(privateKeyPath)
		if err != nil {
			return nil, TokenPrivateKeyFileNotFound
		}

		signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
		if err != nil {
			return nil, err
		}
	}

	// public key
	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, TokenPublicKeyFileNotFound
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}

	return &JwtTokenHandler{
		verifyKey: verifyKey,
		signKey: signKey,
	}, nil
}

// Decode validates a given token using the public key and returns the CustomClaims struct if possible
func (j *JwtTokenHandler) Decode(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.verifyKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// Extract the token from the Metadata map
func (j *JwtTokenHandler) GetBearerToken(md metadata.Metadata) (string, error) {
	authHeader := md["Authorization"]
	if authHeader == "" {
		return "", AuthenticationHeaderMissing
	}

	// Confirm the request is sending Basic Authentication credentials.
	if !strings.HasPrefix(authHeader, BasicSchema) && !strings.HasPrefix(authHeader, BearerSchema) {
		return "", InvalidAuthenticationScheme
	}

	// Get the token from the request header
	// The first six characters are skipped - e.g. "Basic ".
	if strings.HasPrefix(authHeader, BasicSchema) {
		str, err := base64.StdEncoding.DecodeString(authHeader[len(BasicSchema):])
		if err != nil {
			return "", Base64EncodingError
		}
		creds := strings.Split(string(str), ":")
		return creds[0], nil
	}

	return authHeader[len(BearerSchema):], nil
}

// Encode ecreates a new JWT token using the customclaims
func (j *JwtTokenHandler) Encode(user *proto.UserDetails) (token string, err error) {

	if j.signKey == nil {
		return "", PrivateKeyMissing
	}

	expires, err := strconv.Atoi(config.JwtExpireSeconds)
	if err != nil {
		return "", err
	}

	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expires)).Unix(),
			Issuer: config.ServiceName,
		},
	}

	rsaSigner := jwt.New(jwt.GetSigningMethod("RS256"))
	rsaSigner.Claims = claims

	token, err = rsaSigner.SignedString(j.signKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
