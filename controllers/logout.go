package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matt9mg/rawflix-api/repositories"
	"github.com/matt9mg/rawflix-api/services"
)

type LogoutController interface {
	Logout(ctx *fiber.Ctx) error
}

type Logout struct {
	jwt      services.JWTAuthenticator
	userRepo repositories.UserRepository
}

func NewLogout(jwt services.JWTAuthenticator, userRepo repositories.UserRepository) LogoutController {
	return &Logout{
		jwt:      jwt,
		userRepo: userRepo,
	}
}

func (l *Logout) Logout(ctx *fiber.Ctx) error {
	_ = l.userRepo.RemoveTokenFromByUserID(ctx.Locals("claims").(*services.Claims).UserID)

	return ctx.JSON("OK")
}
