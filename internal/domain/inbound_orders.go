package domain

type InboundOrders struct {
	ID             int    `json:"id"`
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeID     int    `json:"employee_id"`
	ProductBatchID int    `json:"product_batch_id"`
	WarehouseID    int    `json:"warehouse_id"`
}

type InboundOrdersReport struct {
	ID                 int    `json:"id"`
	CardNumberID       string `json:"card_number_id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	WarehouseID        int    `json:"warehouse_id"`
	InboundOrdersCount int    `json:"inbound_orders_count"`
}

type InboundOrdersResponseId struct {
	Data InboundOrders `json:"data"`
}

type InboundOrdersReportResponse struct {
	Data InboundOrdersReport `json:"data"`
}
