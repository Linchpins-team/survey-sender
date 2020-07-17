package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func testSurvey() Survey {
	return Survey{
		Name:      "Foo Bar",
		Email:     "foobar@example.com",
		Content:   "Hi!\nHow are you?",
		CreatedAt: time.Now(),
	}
}

func TestSendMail(t *testing.T) {
	sender := Sender{
		Config: Configure,
	}
	assert.Nil(t, sender.Send(testSurvey()))
}
