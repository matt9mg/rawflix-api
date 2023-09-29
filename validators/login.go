package validators

import (
	"github.com/matt9mg/rawflix-api/repositories"
	"github.com/matt9mg/rawflix-api/services"
	"github.com/matt9mg/rawflix-api/types"
)

type LoginValidator interface {
	Validate(login *types.Login) *types.LoginErrors
}

type Login struct {
	userRepo repositories.UserRepository
	pwHasher services.PasswordHasher
}

func NewLogin(userRepo repositories.UserRepository, pwHasher services.PasswordHasher) LoginValidator {
	return &Login{
		userRepo: userRepo,
		pwHasher: pwHasher,
	}
}

func (l *Login) Validate(login *types.Login) *types.LoginErrors {
	loginError := &types.LoginErrors{}
	hasErrors := false

	if login.Username == "" {
		login.Username = "cannot be blank"
		hasErrors = true
	}

	if login.Password == "" {
		login.Password = "cannot be blank"
		hasErrors = true
	}

	if hasErrors == false {
		user, err := l.userRepo.FindOneByUsername(login.Username)

		if err != nil {
			loginError.GeneralError.GeneralError = err.Error()
			hasErrors = true
			goto returnErrors
		}

		if user == nil {
			loginError.GeneralError.GeneralError = "invalid username or password provided"
			hasErrors = true
			goto returnErrors
		}

		if l.pwHasher.CheckPasswordHash(login.Password, user.Password) == false {
			loginError.GeneralError.GeneralError = "invalid username or password provided"
			hasErrors = true
			goto returnErrors
		}
	}

returnErrors:
	if hasErrors == true {
		return loginError
	}

	return nil
}
