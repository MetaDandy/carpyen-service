package warehouse

type Create struct {
	Name        string `json:"name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Description string `json:"description"`
	Phone       string `json:"phone"`
}

type Update struct {
	Name        *string `json:"name"`
	Address     *string `json:"address"`
	Description *string `json:"description"`
	Phone       *string `json:"phone"`
}
