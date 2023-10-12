package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matt9mg/rawflix-api/repositories"
	"github.com/matt9mg/rawflix-api/services"
	"github.com/matt9mg/rawflix-api/transformers"
	"net/http"
)

type HomeController interface {
	Index(ctx *fiber.Ctx) error
}

type Home struct {
	recombee         *services.Recoombe
	movieRepo        repositories.MovieRepository
	movieTransformer transformers.RecommendationMovieTransformer
}

func NewHome(recombee *services.Recoombe, movieRepo repositories.MovieRepository, movieTransformer transformers.RecommendationMovieTransformer) HomeController {
	return &Home{
		recombee:         recombee,
		movieRepo:        movieRepo,
		movieTransformer: movieTransformer,
	}
}

func (h *Home) Index(ctx *fiber.Ctx) error {
	recommendations, err := h.recombee.Recommendation.ReccommendItemsToUser(ctx.Locals("claims").(*services.Claims).UserID, 10, "popular-movies")

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	ids, err := recommendations.GetIDS()

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	movies, err := h.movieRepo.GetByRecommendation(ids)

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	movieList, err := h.movieTransformer.TransformRecommendationMovieWithOrder(recommendations, movies)

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	return ctx.JSON(movieList)
}
