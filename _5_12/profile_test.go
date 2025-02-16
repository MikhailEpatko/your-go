package _5_12

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserProfileToStringMessage(t *testing.T) {
	a := assert.New(t)

	t.Run("check correct message 1", func(t *testing.T) {
		got, err := UserProfileToString(" Name ", 10)
		want := fmt.Sprintf(template, "Name", 10)
		a.Equal(want, got)
		a.Nil(err)
	})
	t.Run("check correct message 2", func(t *testing.T) {
		got, err := UserProfileToString(" Name ", 10)
		want := fmt.Sprintf(template, "Name", 10)
		a.Equal(want, got)
		a.Nil(err)
	})
}

func TestUserProfileToStringError(t *testing.T) {
	a := assert.New(t)

	t.Run("check error empty name", func(t *testing.T) {
		msg, err := UserProfileToString("", 10)
		want := "empty name"
		a.Equal("", msg)
		a.Empty(msg)
		a.NotNil(err)
		a.Equal(want, err.Error())
	})
	t.Run("check error spaces", func(t *testing.T) {
		msg, err := UserProfileToString("   ", 10)
		want := "name cannot contain only spaces"
		a.Equal("", msg)
		a.Empty(msg)
		a.NotNil(err)
		a.Equal(want, err.Error())
	})
	t.Run("check error negative age", func(t *testing.T) {
		msg, err := UserProfileToString("  Name ", -10)
		want := "negative age"
		a.Equal("", msg)
		a.Empty(msg)
		a.NotNil(err)
		a.Equal(want, err.Error())
	})
}
