package _12_1

import (
	"errors"
	"fmt"
)

type Mage struct {
	Name string
}

func NewMage(name string) (*Mage, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Mage{Name: name}, nil
}

func (m Mage) Attack() string {
	return fmt.Sprintf("Маг %s колдует огненный шар.", m.Name)
}
