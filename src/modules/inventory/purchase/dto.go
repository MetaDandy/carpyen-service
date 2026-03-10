package purchase

type Create struct {
	Date          string `json:"date"`
	ReceiptNumber string `json:"receipt_number"`
	SupplierID    string `json:"supplier_id"`
}

type Update struct {
	Date          *string `json:"date"`
	ReceiptNumber *string `json:"receipt_number"`
	SupplierID    *string `json:"supplier_id"`
}
