package supplier

type Create struct {
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
	UserID  string `json:"user_id"`
}

type Update struct {
	Name    *string `json:"name"`
	Contact *string `json:"contact"`
	Phone   *string `json:"phone"`
	Email   *string `json:"email"`
	Address *string `json:"address"`
}
