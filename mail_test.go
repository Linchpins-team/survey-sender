package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func testSurvey() Survey {
	survey := Survey{
		Name:    "Foo Bar",
		Email:   "foobar@example.com",
		Content: "Hi!\nHow are you?",
	}
	survey.CreatedAt = time.Now()
	return survey
}

func TestSendMail(t *testing.T) {
	sender := Sender{
		Config: Configure,
	}
	assert.Nil(t, sender.Send(testSurvey()))
}
