package user

type UserDTO struct {
	InternalId string `json:"_id,omitempty"`
	ExternalId string `json:"externalid"`
	Name       string `json:"name"`
	Login      string `json:"login"`
	Token      string `json:"token"`
}
