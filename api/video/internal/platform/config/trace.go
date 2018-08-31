package config

import "os"
import _ "github.com/joho/godotenv/autoload"

var ZipkinCollectorUrl = os.Getenv("ZIPKIN_COLLECTOR_URL")
