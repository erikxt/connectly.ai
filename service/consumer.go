package service

import (
	"chatbot/model"
)

type ConsumerService struct{}

// get consumer from database
func (s ConsumerService) GetConsumer(ID string) (consumers []model.Consumer, err error) {
	if ID == "" {
		return model.GetConsumers()
	}
	consumer, _ := model.GetConsumerByID(ID)
	// trans consumer to []model.Consumer
	consumers = append(consumers, consumer)
	return
}

// add consumer to database
func (s ConsumerService) AddConsumer(consumer *model.Consumer) error {
	return model.AddConsumer(consumer)
}
