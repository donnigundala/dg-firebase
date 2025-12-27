# dg-firebase

Firebase integration for the DG Framework.

## Installation

```bash
go get github.com/donnigundala/dg-firebase
```

## Configuration

The plugin uses the `firebase` key in your configuration file.

### Configuration Mapping (YAML vs ENV)

| YAML Key | Environment Variable | Default | Description |
| :--- | :--- | :--- | :--- |
| `firebase.credentials_file` | - | - | Path to service account JSON |
| `firebase.credentials_json` | `FIREBASE_CREDENTIALS` | - | Raw JSON string |

### Example YAML

```yaml
firebase:
  credentials_file: "storage/firebase-auth.json"
```

## Observability

This plugin is instrumented with OpenTelemetry metrics. If `dg-observability` is registered and enabled, the following metrics are automatically emitted by the FCM client:

*   `firebase.fcm.message.sent`: Counter (labels: `operation`, `type`, `status`)
*   `firebase.fcm.message.duration`: Histogram (labels: `operation`, `type`, `status`)

To enable observability, ensure the `dg-observability` plugin is registered and configured:

```yaml
observability:
  enabled: true
  service_name: "my-app"
```

## Usage

Register the provider in your application:

```go
package main

import (
    "github.com/donnigundala/dg-core/foundation"
    "github.com/donnigundala/dg-firebase"
)

func main() {
    app := foundation.New(".")
    
    // Register Firebase provider (uses 'firebase' key in config)
    app.Register(dgfirebase.NewFirebaseServiceProvider())
    
    app.Start()
}
```

### Integration via InfrastructureSuite
In your `bootstrap/app.go`, you typically use the declarative suite pattern:

```go
func InfrastructureSuite(workerMode bool) []foundation.ServiceProvider {
	return []foundation.ServiceProvider{
		dgfirebase.NewFirebaseServiceProvider(),
		// ... other providers
	}
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
