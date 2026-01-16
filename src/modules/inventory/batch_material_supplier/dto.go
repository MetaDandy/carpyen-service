package batchmaterialsupplier

type Create struct {
	Quantity   string `json:"quantity"`
	UnitPrice  string `json:"unit_price"`
	MaterialID string `json:"material_id"`
	SupplierID string `json:"supplier_id"`
}

type Update struct {
	Quantity   *string `json:"quantity"`
	UnitPrice  *string `json:"unit_price"`
	MaterialID *string `json:"material_id"`
	SupplierID *string `json:"supplier_id"`
}
