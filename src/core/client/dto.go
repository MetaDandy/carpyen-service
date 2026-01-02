package client

type Create struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone"`
	Address         string `json:"address"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Update struct {
	Name            *string `json:"name"`
	Email           *string `json:"email"`
	Phone           *string `json:"phone"`
	Address         *string `json:"address"`
	Password        *string `json:"password"`
	ConfirmPassword *string `json:"confirm_password"`
}

type UpdateProfile struct {
	Name    *string `json:"name"`
	Phone   *string `json:"phone"`
	Address *string `json:"address"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
