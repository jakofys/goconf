package goenv_test

import (
	"testing"

	"github.com/jakofys/goenv"
)

func TestImportEnvFile(t *testing.T) {
	filename := ".env"
	if err := goenv.ImportEnvFile(filename); err != nil {
		t.Error(err)
	}
}
