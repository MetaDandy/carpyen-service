package batchmaterialsupplier

type Create struct {
	Quantity  string `json:"quantity"`
	UnitPrice string `json:"unit_price"`
	TotalCost string `json:"total_cost"`
	Stock     string `json:"stock"`
}

type Update struct {
	Quantity  *string `json:"quantity"`
	UnitPrice *string `json:"unit_price"`
	TotalCost *string `json:"total_cost"`
	Stock     *string `json:"stock"`
}
