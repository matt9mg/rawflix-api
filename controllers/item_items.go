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

type ItemItemsController interface {
	Like(ctx *fiber.Ctx) error
}

type ItemItems struct {
	recombee         *services.Recoombe
	movieRepo        repositories.MovieRepository
	movieTransformer transformers.RecommendationMovieTransformer
}

func NewItemItems(recombee *services.Recoombe, movieRepo repositories.MovieRepository, movieTransformer transformers.RecommendationMovieTransformer) ItemItemsController {
	return &ItemItems{
		recombee:         recombee,
		movieRepo:        movieRepo,
		movieTransformer: movieTransformer,
	}
}

func (ii *ItemItems) Like(ctx *fiber.Ctx) error {
	movieID, _ := ctx.ParamsInt("id")
	mlRequest := &types.MovieListRequest{}

	if err := ctx.QueryParser(mlRequest); err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	userID := utils.GetUserIDFromClaimsCtx(ctx)
	recommendations, err := ii.recombee.Recommendation.ReccommendItemsToItem(uint(movieID), userID, mlRequest.TotalMovies, mlRequest.Scenario)

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	ids, err := recommendations.GetIDS()

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	data, err := ii.movieRepo.GetByRecommendation(ids, userID, entities.InteractionTypeBookmark)

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	movies, err := ii.movieTransformer.MapInterfaceToMovies(data)

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	movieList, err := ii.movieTransformer.TransformRecommendationMovieWithOrder(recommendations, movies)

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	return ctx.JSON(movieList)
}
