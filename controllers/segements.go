package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matt9mg/rawflix-api/services"
	"log"
	"net/http"
)

type SegmentsController interface {
	GetSegmentsForUser(ctx *fiber.Ctx) error
}

type Segments struct {
	recombee *services.Recoombe
}

func NewSegments(recombee *services.Recoombe) SegmentsController {
	return &Segments{
		recombee: recombee,
	}
}

func (s *Segments) GetSegmentsForUser(ctx *fiber.Ctx) error {
	recommendations, err := s.recombee.Recommendation.RecommendItemSegmentsToUser(103, 5, "home-page-rows")

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(err.Error())
	}

	log.Println(len(recommendations.Recommendations))

	for _, _ = range recommendations.Recommendations {
		log.Println(s.recombee.Recommendation.ReccommendItemsToUserWithFilter(103, 4, "home-page-rows", "Family"))
	}

	return nil
}
