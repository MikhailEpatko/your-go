package _11_8

import (
	"errors"
)

const (
	MinAge                = 15
	MaxAge                = 80
	MinGrade              = 1
	MaxGrade              = 5
	GreatAge              = 30
	MinGradeAfterGreatAge = 3
)

var (
	ErrEmptyName         = errors.New("name cannot be empty")
	ErrTooYoung          = errors.New("too young")
	ErrTooOld            = errors.New("too old")
	ErrGradeOutOfRange   = errors.New("grade out of range")
	ErrTooLowGradeForAge = errors.New("too low grade for age")
	ErrIncorrectEmail    = errors.New("incorrect email")
)

type Student struct {
	Name  string
	Age   int
	Grade int
	Email string
}

func NewStudent(name string, age, grade int, email string) (*Student, error) {
	return &Student{}, nil
}
