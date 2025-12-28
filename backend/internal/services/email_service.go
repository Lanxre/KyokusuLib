package service

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

const (
	EMAIL_PORT = 587
	EMAIL_HOST = "smtp.gmail.com"
)

type EmailRequest struct {
	To      string
	Subject string
	Body    string
}

type EmailService struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewEmailService(name, password string) *EmailService {
	return &EmailService{
		Host:     EMAIL_HOST,
		Port:     EMAIL_PORT,
		Username: name,
		Password: password,
	}
}

func (s *EmailService) NewEmailRequest(sendTo, link string) *EmailRequest {
	return &EmailRequest{
		To:      sendTo,
		Subject: "Confirm Registration on KyokusuLib",
		Body:    s.buildTemplate("Welcome to KyokusuLib!", "Thank you for signing up. To complete your registration, please confirm your email address:", link, "Confirm Registration"),
	}
}

func (s *EmailService) NewResetPasswordEmailRequest(sendTo, link string) *EmailRequest {
	return &EmailRequest{
		To:      sendTo,
		Subject: "Password Reset Request",
		Body:    s.buildTemplate("Reset Your Password", "We received a request to reset your password. Click the button below to proceed:", link, "Reset Password"),
	}
}

func (s *EmailService) SendEmail(req EmailRequest) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.Username)
	m.SetHeader("To", req.To)
	m.SetHeader("Subject", req.Subject)
	m.SetBody("text/html", req.Body)

	d := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password)
	return d.DialAndSend(m)
}

func (s *EmailService) buildTemplate(title, text, link, buttonText string) string {
	return fmt.Sprintf(`
		<div style="font-family: sans-serif; max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #eee; border-radius: 8px;">
			<h2>%s</h2>
			<p style="color: #555;">%s</p>
			<div style="text-align: center; margin: 30px 0;">
				<a href="%s" style="background-color: #007BFF; color: #fff; padding: 12px 25px; text-decoration: none; border-radius: 4px; font-weight: bold; display: inline-block;">%s</a>
			</div>
			<p style="font-size: 12px; color: #999;">If the button doesn't work, copy this link: <br>%[3]s</p>
		</div>
	`, title, text, link, buttonText)
}