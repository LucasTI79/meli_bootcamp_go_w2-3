package domain

type Carry struct {
	ID           int    `json:"id"`
	Cid          string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address      string `json:"address"`
	Telephone    string `json:"telephone"`
	LocalityId  int    `json:"locality_id"`
}

type CarrieResponse struct {
	Data []Carry `json:"data"`
}

type CarrieResponseId struct {
	Data Carry `json:"data"`
}
