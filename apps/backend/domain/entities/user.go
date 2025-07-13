package entities

import (
	"chat/lib/auth"
)

type User struct {
    ID        *uint
    Name      string
    Email     string
    Password  string
}

func (User) TableName() string {
    return "users"
}

func NewUser(name, email, plainPassword string) (*User, error) {
    hashed, err := lib.Hash(plainPassword)
    if err != nil {
        return nil, err
    }
    return &User{
        Name:     name,
        Email:    email,
        Password: hashed,
    }, nil
}

func (u *User) Update(name, email string) {
    u.Name = name
    u.Email = email
}


func (u *User) UpdatePassword(plainPassword string) error {
    hashed, err := lib.Hash(plainPassword)
    if err != nil {
        return err
    }
    u.Password = hashed
    return nil
}


func (u *User) CheckPassword(plainPassword string) bool {
    return lib.Verify(u.Password, plainPassword)
}