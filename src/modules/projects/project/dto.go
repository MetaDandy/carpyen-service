package project

import "github.com/MetaDandy/carpyen-service/src/enum"

type Create struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	ClientID    string `json:"client_id"`
}
type Update struct {
	Name        *string      `json:"name"`
	Description *string      `json:"description"`
	Location    *string      `json:"location"`
	State       *enum.Status `json:"state"`
	ClientID    *string       `json:"client_id"`
}
