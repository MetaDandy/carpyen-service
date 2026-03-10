package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
)

type Purchase struct {
	ID            string `json:"id"`
	Date          string `json:"date"`
	ReceiptNumber string `json:"receipt_number"`

	MaterialDetails *MaterialDetails `json:"material_details"`
	ProductDetails  *ProductDetails  `json:"product_details"`
	Supplier        *Supplier        `json:"supplier"`
	User            *User            `json:"user,omitzero"`
}

func PurchaseToDto(m *model.Purchase) Purchase {
	dto := Purchase{
		ID:            m.ID.String(),
		Date:          m.Date.Format("2006-01-02"),
		ReceiptNumber: m.ReceiptNumber,
	}

	if m.Supplier.ID != (uuid.UUID{}) {
		sup := SupplierToDto(&m.Supplier)
		dto.Supplier = &sup
	}

	if m.User.ID != (uuid.UUID{}) {
		usr := UserToDto(&m.User)
		dto.User = &usr
	}

	return dto
}

func PurchaseToListDto(m []model.Purchase) []Purchase {
	out := make([]Purchase, len(m))
	for i, item := range m {
		out[i] = PurchaseToDto(&item)
	}
	return out
}
