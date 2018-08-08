package service

import (
	"context"
	"github.com/lukasjarosch/educonn-platform/mail/internal/platform/mail"
	pbMail "github.com/lukasjarosch/educonn-platform/mail/proto"
	pbUser "github.com/lukasjarosch/educonn-platform/user/proto"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	"github.com/rs/zerolog/log"
)

type mailService struct {
	userCreatedChan    chan *pbUser.UserCreatedEvent
	userDeletedChan    chan *pbUser.UserDeletedEvent
	videoProcessedChan chan *pbVideo.VideoProcessedEvent
	mail               *mail.SmtpMail
	userClient         pbUser.UserClient
}

func NewMailService(userCreatedChannel chan *pbUser.UserCreatedEvent,
	userDeletedChannel chan *pbUser.UserDeletedEvent,
	videoProcessedChannel chan *pbVideo.VideoProcessedEvent,
	mail *mail.SmtpMail, userClient pbUser.UserClient) pbMail.EmailHandler {

	svc := &mailService{
		userCreatedChan:    userCreatedChannel,
		userDeletedChan:    userDeletedChannel,
		videoProcessedChan: videoProcessedChannel,
		mail:               mail,
		userClient:         userClient,
	}
	go svc.awaitUserCreatedEvent()
	go svc.awaitUserDeletedEvent()
	go svc.awaitVideoProcessedEvent()
	return svc
}

func (m *mailService) Send(ctx context.Context, request *pbMail.EmailRequest, response *pbMail.Response) error {
	log.Info().Msg("rpc call Email.Send is not yet implemented")

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

func (m *mailService) awaitVideoProcessedEvent() {
	for videoProcessed := range m.videoProcessedChan {
		user, err := m.userClient.Get(context.Background(), &pbUser.UserDetails{Id: videoProcessed.UserId})
		if err != nil {
			log.Warn().Err(err).Str("user", videoProcessed.UserId).Msg("unable to fetch user")
			continue
		}

		err = m.mail.SendVideoProcessed(videoProcessed.Video, user.User)
		if err != nil {
			log.Warn().Err(err).Str("recipient", user.User.Email).Msg("unable to send mail")
			continue
		}
		log.Info().Str("recipient", user.User.Email).Msg("sent video-processed notification")
	}
}
