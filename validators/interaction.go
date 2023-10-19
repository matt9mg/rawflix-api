package validators

import (
	"github.com/matt9mg/rawflix-api/repositories"
	"github.com/matt9mg/rawflix-api/types"
)

type InteractionValidator interface {
	Validate(interaction *types.Interaction) *types.InteractionErrors
}

type Interaction struct {
	mRepo repositories.MovieRepository
}

func NewInteraction(mRepo repositories.MovieRepository) InteractionValidator {
	return &Interaction{
		mRepo: mRepo,
	}
}

func (i *Interaction) Validate(favourite *types.Interaction) *types.InteractionErrors {
	fErr := &types.InteractionErrors{}
	hasErrs := false

	if favourite.MovieID == 0 {
		fErr.MovieID = "invalid movie supplied"
		hasErrs = true
	} else {
		movie, err := i.mRepo.FindByID(favourite.MovieID)

		if err != nil {
			fErr.MovieID = "invalid movie supplied"
			hasErrs = true
			goto end
		}

		if movie == nil {
			fErr.MovieID = "invalid movie supplied"
			hasErrs = true
		}
	}

end:
	if hasErrs == true {
		return fErr
	}

	return nil
}
