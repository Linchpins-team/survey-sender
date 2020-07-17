package main

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestSendMail(t *testing.T) {
	survey := Survey{
		Name: "Foo Bar",
		Email: "foobar@example.com",
		Content: "Hi!\nHow are you?",
	}
	survey.CreatedAt = time.Now()
	c, _ := LoadConfig("config.toml")
	sender := Sender{
		Config: c,
	}
	assert.Nil(t, sender.Send(survey))	
}