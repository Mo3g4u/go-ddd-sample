package model

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID   string
	Name string
}

type UserCreateConfig struct {
	Name string
}

func NewUser(conf UserCreateConfig) (*User, error) {
	u := &User{
		ID: uuid.NewString(),
	}
	if err := u.ChangeName(conf.Name); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) ChangeName(name string) error {
	if name == "" {
		return fmt.Errorf("ユーザ名は必須です。")
	}
	if len(name) < 3 {
		return fmt.Errorf("ユーザ名は3文字以上です。%s", name)
	}
	u.Name = name
	return nil
}
