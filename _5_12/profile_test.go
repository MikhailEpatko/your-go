package _5_12

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserProfileToStringMessage(t *testing.T) {
	a := assert.New(t)
	assertCorrectMessage := func(t testing.TB, err error, got, want string) {
		t.Helper()
		a.Equal(want, got)
		a.Nil(err)
	}
	t.Run("check correct message 1", func(t *testing.T) {
		got, err := UserProfileToString(" Name ", 10)
		want := fmt.Sprintf(template, "Name", 10)
		assertCorrectMessage(t, err, got, want)
	})

	t.Run("check correct message 2", func(t *testing.T) {
		got, err := UserProfileToString(" Name ", 10)
		want := fmt.Sprintf(template, "Name", 10)
		assertCorrectMessage(t, err, got, want)
	})
}

func TestUserProfileToStringError(t *testing.T) {
	a := assert.New(t)
	assertCorrectMessage := func(t testing.TB, got error, msg, want string) {
		t.Helper()
		a.Equal("", msg)
		a.Empty(msg)
		a.NotNil(got)
		a.Equal(want, got.Error())
	}

	t.Run("check error empty name", func(t *testing.T) {
		msg, got := UserProfileToString("", 10)
		want := "empty name"
		assertCorrectMessage(t, got, msg, want)
	})
	t.Run("check error spaces", func(t *testing.T) {
		msg, got := UserProfileToString("   ", 10)
		want := "name cannot contain only spaces"
		assertCorrectMessage(t, got, msg, want)
	})
	t.Run("check error negative age", func(t *testing.T) {
		msg, got := UserProfileToString("  Name ", -10)
		want := "negative age"
		assertCorrectMessage(t, got, msg, want)
	})
}
