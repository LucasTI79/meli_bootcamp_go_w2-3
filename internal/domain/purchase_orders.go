package domain

type PurchaseOrders struct {
	ID              int    `json:"id"`
	OrderNumber     string `json:"order_number"`
	OrderDate       string `json:"order_date"`
	TrackingCode    string `json:"tracking_code"`
	BuyerID         int    `json:"buyer_id"`
	ProductRecordID int    `json:"product_record_id"`
	OrderStatusID   int    `json:"order_status_id"`
}

type BuyerOrders struct {
	ID                  int    `json:"id"`
	CardNumberID        string `json:"card_number_id"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	PurchaseOrdersCount int    `json:"purchase_orders_count"`
}

type PurchaseOrdersResponse struct {
	Data []PurchaseOrders `json:"data"`
}
