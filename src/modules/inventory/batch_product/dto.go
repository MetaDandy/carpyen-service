package batchproduct

type Create struct {
	Quantity    string `json:"quantity"`
	UnitPrice   string `json:"unit_price"`
	ProductID   string `json:"product_id"`
	WarehouseID string `json:"warehouse_id"`
}

type Update struct {
	Quantity    *string `json:"quantity"`
	UnitPrice   *string `json:"unit_price"`
	ProductID   *string `json:"product_id"`
	WarehouseID *string `json:"warehouse_id"`
}
