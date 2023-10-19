package transformers

import (
	"encoding/json"
	"github.com/matt9mg/rawflix-api/entities"
	"github.com/matt9mg/rawflix-api/types"
	"github.com/matt9mg/rawflix-api/utils"
	"gorm.io/gorm"
	"strings"
)

type RecommendationMovieTransformer interface {
	TransformRecommendationMovieWithOrder(recommendations *types.RecombeeRecommendations, movies []*entities.Movie) ([]*types.Movie, error)
	MapInterfaceToMovies([]map[string]interface{}) ([]*entities.Movie, error)
	MovieToType(movie *entities.Movie) (*types.Movie, error)
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
				movieInOrder, err := rm.MovieToType(movie)

				if err != nil {
					return nil, err
				}

				moviesInOrder = append(moviesInOrder, movieInOrder)
			}
		}
	}

	return moviesInOrder, nil
}

func (rm *RecommendationMovie) MapInterfaceToMovies(data []map[string]interface{}) ([]*entities.Movie, error) {
	var movies []*entities.Movie

	for _, datum := range data {
		genres := strings.Split(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(datum["genre"].(string), "]", ""), "[", ""), "\"", ""), ",")

		dataGenres, err := json.Marshal(genres)

		if err != nil {
			return nil, err
		}

		movie := &entities.Movie{
			Model: gorm.Model{
				ID: uint(datum["id"].(int64)),
			},
			Title:   datum["title"].(string),
			Runtime: datum["runtime"].(string),
			Genre:   dataGenres,
			Plot:    datum["plot"].(string),
			Poster:  datum["poster"].(string),
		}

		if datum["interaction_id"] != nil {
			movie.Interaction = &entities.Interaction{
				Model: gorm.Model{
					ID: uint(datum["interaction_id"].(int64)),
				},
			}
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func (rm *RecommendationMovie) MovieToType(movie *entities.Movie) (*types.Movie, error) {
	genres, err := utils.FromJsonDataTypeToSliceString(movie.Genre)

	if err != nil {
		return nil, err
	}

	movieType := &types.Movie{
		ID:       movie.ID,
		Title:    movie.Title,
		Plot:     movie.Plot,
		Poster:   movie.Poster,
		Duration: movie.Runtime,
		Genres:   genres,
	}

	if movie.Interaction != nil {
		movieType.Bookmarked = true
	}

	return movieType, nil
}
