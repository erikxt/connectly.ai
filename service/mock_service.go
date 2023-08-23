package service

// mock get current user id
func GetCurrentUserID() (userID string) {
	return "9805893816149728"
}

// mock get right context by event type
// TODO: get context by type
func GetContextByEventType(eventType, msgText string) (context string) {
	if msgText != "" {
		return msgText
	}
	return "hello world"
}
