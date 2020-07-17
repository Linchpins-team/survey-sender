package main

import (
	"net/smtp"
	"fmt"
)

type Sender struct {
	Config
}

func (s Sender) Send(survey Survey) (err error) {
	auth := smtp.PlainAuth("", s.Mail, s.Password, s.Host)

	for _, receiver := range s.Receivers {
		msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: Survey from %s\r\nReply-To: %s\r\n\r\n%s\r\n", s.Mail, receiver, survey.Name, survey.Email, survey.String())
		if err = smtp.SendMail(fmt.Sprintf("%s:%d", s.Host, s.Port), auth, s.Mail, s.Receivers, []byte(msg)); err != nil {
			return
		}
	}
	return nil
}