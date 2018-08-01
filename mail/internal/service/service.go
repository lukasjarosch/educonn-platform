package service

import (
	"context"
	"github.com/lukasjarosch/educonn-platform/mail/internal/platform/mail"
	pbMail "github.com/lukasjarosch/educonn-platform/mail/proto"
	pbUser "github.com/lukasjarosch/educonn-platform/user/proto"
	"github.com/rs/zerolog/log"
)

type mailService struct {
	userCreatedChan chan *pbUser.UserCreatedEvent
	userDeletedChan chan *pbUser.UserDeletedEvent
	mail            *mail.SmtpMail
}

type mailRepository interface {
}

func NewMailService(userCreatedChannel chan *pbUser.UserCreatedEvent,
	userDeletedChannel chan *pbUser.UserDeletedEvent, mail *mail.SmtpMail) pbMail.EmailHandler {

	svc := &mailService{userCreatedChan: userCreatedChannel, userDeletedChan: userDeletedChannel, mail: mail}
	go svc.awaitUserCreatedEvent()
	go svc.awaitUserDeletedEvent()
	return svc
}

func (m *mailService) Send(ctx context.Context, request *pbMail.EmailRequest, response *pbMail.Response) error {
	return nil
}

// awaitUserCreatedEvent is waiting for UserCreatedEvents on the userCreatedChan channel
// The method is running as go routine, called by NewMailService and handles the sending of the emails
func (m *mailService) awaitUserCreatedEvent() {
	for userCreated := range m.userCreatedChan {
		err := m.mail.SendUserCreated(userCreated.User)
		if err != nil {
			log.Warn().Interface("error", err).Str("recipient", userCreated.User.Email).Msg("unable to send mail")
			continue
		}
		log.Info().Str("recipient", userCreated.User.Email).Msg("sent welcome email")
	}
}

func (m *mailService) awaitUserDeletedEvent() {
	for userDeleted := range m.userDeletedChan {
		err := m.mail.SendUserDeleted(userDeleted.User)
		if err != nil {
			log.Warn().Interface("error", err).Str("recipient", userDeleted.User.Email).Msg("unable to send mail")
			continue
		}
		log.Info().Str("recipient", userDeleted.User.Email).Msg("sent farewell email")
	}
}
