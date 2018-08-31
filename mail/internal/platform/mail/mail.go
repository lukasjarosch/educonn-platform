package mail

import (
	"bytes"
	"strconv"

	"net/smtp"

	"github.com/alecthomas/template"
	gomail "github.com/go-mail/mail"
	"context"
	"github.com/opentracing/opentracing-go"
)

//MailRequest struct
type MailRequest struct {
	from    string
	to      string
	subject string
	body    string
}

func NewRequest(to string, from string, subject, body string) *MailRequest {
	return &MailRequest{
		to:      to,
		from:    from,
		subject: subject,
		body:    body,
	}
}
func (r *MailRequest) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

type SmtpMail struct {
	dialer   *gomail.Dialer
	hostname string
	port     int
	auth     smtp.Auth
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

func (m *SmtpMail) SendEmail(ctx context.Context, request *MailRequest) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SmtpMail.SendMail")
	defer span.Finish()

	mail := gomail.NewMessage()

	mail.SetHeader("From", request.from)
	mail.SetHeader("To", request.to)
	mail.SetHeader("Subject", request.subject)
	mail.SetHeader("")
	mail.SetBody("text/html", request.body)

	if err := m.dialer.DialAndSend(mail); err != nil {
		return false, err
	}
	return true, nil
}

