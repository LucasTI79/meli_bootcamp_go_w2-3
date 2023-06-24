package section_test

import (
	"context"
	"errors"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/section"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/section"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAll(t *testing.T) {
	t.Run("should return a list of sections", func(t *testing.T) {
		expectedSections := []domain.Section{
			{
				ID:                 1,
				SectionNumber:      1,
				CurrentTemperature: 10,
				MinimumTemperature: 5,
				CurrentCapacity:    10,
				MinimumCapacity:    5,
				MaximumCapacity:    20,
				WarehouseID:        1,
				ProductTypeID:      1,
			},
			{
				ID:                 2,
				SectionNumber:      2,
				CurrentTemperature: 10,
				MinimumTemperature: 5,
				CurrentCapacity:    10,
				MinimumCapacity:    5,
				MaximumCapacity:    20,
				WarehouseID:        1,
				ProductTypeID:      1,
			},
		}
		mockRepository, service := InitServerWithWarehousesRepository(t)
		mockRepository.On("GetAll").Return(expectedSections, nil)
		sections, err := service.GetAll(context.TODO())
		assert.True(t, len(sections) == 2)
		assert.NoError(t, err)
	})

	t.Run("should not return a list of sections", func(t *testing.T) {
		mockRepository, service := InitServerWithWarehousesRepository(t)
		mockRepository.On("GetAll").Return([]domain.Section{}, errors.New("error"))
		sections, err := service.GetAll(context.Background())
		assert.True(t, len(sections) == 0)
		assert.Error(t, err)
	})
}

func InitServerWithWarehousesRepository(t *testing.T) (*mocks.SectionRepositoryMock, section.Service) {
	t.Helper()
	mockRepository := &mocks.SectionRepositoryMock{}
	mockService := section.NewService(mockRepository)
	return mockRepository, mockService
}
