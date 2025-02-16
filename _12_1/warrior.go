package _12_1

import (
	"errors"
	"fmt"
)

type Warrior struct {
	Name string
}

func NewWarrior(name string) (*Warrior, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Warrior{Name: name}, nil
}

func (w Warrior) Attack() string {
	return fmt.Sprintf("Воин %s бьет мечом.", w.Name)
}
