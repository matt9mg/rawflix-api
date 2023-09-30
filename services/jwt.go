package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matt9mg/rawflix-api/entities"
	"github.com/matt9mg/rawflix-api/repositories"
	"time"
)

const jwtTokenHeader = "x-jwt-token"

type JWTAuthenticator interface {
	Validate(ctx *fiber.Ctx) (*Claims, error)
	CreateToken(user *entities.User) (string, error)
}

type JWT struct {
	secret   []byte
	userRepo repositories.UserRepository
}

func NewJWT(secret []byte, userRepo repositories.UserRepository) JWTAuthenticator {
	return &JWT{
		secret:   secret,
		userRepo: userRepo,
	}
}

type Claims struct {
	UserID uint `json:"user_id"`
	*jwt.RegisteredClaims
}

func (j *JWT) CreateToken(user *entities.User) (string, error) {
	claims := &Claims{
		UserID: user.ID,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.secret)
}

func (j *JWT) Validate(ctx *fiber.Ctx) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(ctx.Get(jwtTokenHeader), claims, func(token *jwt.Token) (any, error) {
		return j.secret, nil
	})

	if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, errors.New("malformed token")
	}

	if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return nil, errors.New("invalid signature")
	}

	if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return nil, errors.New("token is not valid")
	}

	if token.Valid == false {
		return nil, errors.New("unable to handle token")
	}

	match, err := j.userRepo.IDAndTokenExists(claims.UserID, ctx.Get(jwtTokenHeader))

	if err != nil {
		return nil, err
	}

	if match == false {
		return nil, errors.New("unable to validate stored claims")
	}

	return claims, nil
}
