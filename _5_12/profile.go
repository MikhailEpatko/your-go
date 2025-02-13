package _5_12

import (
	"errors"
	"fmt"
	"strings"
)

var template = "Имя человека: %s, возраст: %d."

func UserProfileToString(name string, age int) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	if age < 0 {
		return "", errors.New("negative age")
	}
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return "", errors.New("name cannot contain only spaces")
	}
	return fmt.Sprintf(template, trimmed, age), nil
}
