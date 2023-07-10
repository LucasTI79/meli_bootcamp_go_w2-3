package domain

type Warehouse struct {
	ID                 int    `json:"id"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	WarehouseCode      string `json:"warehouse_code"`
	MinimumCapacity    int    `json:"minimum_capacity"`
	MinimumTemperature int    `json:"minimum_temperature"`
	LocalityId         int    `json:"locality_id"`
}

type WarehouseResponse struct {
	Data []Warehouse `json:"data"`
}

type WarehouseResponseId struct {
	Data Warehouse `json:"data"`
}
