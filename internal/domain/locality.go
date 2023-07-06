package domain

type Locality struct {
	ID           int    `json:"id"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
}

type LocalityInput struct {
	ID           int    `json:"id"`
	LocalityName string `json:"locality_name"`
	IdProvince int `json:"id_province"`
}

type LocalityReport struct {
	IdLocality int `json:"id_locality"`
	LocalityName string `json:"locality_name"`
	SellersCount int `json:"sellers_count"`
}

