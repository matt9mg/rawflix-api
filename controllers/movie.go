package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matt9mg/rawflix-api/repositories"
	"github.com/matt9mg/rawflix-api/transformers"
	"net/http"
)

type MovieController interface {
	Get(ctx *fiber.Ctx) error
}

type Movie struct {
	mRepo        repositories.MovieRepository
	mTransformer transformers.RecommendationMovieTransformer
}

func NewMovie(mRepo repositories.MovieRepository, mTransformer transformers.RecommendationMovieTransformer) MovieController {
	return &Movie{
		mRepo:        mRepo,
		mTransformer: mTransformer,
	}
}

func (m *Movie) Get(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	movie, err := m.mRepo.FindByID(uint(id))

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(err.Error())
	}

	movieType, err := m.mTransformer.MovieToType(movie)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(err.Error())
	}

	return ctx.JSON(movieType)
}
