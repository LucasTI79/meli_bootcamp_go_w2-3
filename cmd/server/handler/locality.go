package handler

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/locality"
)

type LocalityController struct {
	localityService locality.Service
}

func NewLocality(l locality.Service) *LocalityController {
	return &LocalityController{
		localityService: l,
	}
}

