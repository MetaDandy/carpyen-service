package batchmaterial

type Create struct {
	Quantity    string `json:"quantity"`
	UnitPrice   string `json:"unit_price"`
	MaterialID  string `json:"material_id"`
	WarehouseID string `json:"warehouse_id"`
}

type Update struct {
	Quantity    *string `json:"quantity"`
	UnitPrice   *string `json:"unit_price"`
	MaterialID  *string `json:"material_id"`
	WarehouseID *string `json:"warehouse_id"`
}
