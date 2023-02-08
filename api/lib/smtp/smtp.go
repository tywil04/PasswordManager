package smtp

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

var host string
var port string
var from string
var username string
var password string
var auth smtp.Auth

func SendTemplate(to string, subject string, toExecute *template.Template, data any) {
	var writer bytes.Buffer
	toExecute.Execute(&writer, data)
	SendHTML(to, subject, writer.String())
}

func SendHTML(to string, subject string, message string) {
	messagePart := fmt.Sprintf("From: %s\r\n", from)
	messagePart += fmt.Sprintf("To: %s\r\n", to)
	messagePart += fmt.Sprintf("Subject: %s\r\n", subject)
	messagePart += "MIME-version: 1.0;\r\n"
	messagePart += "Content-Type: text/html; charset=\"UTF-8\";\r\n"
	messagePart += fmt.Sprintf("\r\n%s\r\n", message)
	Send(to, []byte(messagePart))
}

func SendString(to string, subject string, message string) {
	messagePart := fmt.Sprintf("From: %s\r\n", from)
	messagePart += fmt.Sprintf("To: %s\r\n", to)
	messagePart += fmt.Sprintf("Subject: %s\r\n", subject)
	messagePart += fmt.Sprintf("\r\n%s\r\n", message)
	Send(to, []byte(messagePart))
}

func Send(to string, message []byte) {
	sendList := []string{to}
	smtp.SendMail(host+":"+port, auth, from, sendList, message)
}

func Connect() {
	host = os.Getenv("SMTP_HOST")
	port = os.Getenv("SMTP_PORT")
	from = os.Getenv("SMTP_FROM")
	username = os.Getenv("SMTP_USERNAME")
	password = os.Getenv("SMTP_PASSWORD")

	auth = smtp.PlainAuth("", username, password, host)

	LoadTemplates()
}
