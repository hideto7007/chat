package lib

import "golang.org/x/crypto/bcrypt"

func Hash(plain string) (string, error) {
    hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
    return string(hashed), err
}

func Verify(hashed, plain string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
    return err == nil
}