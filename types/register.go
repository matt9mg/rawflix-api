package types

import "github.com/matt9mg/rawflix-api/entities"

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Country  string `json:"country"`
}

type RegisterErrors struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Country  string `json:"country,omitempty"`
	GeneralError
}

type RegisterFieldData struct {
	Countries map[entities.UserCountry]entities.UserCountry `json:"countries"`
	Genders   map[entities.UserGender]entities.UserGender   `json:"genders"`
}
