package entities

import (
	"chat/domain/valueObject/passwordHash"
	"time"
)

type User struct {
    ID        *uint
    Name      string
    Email     string
    Password  *passwordHash.PasswordHash
    UpdatedAt time.Time
}

func (User) TableName() string {
    return "users"
}

func NewUser(name, email string, password *passwordHash.PasswordHash) User {
    return User{
        Name:     name,
        Email:    email,
        Password: password,
    }
}

func (u *User) Update() {
    u.UpdatedAt = time.Now()
}
