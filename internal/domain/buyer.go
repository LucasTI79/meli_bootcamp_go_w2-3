package domain

type Buyer struct {
	ID           int    `json:"id"`
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type BuyerRequest struct {
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type BuyerOrders struct {
	ID                  int    `json:"id"`
	CardNumberID        string `json:"card_number_id"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	PurchaseOrdersCount int    `json:"purchase_orders_count"`
}

type BuyerResponse struct {
	Data []Buyer `json:"data"`
}

type BuyerResponseID struct {
	Data Buyer `json:"data"`
}

type BuyerOrdersResponseID struct {
	Data BuyerOrders `json:"data"`
}

type BuyerOrdersResponse struct {
	Data []BuyerOrders `json:"data"`
}
