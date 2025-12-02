# FCM (Firebase Cloud Messaging)

This package provides a clean wrapper around Firebase Cloud Messaging with a fluent message builder API.

## Features

- **Fluent Message Builder**: Build FCM messages with a clean, chainable API
- **Send Methods**: Send to single devices, multiple devices, or topics
- **Topic Management**: Subscribe/unsubscribe tokens to topics
- **Full FCM Access**: Access underlying messaging client for advanced usage

## Usage

### Basic Message

```go
import "github.com/donnigundala/dg-framework/dg-firebase/fcm"

// Get FCM client
fcmClient, err := fbClient.FCM(ctx)

// Build and send a message
message := fcm.NewMessage().
    Token("device-token").
    Notification("Hello", "World").
    Build()

response, err := fcmClient.Send(ctx, message)
```

### Send to Topic

```go
message := fcm.NewMessage().
    Topic("news").
    Notification("Breaking News", "Something happened!").
    Data(map[string]string{
        "article_id": "123",
    }).
    Build()

response, err := fcmClient.Send(ctx, message)
```

### Send to Multiple Devices

```go
tokens := []string{"token1", "token2", "token3"}

multicast := &messaging.MulticastMessage{
    Tokens: tokens,
    Notification: &messaging.Notification{
        Title: "Hello",
        Body:  "World",
    },
}

batchResponse, err := fcmClient.SendMulticast(ctx, multicast)
```

### Topic Subscription

```go
tokens := []string{"token1", "token2"}

// Subscribe to topic
response, err := fcmClient.SubscribeToTopic(ctx, tokens, "news")

// Unsubscribe from topic
response, err := fcmClient.UnsubscribeFromTopic(ctx, tokens, "news")
```

### Advanced Usage

```go
// Access underlying messaging client
messagingClient := fcmClient.Client()
```

## Message Builder API

The fluent builder supports:

- `Token(token)` - Send to a specific device
- `Topic(topic)` - Send to a topic
- `Condition(condition)` - Send based on condition
- `Notification(title, body)` - Add notification
- `NotificationWithImage(title, body, imageURL)` - Add notification with image
- `Data(map[string]string)` - Add custom data
- `Android(config)` - Android-specific config
- `APNS(config)` - iOS-specific config
- `Webpush(config)` - Web push config
- `Build()` - Build the final message
