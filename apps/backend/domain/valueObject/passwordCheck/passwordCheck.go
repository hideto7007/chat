package passwordCheck

import (
	"fmt"
	"regexp"
)

type PasswordCheck struct {
    Password string
}

func NewPasswordCheck(password string) *PasswordCheck {
	return &PasswordCheck{
		Password: password,
	}
}
func (p *PasswordCheck) Validate() error {
    password := p.Password

    if len(password) < 8 || len(password) > 100 {
        return fmt.Errorf("password must be 8-100 characters")
    }
    if !regexp.MustCompile(`[a-z]`).MatchString(password) {
        return fmt.Errorf("password must include a lowercase letter")
    }
    if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
        return fmt.Errorf("password must include an uppercase letter")
    }
    if !regexp.MustCompile(`\d`).MatchString(password) {
        return fmt.Errorf("password must include a number")
    }
    if !regexp.MustCompile(`[!?\-_@]`).MatchString(password) {
        return fmt.Errorf("password must include a special character (!?-_@)")
    }
    if !regexp.MustCompile(`^[a-zA-Z\d!?\-_@]+$`).MatchString(password) {
        return fmt.Errorf("password contains invalid characters")
    }
    return nil
}
