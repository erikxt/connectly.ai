package model

import (
	"time"
)

type Message struct {
	ID          uint   `gorm:"primarykey"`
	MessageID   string `gorm:"index"`
	ConsumerID  string `gorm:"type:varchar(20);index"`
	SmbId       string `gorm:"type:varchar(20);index"`
	Content     string
	Attachments string
	Inbound     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// add message to database
func AddMessage(message *Message) error {
	return DB.Create(message).Error
}

// get all messages from database
func GetMessages() (messages []Message, err error) {
	err = DB.Find(&messages).Error
	return
}

// get messages from database by consumer id
func GetMessagesByConsumerID(consumerID string) (messages []Message, err error) {
	err = DB.Where("consumer_id = ?", consumerID).Find(&messages).Error
	return
}