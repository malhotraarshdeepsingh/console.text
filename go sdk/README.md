# GoSDK CLI Tool - Error Listener

Use this in go sdk, Listens to error using an error channel, when trigerred sends error to recipent or to the one who deployed it


## Installation

just clone or copy as install as an dependency

## Usage

Example:
```go
import "path/to/go sdk"

func main() {
    errorChan := make(chan error)
    go_sdk.ListenForErrors(errorChan, go_sdk.SendEmailNotification)
    // ... your code ...
    // On error:
    errorChan <- fmt.Errorf("something went wrong")
}
```

## Email Notification Setup

You need to setup uop the below ENV vars

- `GOSDK_SMTP_HOST` - SMTP server host (e.g., smtp.gmail.com)
- `GOSDK_SMTP_PORT` - SMTP server port (e.g., 587)
- `GOSDK_SMTP_USER` - SMTP username
- `GOSDK_SMTP_PASS` - SMTP password
- `GOSDK_EMAIL_FROM` - Sender email address
- `GOSDK_EMAIL_TO` - Recipient email address