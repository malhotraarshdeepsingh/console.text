package main

import (
	"net/smtp"
)

type EmailConfig struct {
	SMTPHost string
	SMTPPort string
	Username string
	Password string
	From     string
	To       string
}

func SendEmail(cfg EmailConfig, subject, body string) error {
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.SMTPHost)
	msg := []byte("To: " + cfg.To + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")
	addr := cfg.SMTPHost + ":" + cfg.SMTPPort
	return smtp.SendMail(addr, auth, cfg.From, []string{cfg.To}, msg)
}


