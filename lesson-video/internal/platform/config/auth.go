package config

import "os"

var (
	PublicKeyPath    = os.Getenv("AUTH_PUBLIC_KEY_PATH")
)
