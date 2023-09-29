package types

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginErrors struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	GeneralError
}
