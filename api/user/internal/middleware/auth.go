package middleware

import (
	"github.com/micro/go-micro/client"
	"context"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/metadata"
	"strings"
	"encoding/base64"
	"github.com/lukasjarosch/educonn-platform/user/pkg/jwt_handler"
	"github.com/micro/go-micro/errors"
	"github.com/lukasjarosch/educonn-platform/api/user/internal/platform/config"
)

type authWrapper struct {
	client.Client
}

func (a *authWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption)	error {
	// get metadata from context or create new
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = metadata.Metadata{}
	}
	token, err := getBearerToken(md)
	if err != nil {
		return a.Client.Call(ctx, req, rsp)
	}

	claims, err := decodeToken(token)
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return errors.Unauthorized(config.ServiceName, "%s", jwt_handler.TokenExpired)
		}
		return a.Client.Call(ctx, req, rsp)
	}
	ctx = addUserId(ctx, claims)
	return a.Client.Call(ctx, req, rsp)
}


// Implements the server.HandlerWrapper
func AuthHandlerWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		// get metadata from context or create new
		md, ok := metadata.FromContext(ctx)
		if !ok {
			md = metadata.Metadata{}
		}
		token, err := getBearerToken(md)
		if err != nil {
			return fn(ctx, req, rsp)
		}

		claims, err := decodeToken(token)
		if err != nil {
			if strings.Contains(err.Error(), "token is expired") {
				return errors.Unauthorized(config.ServiceName, "%s", jwt_handler.TokenExpired)
			}
			return fn(ctx, req, rsp)
		}

		ctx = addUserId(ctx, claims)

		return fn(ctx, req, rsp)
	}
}

func decodeToken(token string) (*jwt_handler.CustomClaims, error) {
	jwtHandler, err := jwt_handler.NewJwtTokenHandler(config.PublicKeyPath, "")
	if err != nil {
		return nil, err
	}
	claims, err := jwtHandler.Decode(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func addUserId(ctx context.Context, claims *jwt_handler.CustomClaims) context.Context {
	// get metadata from context or create new
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = metadata.Metadata{}
	}

	if _, ok := md["X-User-Id"]; !ok {
		md["X-User-Id"] = claims.User.Id
		ctx = metadata.NewContext(ctx, md)
	}
	return ctx
}
// Extract the token from the Metadata map
func getBearerToken(md metadata.Metadata) (string, error) {
	authHeader := md["Authorization"]
	if authHeader == "" {
		return "", jwt_handler.AuthenticationHeaderMissing
	}

	// Confirm the request is sending Basic Authentication credentials.
	if !strings.HasPrefix(authHeader, jwt_handler.BasicSchema) && !strings.HasPrefix(authHeader, jwt_handler.BearerSchema) {
		return "", jwt_handler.InvalidAuthenticationScheme
	}

	// Get the token from the request header
	// The first six characters are skipped - e.g. "Basic ".
	if strings.HasPrefix(authHeader, jwt_handler.BasicSchema) {
		str, err := base64.StdEncoding.DecodeString(authHeader[len(jwt_handler.BasicSchema):])
		if err != nil {
			return "", jwt_handler.Base64EncodingError
		}
		creds := strings.Split(string(str), ":")
		return creds[0], nil
	}

	return authHeader[len(jwt_handler.BearerSchema):], nil
}
