package service

import (
	"context"
	"github.com/lukasjarosch/educonn-master-thesis/mail/internal/platform/mail"
	"github.com/lukasjarosch/educonn-master-thesis/mail/proto"
	"github.com/lukasjarosch/educonn-master-thesis/user/proto"
	log "github.com/sirupsen/logrus"
)

type mailService struct {
	userCreatedChan chan *educonn_user.UserCreatedEvent
	userDeletedChan chan *educonn_user.UserDeletedEvent
	mail            *mail.SmtpMail
}

type mailRepository interface {
}

func NewMailService(userCreatedChannel chan *educonn_user.UserCreatedEvent,
	userDeletedChannel chan *educonn_user.UserDeletedEvent, mail *mail.SmtpMail) educonn_mail.EmailHandler {

	svc := &mailService{userCreatedChan: userCreatedChannel, userDeletedChan: userDeletedChannel, mail: mail}
	go svc.awaitUserCreatedEvent()
	go svc.awaitUserDeletedEvent()
	return svc
}

func (m *mailService) Send(ctx context.Context, request *educonn_mail.EmailRequest, response *educonn_mail.Response) error {
	return nil
}

// awaitUserCreatedEvent is waiting for UserCreatedEvents on the userCreatedChan channel
// The method is running as go routine, called by NewMailService and handles the sending of the emails
func (m *mailService) awaitUserCreatedEvent() {
	for userCreated := range m.userCreatedChan {
		err := m.mail.SendUserCreated(userCreated.User)
		if err != nil {
			log.Infof("Unable to send email: %v", err)
			continue
		}
		log.Infof("Sent welcome email to %s", userCreated.User.Email)
	}
}

func (m *mailService) awaitUserDeletedEvent() {
	for userDeleted := range m.userDeletedChan {
		err := m.mail.SendUserDeleted(userDeleted.User)
		if err != nil {
			log.Infof("Unable to send email: %v", err)
			continue
		}
		log.Infof("Sent goodbye email to: %v", userDeleted.User)
	}
}
