package transformers

import (
	"github.com/matt9mg/rawflix-api/entities"
	"github.com/matt9mg/rawflix-api/types"
	"github.com/matt9mg/rawflix-api/utils"
)

type RecommendationMovieTransformer interface {
	TransformRecommendationMovieWithOrder(recommendations *types.RecombeeRecommendations, movies []*entities.Movie) ([]*types.Movie, error)
}

type RecommendationMovie struct {
}

func NewRecommendationMovie() RecommendationMovieTransformer {
	return &RecommendationMovie{}
}

func (rm *RecommendationMovie) TransformRecommendationMovieWithOrder(recommendations *types.RecombeeRecommendations, movies []*entities.Movie) ([]*types.Movie, error) {
	var moviesInOrder []*types.Movie

	for _, recommended := range recommendations.Recommendations {
		idToCompare, err := utils.StringToUint(recommended.ID)

		if err != nil {
			return nil, err
		}

		for _, movie := range movies {
			if idToCompare == movie.ID {
				moviesInOrder = append(moviesInOrder, &types.Movie{
					ID:               movie.ID,
					Title:            movie.Title,
					Plot:             movie.Plot,
					Poster:           movie.Poster,
					RecommendationID: recommendations.RecommendID,
				})
			}
		}
	}

	return moviesInOrder, nil
}
