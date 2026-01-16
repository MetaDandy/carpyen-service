package batchproductsupplier

type Create struct {
	Quantity   string `json:"quantity"`
	UnitPrice  string `json:"unit_price"`
	ProductID  string `json:"product_id"`
	SupplierID string `json:"supplier_id"`
}

type Update struct {
	Quantity   *string `json:"quantity"`
	UnitPrice  *string `json:"unit_price"`
	ProductID  *string `json:"product_id"`
	SupplierID *string `json:"supplier_id"`
}
