package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	ZipkinConnectionString = os.Getenv("ZIPKIN_COLLECTOR_URL")
)
