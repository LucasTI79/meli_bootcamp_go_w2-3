package domain

type ProductBatch struct {
	ID                 int    `json:"id"`
	BatchNumber        int    `json:"batch_number"`
	CurrentQuantity    int    `json:"current_quantity"`
	CurrentTemperature int    `json:"current_temperature"`
	DueDate            string `json:"due_date"`
	InitialQuantity    int    `json:"initial_quantity"`
	ManufacturingDate  string `json:"manufacturing_date"`
	ManufacturingHour  string `json:"manufacturing_hour"`
	MinimumTemperature int    `json:"minimum_temperature"`
	ProductID          int    `json:"product_id"`
	SectionID          int    `json:"section_id"`
}
