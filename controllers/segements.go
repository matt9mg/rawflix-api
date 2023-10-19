package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matt9mg/rawflix-api/entities"
	"github.com/matt9mg/rawflix-api/repositories"
	"github.com/matt9mg/rawflix-api/services"
	"github.com/matt9mg/rawflix-api/transformers"
	"github.com/matt9mg/rawflix-api/types"
	"net/http"
)

type SegmentsController interface {
	GetSegmentsForUser(ctx *fiber.Ctx) error
}

type Segments struct {
	recombee     *services.Recoombe
	mTransformer transformers.RecommendationMovieTransformer
	mRepo        repositories.MovieRepository
}

func NewSegments(recombee *services.Recoombe, mTransformer transformers.RecommendationMovieTransformer, mRepo repositories.MovieRepository) SegmentsController {
	return &Segments{
		recombee:     recombee,
		mTransformer: mTransformer,
		mRepo:        mRepo,
	}
}

func (s *Segments) GetSegmentsForUser(ctx *fiber.Ctx) error {
	//userID := utils.GetUserIDFromClaimsCtx(ctx)

	recommendations, err := s.recombee.Recommendation.RecommendItemSegmentsToUser(103, 5, "home-page-rows")

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(err.Error())
	}

	var collection []*types.MovieCollection

	for _, recommendation := range recommendations.Recommendations {
		r, _ := s.recombee.Recommendation.ReccommendItemsToUserWithFilter(103, 4, "home-page-rows", recommendation.ID)

		ids, err := r.GetIDS()

		if err != nil {
			ctx.SendStatus(http.StatusBadRequest)
			return ctx.JSON(err)
		}

		movies, err := s.mRepo.GetByRecommendation(ids, 103, entities.InteractionTypeBookmark)

		if err != nil {
			ctx.SendStatus(http.StatusBadRequest)
			return ctx.JSON(err)
		}

		tMovies, err := s.mTransformer.MapInterfaceToMovies(movies)

		if err != nil {
			ctx.SendStatus(http.StatusBadRequest)
			return ctx.JSON(err)
		}

		movieList, err := s.mTransformer.TransformRecommendationMovieWithOrder(r, tMovies)

		if err != nil {
			ctx.SendStatus(http.StatusBadRequest)
			return ctx.JSON(err)
		}

		collection = append(collection, &types.MovieCollection{
			Title:  recommendation.ID,
			Movies: movieList,
		})
	}

	return ctx.JSON(collection)
}
