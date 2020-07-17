package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	survey := testSurvey()
	assert.Nil(t, Save(survey))
}
