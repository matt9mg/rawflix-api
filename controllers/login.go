package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matt9mg/rawflix-api/repositories"
	"github.com/matt9mg/rawflix-api/services"
	"github.com/matt9mg/rawflix-api/types"
	"github.com/matt9mg/rawflix-api/validators"
	"net/http"
)

type LoginController interface {
	Login(ctx *fiber.Ctx) error
}

type Login struct {
	loginValidator validators.LoginValidator
	jwt            services.JWTAuthenticator
	userRepo       repositories.UserRepository
}

func NewLogin(loginValidator validators.LoginValidator, jwt services.JWTAuthenticator, userRepo repositories.UserRepository) LoginController {
	return &Login{
		loginValidator: loginValidator,
		jwt:            jwt,
		userRepo:       userRepo,
	}
}

func (l *Login) Login(ctx *fiber.Ctx) error {
	var login *types.Login

	if err := ctx.BodyParser(&login); err != nil {
		ctx.Response().SetStatusCode(http.StatusBadRequest)
		return ctx.JSON(&types.GeneralError{GeneralError: err.Error()})
	}

	if err := l.loginValidator.Validate(login); err != nil {
		ctx.Response().SetStatusCode(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	user, _ := l.userRepo.FindOneByUsername(login.Username)

	token, err := l.jwt.CreateToken(user)

	if err != nil {
		ctx.Response().SetStatusCode(http.StatusBadRequest)
		return ctx.JSON(&types.GeneralError{GeneralError: err.Error()})
	}

	user.Token = token

	if err = l.userRepo.Save(user); err != nil {
		ctx.Response().SetStatusCode(http.StatusBadRequest)
		return ctx.JSON(&types.GeneralError{GeneralError: err.Error()})
	}

	return ctx.JSON(user.Token)
}
