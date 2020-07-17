package main

import (
	"fmt"
	"regexp"
	"time"
)

type Survey struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"-"`
}

func (s Survey) String() string {
	return fmt.Sprintf(
		`Name: %s
Email: %s
Time: %s,
Content: 

%s

`, s.Name, s.Email, s.CreatedAt.Format("2006 01/02 15:04"), s.Content)
}

func (s Survey) CheckAvailable() bool {
	return checkEmail(s.Email)
}

func checkEmail(email string) bool {
	return regexp.MustCompile(`(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$)`).MatchString(email)
}
