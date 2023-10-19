package transformers

import (
	"github.com/matt9mg/rawflix-api/entities"
	"github.com/matt9mg/rawflix-api/types"
)

type InteractionTransformer interface {
	TransformFromTypeToEntity(i *types.Interaction, userID uint, interactionType entities.InteractionType) *entities.Interaction
}

type Interaction struct {
}

func NewInteraction() InteractionTransformer {
	return &Interaction{}
}

func (*Interaction) TransformFromTypeToEntity(i *types.Interaction, userID uint, interactionType entities.InteractionType) *entities.Interaction {
	return &entities.Interaction{
		UserID:        userID,
		MovieID:       i.MovieID,
		Type:          interactionType,
		RecommenderID: i.RecommenderID,
	}
}
