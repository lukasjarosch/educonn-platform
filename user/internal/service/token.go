package service

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	pb "github.com/lukasjarosch/educonn-platform/user/proto"
	"io/ioutil"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/errors"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/config"
	"strconv"
)

type CustomClaims struct {
	User *pb.UserDetails
	jwt.StandardClaims
}

type Auth interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (token string, err error)
}

type TokenService struct {
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
}

func NewTokenService(publicKeyPath string, privateKeyPath string) (*TokenService, error) {
	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
	    return nil, errors.PrivateKeyFileNotFound
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
	    return nil, err
	}

	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
	    return nil, errors.PublicKeyFileNotFound
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
	    return nil, err
	}

	return &TokenService{
		signKey: signKey,
		verifyKey: verifyKey,
	}, nil
}

func (t *TokenService) Encode(user *pb.UserDetails) (token string, err error) {
	expires, err := strconv.ParseInt(config.JwtExpireSeconds, 10, 64)
	if err != nil {
	    return "", err
	}

	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expires,
			Issuer: config.ServiceName,
		},
	}

	rsaSigner := jwt.New(jwt.GetSigningMethod("RS256"))
	rsaSigner.Claims = claims

	token, err = rsaSigner.SignedString(t.signKey)
	if err != nil {
	    return "", err
	}
	return token, nil
}
