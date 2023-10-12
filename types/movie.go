package types

type Movie struct {
	ID               uint   `json:"id"`
	Title            string `json:"title"`
	Plot             string `json:"plot"`
	Poster           string `json:"poster"`
	RecommendationID string `json:"recommendation_id"`
}
