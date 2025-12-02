package fcm

import "firebase.google.com/go/v4/messaging"

// MessageBuilder helps build FCM messages fluently.
type MessageBuilder struct {
	message *messaging.Message
}

// NewMessage creates a new message builder.
func NewMessage() *MessageBuilder {
	return &MessageBuilder{
		message: &messaging.Message{},
	}
}

// Token sets the device token.
func (b *MessageBuilder) Token(token string) *MessageBuilder {
	b.message.Token = token
	return b
}

// Topic sets the topic.
func (b *MessageBuilder) Topic(topic string) *MessageBuilder {
	b.message.Topic = topic
	return b
}

// Condition sets the condition.
func (b *MessageBuilder) Condition(condition string) *MessageBuilder {
	b.message.Condition = condition
	return b
}

// Notification sets the notification payload.
func (b *MessageBuilder) Notification(title, body string) *MessageBuilder {
	b.message.Notification = &messaging.Notification{
		Title: title,
		Body:  body,
	}
	return b
}

// NotificationWithImage sets the notification with an image.
func (b *MessageBuilder) NotificationWithImage(title, body, imageURL string) *MessageBuilder {
	b.message.Notification = &messaging.Notification{
		Title:    title,
		Body:     body,
		ImageURL: imageURL,
	}
	return b
}

// Data sets custom data payload.
func (b *MessageBuilder) Data(data map[string]string) *MessageBuilder {
	b.message.Data = data
	return b
}

// Android sets Android-specific options.
func (b *MessageBuilder) Android(config *messaging.AndroidConfig) *MessageBuilder {
	b.message.Android = config
	return b
}

// APNS sets Apple Push Notification Service options.
func (b *MessageBuilder) APNS(config *messaging.APNSConfig) *MessageBuilder {
	b.message.APNS = config
	return b
}

// Webpush sets Web Push options.
func (b *MessageBuilder) Webpush(config *messaging.WebpushConfig) *MessageBuilder {
	b.message.Webpush = config
	return b
}

// Build returns the constructed message.
func (b *MessageBuilder) Build() *messaging.Message {
	return b.message
}
