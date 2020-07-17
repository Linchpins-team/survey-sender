package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSurvey(t *testing.T) {
	s := testSurvey()
	handler := NewHandler()
	body, err := json.Marshal(s)
	assert.Nil(t, err)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	w := httptest.NewRecorder()
	handler.newSurvey(w, req)
	response := w.Result()
	if response.StatusCode != 200 {
		t.Error(ioutil.ReadAll(response.Body))
	}
}
