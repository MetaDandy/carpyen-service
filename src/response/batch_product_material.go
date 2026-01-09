package response

import "github.com/MetaDandy/carpyen-service/src/model"

type BatchProductMaterial struct {
	ID        string `json:"id"`
	Quantity  uint   `json:"quantity"`
	UnitPrice string `json:"unit_price"`
	TotalCost string `json:"total_cost"`
	Stock     string `json:"stock"`

	Product Product `json:"product"`
	User    User    `json:"user"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func BatchProductMaterialToDto(bpm *model.BatchProductMaterial) BatchProductMaterial {
	return BatchProductMaterial{
		ID:        bpm.ID.String(),
		Quantity:  bpm.Quantity,
		UnitPrice: bpm.UnitPrice.String(),
		TotalCost: bpm.TotalCost.String(),
		Stock:     bpm.Stock.String(),

		Product: ProductToDto(&bpm.Product),
		User:    UserToDto(&bpm.User),

		CreatedAt: bpm.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: bpm.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func BatchProductMaterialToListDto(bpm []model.BatchProductMaterial) []BatchProductMaterial {
	result := make([]BatchProductMaterial, len(bpm))
	for i, v := range bpm {
		result[i] = BatchProductMaterialToDto(&v)
	}
	return result
}
