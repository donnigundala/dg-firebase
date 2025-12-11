# dg-firebase

Firebase integration for the DG Framework.

## Installation

```bash
go get github.com/donnigundala/dg-firebase
```

## Configuration

Add the following to your `config/firebase.yaml` or `config/app.yaml`:

### Option 1: Using a credentials file

```yaml
firebase:
  credentials_file: "path/to/service-account.json"
```

### Option 2: Using raw JSON (via environment variable)

```bash
export FIREBASE_CREDENTIALS='{"type": "service_account", "project_id": "...", ...}'
```

```yaml
firebase:
  credentials_json: "${FIREBASE_CREDENTIALS}"
```

## Usage

Register the provider in your application:

```go
import (
    "github.com/donnigundala/dg-firebase"
)

func main() {
    app := foundation.New(".")
    
    // Register Firebase provider
    app.Register(&firebase.FirebaseServiceProvider{})
    
    app.Boot()
}
```

Access the client:

```go
import (
    "github.com/donnigundala/dg-firebase"
    "github.com/donnigundala/dg-firebase/fcm"
)

client, err := app.Make("firebase")
if err != nil {
    log.Fatal(err)
}

fbClient := client.(*firebase.Client)

// Use Firestore
firestore, err := fbClient.Firestore(ctx)

// Use Auth
auth, err := fbClient.Auth(ctx)

// Use FCM (Firebase Cloud Messaging)
fcmClient, err := fbClient.FCM(ctx)
if err != nil {
    log.Fatal(err)
}

// Send a message using the fluent builder
message := fcm.NewMessage().
    Token("device-token").
    Notification("Hello", "World").
    Data(map[string]string{
        "key": "value",
    }).
    Build()

response, err := fcmClient.Send(ctx, message)

// Or send to a topic
topicMessage := fcm.NewMessage().
    Topic("news").
    Notification("Breaking News", "Something happened!").
    Build()

response, err = fcmClient.Send(ctx, topicMessage)
```
