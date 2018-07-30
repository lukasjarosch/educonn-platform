package config

import "os"
import _ "github.com/joho/godotenv/autoload"

var (
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPass = os.Getenv("DB_PASS")
	DbName = os.Getenv("DB_NAME")
)

const VideoCollectionName = "videos"
