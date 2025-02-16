package _12_1

import (
	"errors"
	"fmt"
)

type Archer struct {
	Name string
}

func NewArcher(name string) (*Archer, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Archer{Name: name}, nil
}

func (a Archer) Attack() string {
	return fmt.Sprintf("Лучник %s выпускает град стрел.", a.Name)
}
