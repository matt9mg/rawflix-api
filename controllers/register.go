package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matt9mg/rawflix-api/entities"
	"github.com/matt9mg/rawflix-api/repositories"
	"github.com/matt9mg/rawflix-api/services"
	"github.com/matt9mg/rawflix-api/types"
	"github.com/matt9mg/rawflix-api/validators"
	"net/http"
)

type RegisterController interface {
	Register(*fiber.Ctx) error
	GetRegisterFieldData(*fiber.Ctx) error
}

type Register struct {
	validator validators.RegisterValidator
	pwHasher  services.PasswordHasher
	userRepo  repositories.UserRepository
}

func NewRegister(validator validators.RegisterValidator, pwHasher services.PasswordHasher, userRepo repositories.UserRepository) RegisterController {
	return &Register{
		validator: validator,
		pwHasher:  pwHasher,
		userRepo:  userRepo,
	}
}

func (r *Register) Register(ctx *fiber.Ctx) error {
	var register *types.Register

	if err := ctx.BodyParser(&register); err != nil {
		ctx.Response().Header.SetStatusCode(http.StatusBadRequest)
		return ctx.JSON(&types.RegisterErrors{
			GeneralError: types.GeneralError{
				GeneralError: err.Error(),
			},
		})
	}

	if err := r.validator.Validate(register); err != nil {
		ctx.Response().Header.SetStatusCode(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	pw, err := r.pwHasher.HashPassword(register.Password)

	if err != nil {
		ctx.Response().Header.SetStatusCode(http.StatusBadRequest)
		return ctx.JSON(&types.RegisterErrors{
			GeneralError: types.GeneralError{
				GeneralError: err.Error(),
			},
		})
	}

	newUser := &entities.User{
		Username: register.Username,
		Password: pw,
		Country:  entities.UserCountry(register.Country),
		Gender:   entities.UserGender(register.Gender),
	}

	if err = r.userRepo.Create(newUser); err != nil {
		ctx.Response().Header.SetStatusCode(http.StatusBadRequest)
		return ctx.JSON(&types.RegisterErrors{
			GeneralError: types.GeneralError{
				GeneralError: err.Error(),
			},
		})
	}

	ctx.Response().Header.SetStatusCode(http.StatusCreated)
	return ctx.JSON("success")
}

func (r *Register) GetRegisterFieldData(ctx *fiber.Ctx) error {
	return ctx.JSON(&types.RegisterFieldData{
		Countries: entities.UserCountries,
		Genders:   entities.UserGenders,
	})
}
