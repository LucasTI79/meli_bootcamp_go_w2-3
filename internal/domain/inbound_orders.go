package domain

type InboundOrders struct {
	ID             int    `json:"id"`
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeID     string `json:"employee_id"`
	ProductBatchID string `json:"product_batch_id"`
	WarehouseID    string `json:"warehouse_id"`
}
