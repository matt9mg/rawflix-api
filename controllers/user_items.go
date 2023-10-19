package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matt9mg/rawflix-api/entities"
	"github.com/matt9mg/rawflix-api/repositories"
	"github.com/matt9mg/rawflix-api/services"
	"github.com/matt9mg/rawflix-api/transformers"
	"github.com/matt9mg/rawflix-api/types"
	"github.com/matt9mg/rawflix-api/utils"
	"net/http"
)

type UserItemsController interface {
	RecommendForScenario(ctx *fiber.Ctx) error
}

type UserItems struct {
	recombee         *services.Recoombe
	movieRepo        repositories.MovieRepository
	movieTransformer transformers.RecommendationMovieTransformer
}

func NewUserItems(recombee *services.Recoombe, movieRepo repositories.MovieRepository, movieTransformer transformers.RecommendationMovieTransformer) UserItemsController {
	return &UserItems{
		recombee:         recombee,
		movieRepo:        movieRepo,
		movieTransformer: movieTransformer,
	}
}

func (ui *UserItems) RecommendForScenario(ctx *fiber.Ctx) error {
	mlRequest := &types.MovieListRequest{}

	if err := ctx.QueryParser(mlRequest); err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	userID := utils.GetUserIDFromClaimsCtx(ctx)
	recommendations, err := ui.recombee.Recommendation.ReccommendItemsToUser(userID, mlRequest.TotalMovies, mlRequest.Scenario)

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	ids, err := recommendations.GetIDS()

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	data, err := ui.movieRepo.GetByRecommendation(ids, userID, entities.InteractionTypeBookmark)

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	movies, err := ui.movieTransformer.MapInterfaceToMovies(data)

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	movieList, err := ui.movieTransformer.TransformRecommendationMovieWithOrder(recommendations, movies)

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	return ctx.JSON(movieList)
}
