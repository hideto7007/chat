package passwordHash

import (
	lib "chat/lib/auth"
	"fmt"
)

type PasswordHash struct {
    passwordHash string
}

func NewPasswordHash(password string) (*PasswordHash, error) {
    hashed, err := lib.Hash(password)
    if err != nil {
        return nil, fmt.Errorf("failed to hash password: %w", err)
    }
    return &PasswordHash{passwordHash: hashed}, nil
}

func (p *PasswordHash) ToString() string {
	return p.passwordHash
}