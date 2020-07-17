package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSurvey(t *testing.T) {
	s := testSurvey()
	response := testRequest(s)
	if response.StatusCode != 200 {
		t.Error(ioutil.ReadAll(response.Body))
	}
}

func TestInvalidEmail(t *testing.T) {
	s := Survey{
		Name:    "Foo Bar",
		Email:   "invalid.email",
		Content: "Nothing",
	}
	response := testRequest(s)
	assert.Equal(t, response.StatusCode, 403)
}

func testRequest(survey Survey) (response *http.Response) {
	handler := NewHandler()
	body, _ := json.Marshal(survey)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	w := httptest.NewRecorder()
	handler.newSurvey(w, req)
	return w.Result()
}
