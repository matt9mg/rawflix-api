package types

type Interaction struct {
	MovieID       uint   `json:"movie_id"`
	RecommenderID string `json:"recommender_id"`
}

type InteractionErrors struct {
	MovieID       string `json:"movie_id,omitempty"`
	RecommenderID string `json:"recommender_id,omitempty"`
}
