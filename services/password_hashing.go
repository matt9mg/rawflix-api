package services

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password string, hash string) bool
}

type Password struct {
	config *PasswordConfig
}

type PasswordConfig struct {
	Cost int
}

func NewPassword(passwordCfg *PasswordConfig) PasswordHasher {
	return &Password{
		config: passwordCfg,
	}
}

func (p *Password) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), p.config.Cost)
	return string(bytes), err
}

func (p *Password) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
