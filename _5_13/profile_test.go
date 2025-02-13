package _5_13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserProfileMessage(t *testing.T) {
	a := assert.New(t)
	assertInfo := func(t testing.TB, got, want string, err error) {
		t.Helper()
		a.Nil(err)
		a.Equal(want, got)
	}

	got, err := UserProfile("userId")
	want := "Пользователь с id userId имеет на счету 123.45 руб."
	assertInfo(t, got, want, err)
}

func TestUserProfileError(t *testing.T) {
	a := assert.New(t)
	assertInfo := func(t testing.TB, msg, want string, got error) {
		t.Helper()
		a.Empty(msg)
		a.NotNil(got)
		a.Equal(want, got.Error())
	}

	msg, err := UserProfile("unknownId")
	want := "fetch error: not found"
	assertInfo(t, msg, want, err)
}
