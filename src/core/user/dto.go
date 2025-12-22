package user

type Create struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone"`
	Address         string `json:"address"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	Role            string `json:"role" validate:"required"`
}

type Update struct {
	Name            *string `json:"name"`
	Email           *string `json:"email"`
	Phone           *string `json:"phone"`
	Address         *string `json:"address"`
	Password        *string `json:"password"`
	ConfirmPassword *string `json:"confirm_password"`
	Role            *string `json:"role"`
}

type UpdateProfile struct {
	Name    *string `json:"name"`
	Phone   *string `json:"phone"`
	Address *string `json:"address"`
	Email   *string `json:"email"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
