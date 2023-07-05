package domain

type Carry struct {
	ID          int    `json:"id"`
	Cid         string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityId  int    `json:"locality_id"`
}

type LocalityCarriersReport struct {
	LocalityID    int    `json:"locality_id"`
	LocalityName  string `json:"locality_name"`
	CarriersCount int    `json:"carriers_count"`
}

type CarrieResponse struct {
	Data []Carry `json:"data"`
}

type CarrieResponseId struct {
	Data Carry `json:"data"`
}
