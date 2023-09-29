package services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matt9mg/rawflix-api/entities"
	"github.com/matt9mg/rawflix-api/repositories"
)

type JWTAuthenticator interface {
	Validate(token string) (*jwt.Token, error)
	CreateToken(user *entities.User) (string, error)
}

type JWT struct {
	secret   []byte
	expires  *jwt.NumericDate
	userRepo repositories.UserRepository
}

func NewJWT(secret []byte, expires *jwt.NumericDate, userRepo repositories.UserRepository) JWTAuthenticator {
	return &JWT{
		secret:   secret,
		expires:  expires,
		userRepo: userRepo,
	}
}

func (j *JWT) CreateToken(user *entities.User) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": j.expires,
		"userID":    user.ID,
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.secret)
}

func (j *JWT) Validate(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

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

	claims := token.Claims.(jwt.MapClaims)

	match, err := j.userRepo.IDAndTokenExists(claims["userID"].(uint), tokenString)

	if err != nil {
		return nil, err
	}

	if match == false {
		return nil, errors.New("unable to validate stored claims")
	}

	return token, nil
}
