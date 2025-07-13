package entities

import (
	lib "chat/lib/auth"
	"time"
)

type User struct {
    ID        *uint
    Name      string
    Email     string
    Password  string
    UpdatedAt time.Time
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

func (u *User) Update() {
    u.UpdatedAt = time.Now()
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