package mail

import (
	"fmt"
	"github.com/lukasjarosch/educonn-platform/mail/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/user/proto"
	"strconv"
	gomail "github.com/go-mail/mail"
)

type SmtpMail struct {
	dialer   *gomail.Dialer
	hostname string
	port     int
	username string
	password string
}

// New creates a new mail sender
func NewSmtpMail(hostname string, port string, username string, password string) (mail *SmtpMail, err error) {
	portVal, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	dialer := gomail.NewDialer(hostname, portVal, username, password)
	return &SmtpMail{dialer: dialer}, nil
}

// SendUserCreated sends an UserCreated email aka. the welcome email
func (m *SmtpMail) SendUserCreated(user *educonn_user.UserDetails) (err error) {
	from := config.DefaultSenderAddress
	to := user.Email
	name := user.FirstName
	subject := fmt.Sprintf(config.UserCreatedSubject, name)
	body := "Ohai" // TODO: proper mail text

	if err = m.send(from, to, subject, body); err != nil {
		return err
	}
	return nil
}

// SendUserDeleted sends an UserDeleted email aka. goodbye email
func (m *SmtpMail) SendUserDeleted(user *educonn_user.UserDetails) (err error) {
	from := config.DefaultSenderAddress
	to := user.Email
	name := user.FirstName
	subject := fmt.Sprintf(config.UserDeletedSubject, name)
	body := "Goodbye" // TODO: proper mail text

	if err = m.send(from, to, subject, body); err != nil {
		return err
	}
	return nil
}

// send actually sends out the email
func (m *SmtpMail) send(from string, to string, subject string, body string) (err error) {
	mail := gomail.NewMessage()

	mail.SetHeader("From", from)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	if err := m.dialer.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}
