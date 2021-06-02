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

func TestImportFromOS(t *testing.T) {
	pattern := "HOM*"
	goenv.ImportFromOS(pattern)
	if goenv.GetEnv("HOME") == "" {
		t.Error("After import env var of the linux os, HOME env var not found")
	}
}

func TestSprintf(t *testing.T) {
	filename := ".env"
	if err := goenv.ImportEnvFile(filename); err != nil {
		t.Error(err)
	}

	if goenv.Sprintf("hello, i'm %name% and i'm %age% years!") != "hello, i'm jacques and i'm 22 years!" {
		t.Error("Not corresponding after print interpolation syntax")
	}

}
