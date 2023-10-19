package types

type MovieListRequest struct {
	TotalMovies int    `json:"total_movies" query:"total_movies"`
	Scenario    string `json:"scenario" query:"scenario"`
}
