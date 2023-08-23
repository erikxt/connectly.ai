package model

import "time"

type Consumer struct {
	ConsumerID     string `gorm:"primarykey"`
	FirstName      string
	LastName       string
	Gender         string
	ProfilePicture string
	Locale         string
	Timezone       float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (Consumer) TableName() string {
	return "consumers"
}

// get all consumers from database
func GetConsumers() (consumers []Consumer, err error) {
	err = DB.Find(&consumers).Error
	return
}

// get consumer from database
func GetConsumerByID(ID string) (consumer Consumer, err error) {
	err = DB.First(&consumer, "consumer_id = ?", ID).Error
	return
}

// add consumer to database
func AddConsumer(consumer *Consumer) error {
	return DB.Create(consumer).Error
}
