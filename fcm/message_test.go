package fcm

import (
	"testing"

	"firebase.google.com/go/v4/messaging"
	"github.com/stretchr/testify/assert"
)

func TestNewMessage(t *testing.T) {
	builder := NewMessage()

	assert.NotNil(t, builder)
	assert.NotNil(t, builder.message)
}

func TestMessageBuilder_Token(t *testing.T) {
	token := "test-device-token"
	message := NewMessage().
		Token(token).
		Build()

	assert.Equal(t, token, message.Token)
}

func TestMessageBuilder_Topic(t *testing.T) {
	topic := "news"
	message := NewMessage().
		Topic(topic).
		Build()

	assert.Equal(t, topic, message.Topic)
}

func TestMessageBuilder_Condition(t *testing.T) {
	condition := "'stock-GOOG' in topics || 'industry-tech' in topics"
	message := NewMessage().
		Condition(condition).
		Build()

	assert.Equal(t, condition, message.Condition)
}

func TestMessageBuilder_Notification(t *testing.T) {
	title := "Test Title"
	body := "Test Body"

	message := NewMessage().
		Notification(title, body).
		Build()

	assert.NotNil(t, message.Notification)
	assert.Equal(t, title, message.Notification.Title)
	assert.Equal(t, body, message.Notification.Body)
}

func TestMessageBuilder_NotificationWithImage(t *testing.T) {
	title := "Test Title"
	body := "Test Body"
	imageURL := "https://example.com/image.png"

	message := NewMessage().
		NotificationWithImage(title, body, imageURL).
		Build()

	assert.NotNil(t, message.Notification)
	assert.Equal(t, title, message.Notification.Title)
	assert.Equal(t, body, message.Notification.Body)
	assert.Equal(t, imageURL, message.Notification.ImageURL)
}

func TestMessageBuilder_Data(t *testing.T) {
	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	message := NewMessage().
		Data(data).
		Build()

	assert.Equal(t, data, message.Data)
}

func TestMessageBuilder_Android(t *testing.T) {
	androidConfig := &messaging.AndroidConfig{
		Priority: "high",
	}

	message := NewMessage().
		Android(androidConfig).
		Build()

	assert.Equal(t, androidConfig, message.Android)
}

func TestMessageBuilder_APNS(t *testing.T) {
	apnsConfig := &messaging.APNSConfig{
		Headers: map[string]string{
			"apns-priority": "10",
		},
	}

	message := NewMessage().
		APNS(apnsConfig).
		Build()

	assert.Equal(t, apnsConfig, message.APNS)
}

func TestMessageBuilder_Webpush(t *testing.T) {
	webpushConfig := &messaging.WebpushConfig{
		Headers: map[string]string{
			"TTL": "3600",
		},
	}

	message := NewMessage().
		Webpush(webpushConfig).
		Build()

	assert.Equal(t, webpushConfig, message.Webpush)
}

func TestMessageBuilder_Chaining(t *testing.T) {
	// Test that methods can be chained together
	message := NewMessage().
		Token("device-token").
		Notification("Title", "Body").
		Data(map[string]string{"key": "value"}).
		Build()

	assert.Equal(t, "device-token", message.Token)
	assert.NotNil(t, message.Notification)
	assert.Equal(t, "Title", message.Notification.Title)
	assert.Equal(t, "Body", message.Notification.Body)
	assert.Equal(t, "value", message.Data["key"])
}

func TestMessageBuilder_CompleteMessage(t *testing.T) {
	// Test building a complete message with all fields
	message := NewMessage().
		Token("test-token").
		Notification("Breaking News", "Something important happened").
		Data(map[string]string{
			"article_id": "123",
			"category":   "news",
		}).
		Android(&messaging.AndroidConfig{
			Priority: "high",
		}).
		Build()

	assert.Equal(t, "test-token", message.Token)
	assert.NotNil(t, message.Notification)
	assert.Equal(t, "Breaking News", message.Notification.Title)
	assert.Equal(t, 2, len(message.Data))
	assert.NotNil(t, message.Android)
	assert.Equal(t, "high", message.Android.Priority)
}
