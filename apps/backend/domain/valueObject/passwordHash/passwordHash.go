package passwordHash

import (
	lib "chat/lib/auth"
	"fmt"
)

type PasswordHash struct {
    passwordHash string
}

// 生パスワードからハッシュを生成（新規登録用）
func NewPasswordHash(password string) (*PasswordHash, error) {
    hashed, err := lib.Hash(password)
    if err != nil {
        return nil, fmt.Errorf("failed to hash password: %w", err)
    }
    return &PasswordHash{passwordHash: hashed}, nil
}

// 既存のハッシュ値からVOを生成（ログイン・検証用）
func NewPasswordHashFromHash(hash string) *PasswordHash {
    return &PasswordHash{passwordHash: hash}
}

// ハッシュ値を文字列で取得
func (p *PasswordHash) ToString() string {
    return p.passwordHash
}

// 生パスワードとハッシュ値を検証
func (p *PasswordHash) Verify(plain string) bool {
    return lib.Verify(p.passwordHash, plain)
}