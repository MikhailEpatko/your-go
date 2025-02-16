package _5_13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserProfileMessage(t *testing.T) {
	a := assert.New(t)

	got, err := UserProfile("userId")
	want := "Пользователь с id userId имеет на счету 123.45 руб."
	a.Nil(err)
	a.Equal(want, got)
}

func TestUserProfileError(t *testing.T) {
	a := assert.New(t)

	msg, err := UserProfile("unknownId")
	want := "fetch error: not found"

	a.Empty(msg)
	a.NotNil(err)
	a.Equal(want, err.Error())
}
