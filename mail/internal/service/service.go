package service

import (
	"context"
	"github.com/lukasjarosch/educonn-platform/mail/internal/platform/config"
	"github.com/lukasjarosch/educonn-platform/mail/internal/platform/mail"
	pbMail "github.com/lukasjarosch/educonn-platform/mail/proto"
	pbUser "github.com/lukasjarosch/educonn-platform/user/proto"
	pbVideo "github.com/lukasjarosch/educonn-platform/video/proto"
	"github.com/rs/zerolog/log"
)

const (
	userCreatedTemplate    = "user-created.html"
	userDeletedTemplate    = "user-deleted.html"
	videoProcessedTemplate = "video-processed.html"
)

type mailService struct {
	userCreatedChan    chan *pbUser.UserCreatedEvent
	userDeletedChan    chan *pbUser.UserDeletedEvent
	videoProcessedChan chan *pbVideo.VideoProcessedEvent
	mail               *mail.SmtpMail
	userClient         pbUser.UserClient
	videoClient        pbVideo.VideoClient
}

func NewMailService(userCreatedChannel chan *pbUser.UserCreatedEvent,
	userDeletedChannel chan *pbUser.UserDeletedEvent,
	videoProcessedChannel chan *pbVideo.VideoProcessedEvent,
	mail *mail.SmtpMail,
	userClient pbUser.UserClient,
	videoClient pbVideo.VideoClient) pbMail.EmailHandler {

	svc := &mailService{
		userCreatedChan:    userCreatedChannel,
		userDeletedChan:    userDeletedChannel,
		videoProcessedChan: videoProcessedChannel,
		mail:               mail,
		userClient:         userClient,
		videoClient:        videoClient,
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

		templateData := struct {
			Name string
			URL  string
		}{
			Name: userCreated.User.FirstName,
			URL:  "https://educonn.de/user/validate_email?token=asdfasdfkl23klön234ölkn",
		}

		// parse template
		mailRequest := mail.NewRequest(userCreated.User.Email, "notify@educonn.de", "Welcome to EduConn!", "awesome")
		err := mailRequest.ParseTemplate(config.TemplatePath+"/"+userCreatedTemplate, templateData)
		if err != nil {
			log.Info().Err(err).Msg("ParseTemplate failed")
		}

		// send mail
		_, err = m.mail.SendEmail(mailRequest)
		if err != nil {
			log.Warn().Err(err).Msg("unable to send email")
			continue
		}

		log.Info().Str("user", userCreated.User.Id).Msg("sent welcome email")
	}
}

func (m *mailService) awaitUserDeletedEvent() {
	for userDeleted := range m.userDeletedChan {
		templateData := struct {
			Name string
			URL  string
		}{
			Name: userDeleted.User.FirstName,
			URL:  "https://educonn.de/user/validate_email?token=asdfasdfkl23klön234ölkn",
		}

		// parse template
		mailRequest := mail.NewRequest(userDeleted.User.Email, "notify@educonn.de", "Farewell,"+userDeleted.User.FirstName, "sorry to see you leave")
		err := mailRequest.ParseTemplate(config.TemplatePath+"/"+userDeletedTemplate, templateData)
		if err != nil {
			log.Info().Err(err).Msg("ParseTemplate failed")
		}

		// send mail
		_, err = m.mail.SendEmail(mailRequest)
		if err != nil {
			log.Warn().Err(err).Msg("unable to send email")
			continue
		}

		log.Info().Str("user", userDeleted.User.Id).Msg("sent farewell email")
	}
}

func (m *mailService) awaitVideoProcessedEvent() {
	for videoProcessed := range m.videoProcessedChan {

		// fetch user
		user, err := m.userClient.Get(context.Background(), &pbUser.UserDetails{Id: videoProcessed.UserId})
		if err != nil {
			log.Warn().Err(err).Str("user", videoProcessed.UserId).Msg("unable to fetch user")
			continue
		}

		// fetch video url
		video, err := m.videoClient.GetById(context.Background(), &pbVideo.GetVideoRequest{Id:videoProcessed.Video.Id})
		if err != nil {
			log.Warn().Err(err).Str("video", videoProcessed.Video.Id).Msg("unable to fetch video")
			continue
		}

		templateData := struct {
			Name       string
			URL        string
			VideoTitle string
		}{
			Name:       user.User.FirstName,
			URL:		video.SignedUrl,
			VideoTitle: videoProcessed.Video.Title,
		}

		// parse template
		mailRequest := mail.NewRequest(user.User.Email, "notify@educonn.de", "Your video is ready to watch", "video ready")
		err = mailRequest.ParseTemplate(config.TemplatePath+"/"+videoProcessedTemplate, templateData)
		if err != nil {
			log.Info().Err(err).Msg("ParseTemplate failed")
		}

		// send mail
		_, err = m.mail.SendEmail(mailRequest)
		if err != nil {
			log.Warn().Err(err).Msg("unable to send email")
			continue
		}

		log.Debug().Str("user", user.User.Id).Msg("user notified that the video is processed")
	}
}
