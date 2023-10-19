package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matt9mg/rawflix-api/services"
	"net/http"
)

type JWTMiddleware interface {
	Validate(ctx *fiber.Ctx) error
}

type JWT struct {
	jwt services.JWTAuthenticator
}

func NewJWT(jwt services.JWTAuthenticator) JWTMiddleware {
	return &JWT{
		jwt: jwt,
	}
}

func (j *JWT) Validate(ctx *fiber.Ctx) error {
	claims, err := j.jwt.Validate(ctx)

	if err != nil {
		ctx.SendStatus(http.StatusForbidden)
		return ctx.JSON(err.Error())
	}

	ctx.Locals("claims", claims)

	return ctx.Next()
}
