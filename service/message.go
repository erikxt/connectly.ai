package service

import (
	"chatbot/model"
)

type MessageService struct{}

// if consumerId is empty, get all messages, otherwise get messages by consumer id
func (messageService *MessageService) GetMessages(consumerId string) (messages []model.Message, err error) {
	if consumerId != "" {
		return model.GetMessagesByConsumerID(consumerId)
	}
	return model.GetMessages()
}

// add message to database
func (messageService *MessageService) AddMessage(msg *model.Message) (err error) {
	return model.AddMessage(msg)
}
