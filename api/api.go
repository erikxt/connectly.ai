package api

import (
	"chatbot/messenger"
	"chatbot/model"
	"chatbot/service"

	"github.com/gin-gonic/gin"
)

// get consumer by id
func GetConsumer(c *gin.Context) {
	service := service.ConsumerService{}
	consumerID := c.Query("consumer_id")

	consumer, err := service.GetConsumer(consumerID)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, consumer)
}

// get all messages
func GetMessages(c *gin.Context) {
	service := service.MessageService{}
	consumerID := c.Query("consumer_id")

	messages, err := service.GetMessages(consumerID)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, messages)
}

// simplify the template message into send text message
func SendMessage(c *gin.Context) {
	// get parameters from post form
	consumerID := c.PostForm("consumer_id")
	eventType := c.PostForm("event_type")
	msgText := c.PostForm("msg_text")

	messageService := service.MessageService{}

	content := service.GetContextByEventType(eventType, msgText)
	messageID, _ := messenger.SendSimpleMessage(consumerID, content)
	err := messageService.AddMessage(&model.Message{
		ConsumerID: consumerID,
		SmbId:      service.GetCurrentUserID(),
		MessageID:  messageID,
		Content:    content,
		Inbound:    false,
	})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "ok")
}
