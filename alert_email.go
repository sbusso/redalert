package main

import (
	"net/smtp"
	"strings"
)

type Email struct {
	user                  string
	pass                  string
	notificationAddresses []string
}

func (a Email) Trigger(event *Event) error {

	body := "To: " + strings.Join(a.notificationAddresses, ",") +
		"\r\nSubject: " + event.ShortMessage() +
		"\r\n\r\n" + event.ShortMessage()

	auth := smtp.PlainAuth("", a.user, a.pass, "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, a.user,
		a.notificationAddresses, []byte(body))
	if err != nil {
		return err
	}

	event.server.log.Println(white, "Email alert successfully triggered.", reset)
	return nil
}
