package types

import (
	"github.com/matt9mg/rawflix-api/entities"
	"strconv"
)

type RecombeeCascadeCreate struct {
	CascadeCreate bool `json:"!cascadeCreate"`
}

type RecombeeMovieItem struct {
	*RecombeeCascadeCreate
	ItemId    string   `json:"item_id,omitempty"`
	Actors    []string `json:"actors"`
	Country   []string `json:"country"`
	Directors []string `json:"directors"`
	Genres    []string `json:"genres"`
	Language  []string `json:"language"`
	Plot      string   `json:"plot"`
	Poster    string   `json:"poster"`
	Rated     string   `json:"rated"`
	Released  string   `json:"released"`
	Runtime   string   `json:"runtime"`
	Title     string   `json:"title"`
	Writers   []string `json:"writers"`
	Year      string   `json:"year"`
	Added     int64    `json:"added"`
}

type RecombeeUserItem struct {
	*RecombeeCascadeCreate
	UserId  string               `json:"user_id,omitempty"`
	Country entities.UserCountry `json:"country"`
	Gender  entities.UserGender  `json:"gender"`
}

type RecombeeRecommendations struct {
	RecommendID     string                    `json:"recommId"`
	Recommendations []*RecombeeRecommendation `json:"recomms"`
}

type RecombeeRecommendation struct {
	ID string `json:"id"`
}

type RecombeeUserItemInteraction struct {
	*RecombeeCascadeCreate
	UserID           string `json:"userId"`
	ItemID           string `json:"itemId"`
	TimeStamp        int64  `json:"timestamp"`
	Duration         int    `json:"duration,omitempty"`
	RecommendationID string `json:"recommId,omitempty"`
}

func (rr *RecombeeRecommendations) GetIDS() ([]uint, error) {
	var ids []uint

	for _, recommendation := range rr.Recommendations {
		id, err := strconv.Atoi(recommendation.ID)

		if err != nil {
			return nil, err
		}

		ids = append(ids, uint(id))
	}

	return ids, nil
}
