package config

import "os"
import _ "github.com/joho/godotenv/autoload"

var (
	AuthPublicKeyPath    = os.Getenv("AUTH_PUBLIC_KEY_PATH")
)

