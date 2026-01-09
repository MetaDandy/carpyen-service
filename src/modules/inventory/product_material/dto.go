package productmaterial

type Create struct {
	Quantity               string `json:"quantity" validate:"required"`
	BatchProductMaterialID string `json:"batch_product_material_id" validate:"required"`
	MaterialID             string `json:"material_id" validate:"required"`
}

type Update struct {
	Quantity   *string `json:"quantity" validate:"required"`
	UnitPrice  *string `json:"unit_price" validate:"required"`
	MaterialID *string `json:"material_id" validate:"required"`
}
