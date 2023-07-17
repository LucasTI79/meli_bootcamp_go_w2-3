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

type CarryResponse struct {
	Data []Carry `json:"data"`
}

type CarryResponseId struct {
	Data Carry `json:"data"`
}

type LocalityCarriersResponse struct {
	Data []LocalityCarriersReport `json:"data"`
}
