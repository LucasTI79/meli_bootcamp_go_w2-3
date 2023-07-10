package domain

type InboundOrder struct {
	ID             int    `json:"id"`
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeID     string `json:"lemployee_id"`
	ProductBatchID string `json:"product_batch_id"`
	WarehouseID    string `json:"warehouse_id"`
}

type RequestCreateInboundOrders struct {
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeID     string `json:"lemployee_id"`
	ProductBatchID string `json:"product_batch_id"`
	WarehouseID    string `json:"warehouse_id"`
}

type RequestUpdateInboundOrders struct {
	OrderDate      *string `json:"order_date"`
	OrderNumber    *string `json:"order_number"`
	EmployeeID     *string `json:"lemployee_id"`
	ProductBatchID *string `json:"product_batch_id"`
	WarehouseID    *string `json:"warehouse_id"`
}
