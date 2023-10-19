package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matt9mg/rawflix-api/entities"
	"github.com/matt9mg/rawflix-api/repositories"
	"github.com/matt9mg/rawflix-api/services"
	"github.com/matt9mg/rawflix-api/transformers"
	"github.com/matt9mg/rawflix-api/types"
	"github.com/matt9mg/rawflix-api/utils"
	"github.com/matt9mg/rawflix-api/validators"
	"net/http"
	"time"
)

type ViewedController interface {
	Viewed(ctx *fiber.Ctx) error
}

type Viewed struct {
	iRepo        repositories.InteractionRepository
	iValidator   validators.InteractionValidator
	iTransformer transformers.InteractionTransformer
	recombee     *services.Recoombe
}

func NewViewed(iRepo repositories.InteractionRepository, iValidator validators.InteractionValidator, iTransformer transformers.InteractionTransformer, recombee *services.Recoombe) ViewedController {
	return &Viewed{
		iRepo:        iRepo,
		iValidator:   iValidator,
		iTransformer: iTransformer,
		recombee:     recombee,
	}
}

func (v *Viewed) Viewed(ctx *fiber.Ctx) error {
	var interaction *types.Interaction

	if err := ctx.BodyParser(&interaction); err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err.Error())
	}

	if err := v.iValidator.Validate(interaction); err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err)
	}

	interactionEntity := v.iTransformer.TransformFromTypeToEntity(interaction, utils.GetUserIDFromClaimsCtx(ctx), entities.InteractionTypeDetailView)

	if err := v.iRepo.Create(interactionEntity); err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err.Error())
	}

	err := v.recombee.UserItemInteraction.AddDetailView(&types.RecombeeUserItemInteraction{
		RecombeeCascadeCreate: &types.RecombeeCascadeCreate{
			CascadeCreate: true,
		},
		UserID:           utils.UintToString(utils.GetUserIDFromClaimsCtx(ctx)),
		ItemID:           utils.UintToString(interactionEntity.MovieID),
		TimeStamp:        time.Now().Unix(),
		RecommendationID: interaction.RecommenderID,
	})

	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err.Error())
	}

	if err = v.iRepo.MarkAsSynced(interactionEntity); err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(err.Error())
	}

	return ctx.JSON("OK")
}
