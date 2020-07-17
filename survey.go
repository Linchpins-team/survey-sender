package main

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Survey struct {
	gorm.Model
	Name string
	Email string
	Content string
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