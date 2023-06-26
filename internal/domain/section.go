package domain

import "errors"

var (
	ErrNotFound     = errors.New("section not found")
	ErrInvalidId    = errors.New("invalid id")
	ErrTryAgain     = errors.New("error, try again %s")
	ErrAlreadyExists = errors.New("section already exists")
	ErrModifySection = errors.New("cannot modify Section")
)

type Section struct {
	ID                 int `json:"id"`
	SectionNumber      int `json:"section_number"`
	CurrentTemperature int `json:"current_temperature"`
	MinimumTemperature int `json:"minimum_temperature"`
	CurrentCapacity    int `json:"current_capacity"`
	MinimumCapacity    int `json:"minimum_capacity"`
	MaximumCapacity    int `json:"maximum_capacity"`
	WarehouseID        int `json:"warehouse_id"`
	ProductTypeID      int `json:"product_type_id"`
}

type SectionsResponse struct {
	Data []Section `json:"data"`
}

type SectionResponse struct {
	Data Section `json:"data"`
}

type SectionRequest struct {
	SectionNumber      int `json:"section_number"`
	CurrentTemperature int `json:"current_temperature"`
	MinimumTemperature int `json:"minimum_temperature"`
	CurrentCapacity    int `json:"current_capacity"`
	MinimumCapacity    int `json:"minimum_capacity"`
	MaximumCapacity    int `json:"maximum_capacity"`
	WarehouseID        int `json:"warehouse_id"`
	ProductTypeID      int `json:"product_type_id"`
}