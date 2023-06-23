package domain

type Seller struct {
	ID          int    `json:"id"`
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

type SellerResponse struct {
	Data []Seller `json:"data"`
}
type SellerResponseId struct {
	Data Seller `json:"data"`
}