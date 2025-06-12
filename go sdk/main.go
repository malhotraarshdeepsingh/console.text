package main

import (
	"fmt"
	"os"
)


func ListenForErrors(errorChan <-chan error, notifyFunc func(error)) {
	go func() {
		for err := range errorChan {
			fmt.Println("Error detected:", err)
			notifyFunc(err)
		}
	}()
}

func getEmailConfigFromEnv() EmailConfig {
	return EmailConfig{
		SMTPHost: os.Getenv("GOSDK_SMTP_HOST"),
		SMTPPort: os.Getenv("GOSDK_SMTP_PORT"),
		Username: os.Getenv("GOSDK_SMTP_USER"),
		Password: os.Getenv("GOSDK_SMTP_PASS"),
		From:     os.Getenv("GOSDK_EMAIL_FROM"),
		To:       os.Getenv("GOSDK_EMAIL_TO"),
	}
}

func sendEmailNotification(err error) {
	cfg := getEmailConfigFromEnv()
	subject := "[GoSDK] Error Detected in Deployment"
	body := fmt.Sprintf("An error occurred: %v", err)
	e := SendEmail(cfg, subject, body)
	if e != nil {
		fmt.Println("Failed to send email notification:", e)
	}
}

func SendEmailNotification(err error) {
	cfg := getEmailConfigFromEnv()
	subject := "[GoSDK] Error Detected in Deployment"
	body := fmt.Sprintf("An error occurred: %v", err)
	e := SendEmail(cfg, subject, body)
	if e != nil {
		fmt.Println("Failed to send email notification:", e)
	}
}
