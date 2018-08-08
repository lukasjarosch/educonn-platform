package config

import "os"
import _ "github.com/joho/godotenv/autoload"

const DefaultSenderAddress = "service@educonn.de"

var (
	SmtpHostname = os.Getenv("SMTP_HOST")
	SmtpPort     = os.Getenv("SMTP_PORT")
	SmtpUsername = os.Getenv("SMTP_USER")
	SmtpPassword = os.Getenv("SMTP_PASS")
)

var (
	TemplatePath = os.Getenv("MAIL_TEMPLATE_PATH")
	UserCreatedSubject = "Welcome to EduConn %s"
	UserDeletedSubject = "We're sorry to see you go %s"
)
