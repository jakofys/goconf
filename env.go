package goenv

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var env = make(map[string]string)

type WrongEnvFileFormatError struct {
	filename string
}

func (e *WrongEnvFileFormatError) Error() string {
	return fmt.Sprintf("Env filename '%s' is wrong formatted, verify correct syntax of content file!", e.filename)
}

func ImportEnvFile(filename string) error {
	dir, _ := os.Getwd()
	if file, err := ioutil.ReadFile(dir + "/" + filename); err != nil {
		return err
	} else {
		regComment := regexp.MustCompile(``)
		regVar := regexp.MustCompile(``)
		for _, line := range strings.Split(string(file), "\n") {
			if regVar.MatchString(line) {
				envVar := strings.Split(line, "=")
				env[strings.ToLower(envVar[0])] = envVar[1]
			} else if !regComment.MatchString(line) {
				return &WrongEnvFileFormatError{
					filename: filename,
				}
			}
		}
	}
	return nil
}

func ImportFromOS(pattern string) {
	reg := regexp.MustCompile(strings.ReplaceAll(pattern, "*", ".*"))
	for _, line := range os.Environ() {
		envVar := strings.Split(line, "=")
		if reg.MatchString(envVar[0]) {
			env[strings.ToLower(envVar[0])] = envVar[1]
		}
	}
}

func Sprintf(text string) string {
	reg := regexp.MustCompile(`%\w+%`)
	for _, key := range reg.FindAllString(text, -1) {
		if value, ok := env[key[1:len(key)-1]]; ok {
			text = strings.ReplaceAll(text, key, value)
		}
	}
	return text
}

func GetEnv(key string) string {
	if envVar, ok := env[strings.ToLower(key)]; ok {
		return envVar
	} else {
		return ""
	}
}
