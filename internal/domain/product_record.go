package domain

type ProductRecord struct {
	ID             int    `json:"id"`
	LastUpdateDate string `json:"last_update_date"`
	PurchasePrice  int    `json:"purchase_price"`
	SalePrice      int    `json:"sale_price"`
	ProductID      int    `json:"product_id"`
}

type ProductRecordRequest struct {
	LastUpdateDate string `json:"last_update_date"`
	PurchasePrice  int    `json:"purchase_price"`
	SalePrice      int    `json:"sale_price"`
	ProductID      int    `json:"product_id"`
}
