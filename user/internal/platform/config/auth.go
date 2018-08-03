package config

import "os"

var (
	PrivateKeyPath   = os.Getenv("PRIVATE_KEY_PATH")
	PublicKeyPath    = os.Getenv("PUBLIC_KEY_PATH")
	JwtExpireSeconds = os.Getenv("JWT_EXPIRE_SECONDS")
)
