package batchproductmaterial

type Create struct {
	Quantity  string `json:"quantity" binding:"required"`
	UnitPrice string `json:"unit_price" binding:"required"`
	ProductID string `json:"product_id" binding:"required,uuid"`
}

type Update struct {
	Quantity  *string `json:"quantity" binding:"required"`
	UnitPrice *string `json:"unit_price" binding:"required"`
	ProductID *string `json:"product_id" binding:"required,uuid"`
}
