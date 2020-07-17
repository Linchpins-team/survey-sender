package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	initConfig()
	os.Exit(m.Run())
}
