package v1

import (
	"os"
	"testing"
)

func Setup() {
}

func Teardown() {
}

func TestMain(m *testing.M) {
	Setup()
	code := m.Run()
	Teardown()
	os.Exit(code)
}

func TestLogin(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
