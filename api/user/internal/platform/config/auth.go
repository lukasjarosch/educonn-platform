package config

import "os"
import _ "github.com/joho/godotenv/autoload"

var (
	PublicKeyPath    = os.Getenv("PUBLIC_KEY_PATH")
)
