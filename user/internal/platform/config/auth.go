package config

import "os"
import _ "github.com/joho/godotenv/autoload"

var (
	PrivateKeyPath   = os.Getenv("PRIVATE_KEY_PATH")
	PublicKeyPath    = os.Getenv("PUBLIC_KEY_PATH")
	JwtExpireSeconds = os.Getenv("JWT_EXPIRE_SECONDS")
)
