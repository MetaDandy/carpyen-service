package batchproductmaterial

type Create struct {
	Quantity  uint   `json:"quantity" binding:"required"`
	UnitPrice string `json:"unit_price" binding:"required"`
	Stock     string `json:"stock" binding:"required"`
	ProductID string `json:"product_id" binding:"required,uuid"`
}

type Update struct {
	Quantity  *uint   `json:"quantity" binding:"required"`
	UnitPrice *string `json:"unit_price" binding:"required"`
	Stock     *string `json:"stock" binding:"required"`
	ProductID *string `json:"product_id" binding:"required,uuid"`
}
