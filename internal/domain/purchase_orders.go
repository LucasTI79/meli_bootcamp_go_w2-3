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

type PurchaseOrdersResponse struct {
	Data []PurchaseOrders `json:"data"`
}
