package messenger

import (
	"chatbot/model"
	"chatbot/service"
	"encoding/json"
	"fmt"
	"os"

	"github.com/maciekmm/messenger-platform-go-sdk"
	"github.com/sirupsen/logrus"
)

var mess = &messenger.Messenger{
	AccessToken: os.Getenv("FB_PAGE_ACCESS_TOKEN"),
	VerifyToken: "eriktoken",
	Debug:       messenger.DebugAll,
}

// init messenger
func init() {
	mess.MessageReceived = messageReceived
}

func messageReceived(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
	var messageService = &service.MessageService{}
	attachments, err := json.Marshal(msg.Attachments)
	if err != nil {
		logrus.Error(fmt.Sprintf("marshal attachments error: %v", err))
		return
	}
	var message = &model.Message{
		MessageID:   msg.ID,
		ConsumerID:  opts.Sender.ID,
		SmbId:       opts.Recipient.ID,
		Content:     msg.Text,
		Attachments: string(attachments),
		Inbound:     true,
	}
	messageService.AddMessage(message)
	// check if consumer is in database
	var consumerService = &service.ConsumerService{}
	consumers, err := consumerService.GetConsumer(opts.Sender.ID)
	if err != nil {
		logrus.Error(fmt.Sprintf("get consumer error: %v", err))
		return
	}
	// check consumer is nil
	if consumers[0].ConsumerID == "" {
		// get user profile
		profile, err := GetProfile(opts.Sender.ID)
		if err != nil {
			logrus.Error(fmt.Sprintf("get profile error: %v", err))
			return
		}
		// new a consumer
		var consumerNew = &model.Consumer{
			ConsumerID:     opts.Sender.ID,
			FirstName:      profile.FirstName,
			LastName:       profile.LastName,
			Gender:         profile.Gender,
			ProfilePicture: profile.ProfilePicture,
			Locale:         profile.Locale,
			Timezone:       profile.Timezone,
		}
		// add consumer to database
		consumerService.AddConsumer(consumerNew)
	}
}

// get a mess instance
func GetInstance() *messenger.Messenger {
	return mess
}

// send simple message
func SendSimpleMessage(recipientID, text string) (string, error) {
	resp, err := mess.SendSimpleMessage(recipientID, text)
	return resp.MessageID, err
}

// get user profile
func GetProfile(userID string) (*messenger.Profile, error) {
	return mess.GetProfile(userID)
}
