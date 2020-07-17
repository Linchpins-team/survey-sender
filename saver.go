package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Save(survey Survey) (err error) {
	c := Configure
	// mkdir
	if err = os.MkdirAll(c.SaveAt, 0755); err != nil {
		return
	}
	// save file
	filename := fmt.Sprintf("%s/%s_%s.txt", c.SaveAt, survey.CreatedAt.Format("2006_01-02_15-04"), survey.Name)
	if err = ioutil.WriteFile(filename, []byte(survey.String()), 0644); err != nil {
		return
	}
	return nil
}
