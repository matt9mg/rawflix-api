package validators

import (
	"github.com/matt9mg/rawflix-api/entities"
	"github.com/matt9mg/rawflix-api/repositories"
	"github.com/matt9mg/rawflix-api/types"
)

type RegisterValidator interface {
	Validate(*types.Register) *types.RegisterErrors
}

type Register struct {
	userRepo repositories.UserRepository
}

func NewRegister(userRepo repositories.UserRepository) RegisterValidator {
	return &Register{
		userRepo: userRepo,
	}
}

func (r *Register) Validate(register *types.Register) *types.RegisterErrors {
	err := &types.RegisterErrors{}
	hasErrors := false

	if register.Username == "" {
		err.Username = "username cannot be blank"
		hasErrors = true
	}

	if register.Password == "" {
		err.Password = "password cannot be blank"
		hasErrors = true
	}

	if ok := entities.UserGenders[entities.UserGender(register.Gender)]; ok == "" {
		err.Gender = "invalid gender supplied"
		hasErrors = true
	}

	if ok := entities.UserCountries[entities.UserCountry(register.Country)]; ok == "" {
		err.Country = "invalid country supplied"
		hasErrors = true
	}

	if err == nil || err.Username == "" {
		exists, er := r.userRepo.UsernameExists(register.Username)

		if er != nil {
			err.Username = "unable to register at this time"
			goto returnErrs
		}

		if exists == true {
			err.Username = "username already exists"
		}
	}

returnErrs:
	if hasErrors == true {
		return err
	}

	return nil
}
