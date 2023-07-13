package domain

import (
	"errors"
	"time"
)

type ProductBatch struct {
	ID                 int    `json:"id"`
	BatchNumber        int    `json:"batch_number" binding:"required"`
	CurrentQuantity    int    `json:"current_quantity" binding:"required" `
	CurrentTemperature int    `json:"current_temperature" binding:"required"`
	DueDate            string `json:"due_date" binding:"required"`
	InitialQuantity    int    `json:"initial_quantity" binding:"required"`
	ManufacturingDate  string `json:"manufacturing_date" binding:"required"`
	ManufacturingHour  int    `json:"manufacturing_hour" binding:"required"`
	MinimumTemperature int    `json:"minimum_temperature" binding:"required"`
	ProductID          int    `json:"product_id" binding:"required" min:"1"`
	SectionID          int    `json:"section_id" binding:"required" min:"1"`
}

var (
	ErrInvalidManufacturingDate = errors.New("invalid manufacturing date")
)

const layout = "2006-01-02"

func (p *ProductBatch) Validate() error {
	_, err := time.Parse(layout, p.ManufacturingDate)
	if err != nil {
		return ErrInvalidManufacturingDate
	}
	_, err = time.Parse(layout, p.DueDate)
	if err != nil {
		return ErrInvalidManufacturingDate
	}
	return nil
}
