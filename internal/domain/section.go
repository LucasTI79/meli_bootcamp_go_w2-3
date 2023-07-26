package domain

import "errors"

var (
	ErrNotFound             = errors.New("section not found")
	ErrInvalidId            = errors.New("invalid id")
	ErrInvalidIdWareHouse   = errors.New("invalid warehouse_id field")
	ErrInvalidIdProductType = errors.New("invalid product_type_id field")
	ErrTryAgain             = errors.New("error, try again %s")
	ErrAlreadyExists        = errors.New("section already exists")
	ErrModifySection        = errors.New("cannot modify Section")
)

type Section struct {
	ID                 int `json:"id"`
	SectionNumber      int `json:"section_number" binding:"required"`
	CurrentTemperature int `json:"current_temperature" binding:"required"`
	MinimumTemperature int `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    int `json:"current_capacity" binding:"required"`
	MinimumCapacity    int `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    int `json:"maximum_capacity" binding:"required"`
	WarehouseID        int `json:"warehouse_id"`
	ProductTypeID      int `json:"product_type_id"`
}

type SectionsResponse struct {
	Data []Section `json:"data"`
}

type ProductBySection struct {
	ProductsCount int    `json:"products_count"`
	SectionID     int    `json:"section_id"`
	SectionNumber string `json:"section_number"`
}

type ProductBySectionResponse struct {
	Data []ProductBySection `json:"data"`
}

type SectionResponse struct {
	Data Section `json:"data"`
}

type SectionRequest struct {
	SectionNumber      int `json:"section_number" binding:"required"`
	CurrentTemperature int `json:"current_temperature" binding:"required"`
	MinimumTemperature int `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    int `json:"current_capacity" binding:"required"`
	MinimumCapacity    int `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    int `json:"maximum_capacity" binding:"required"`
	WarehouseID        int `json:"warehouse_id" binding:"required"`
	ProductTypeID      int `json:"product_type_id" binding:"required"`
}

func (s *Section) Validate() error {
	if s.WarehouseID == 0 {
		return ErrInvalidIdWareHouse
	}

	if s.ProductTypeID == 0 {
		return ErrInvalidIdProductType
	}

	return nil
}
