package section_test

import (
	"context"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/section"
	mocks "github.com/extmatperez/meli_bootcamp_go_w2-3/tests/section"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestGetById(t *testing.T) {
	t.Run("should return a section", func(t *testing.T) {
		expectedSection := domain.Section{
			ID:                 1,
			SectionNumber:      1,
			CurrentTemperature: 10,
			MinimumTemperature: 5,
			CurrentCapacity:    10,
			MinimumCapacity:    5,
			MaximumCapacity:    20,
			WarehouseID:        1,
			ProductTypeID:      1,
		}
		mockRepository, service := InitServerWithWarehousesRepository(t)
		mockRepository.On("Get", 1).Return(expectedSection, nil)
		section, err := service.Get(context.Background(), 1)
		assert.Equal(t, expectedSection, section)
		assert.NoError(t, err)
	})

	t.Run("should not return a section", func(t *testing.T) {
		mockRepository, service := InitServerWithWarehousesRepository(t)
		mockRepository.On("Get", 1).Return(domain.Section{}, errors.New("error"))
		section, err := service.Get(context.Background(), 1)
		assert.Equal(t, domain.Section{}, section)
		assert.Error(t, err)
	})
}

func TestSave(t *testing.T) {
	t.Run("should save a section", func(t *testing.T) {
		expectedSection := domain.Section{
			SectionNumber:      1,
			CurrentTemperature: 10,
			MinimumTemperature: 5,
			CurrentCapacity:    10,
			MinimumCapacity:    5,
			MaximumCapacity:    20,
			WarehouseID:        1,
			ProductTypeID:      1,
		}
		mockRepository, service := InitServerWithWarehousesRepository(t)
		mockRepository.On("Exists", context.TODO(), 0).Return(true)
		mockRepository.On("Save", mock.AnythingOfType("domain.Section")).Return(1, nil)
		_, err := service.Save(context.Background(), expectedSection)
		expectedSection.ID = 1
		assert.Equal(t, 1, expectedSection.ID)
		assert.NoError(t, err)
	})

	t.Run("should not save a section on any error", func(t *testing.T) {
		section := domain.Section{
			SectionNumber:      1,
			CurrentTemperature: 10,
			MinimumTemperature: 5,
			CurrentCapacity:    10,
			MinimumCapacity:    5,
			MaximumCapacity:    20,
			WarehouseID:        1,
			ProductTypeID:      1,
		}
		mockRepository, service := InitServerWithWarehousesRepository(t)
		mockRepository.On("Exists", context.TODO(), 0).Return(true)
		mockRepository.On("Save", section).Return(0, errors.New("error"))
		id, err := service.Save(context.Background(), section)
		assert.Equal(t, 0, id)
		assert.Error(t, err)
	})

	t.Run("should not save a section if it already exists", func(t *testing.T) {
		section := domain.Section{
			ID:                 1,
			SectionNumber:      1,
			CurrentTemperature: 10,
			MinimumTemperature: 5,
			CurrentCapacity:    10,
			MinimumCapacity:    5,
			MaximumCapacity:    20,
			WarehouseID:        1,
			ProductTypeID:      1,
		}
		mockRepository, service := InitServerWithWarehousesRepository(t)
		mockRepository.On("Exists", context.TODO(), 1).Return(false)
		mockRepository.On("Save", mock.Anything).Return(false)
		id, err := service.Save(context.Background(), section)
		assert.Equal(t, 0, id)
		assert.Error(t, err)
	})
}

// Tests cases for delete service section

func TestDelete(t *testing.T) {
	t.Run("should delete a section", func(t *testing.T) {
		mockRepository, service := InitServerWithWarehousesRepository(t)
		mockRepository.On("Delete", 1).Return(nil)
		err := service.Delete(context.Background(), 1)
		assert.NoError(t, err)
	})

	t.Run("should not delete a section", func(t *testing.T) {
		mockRepository, service := InitServerWithWarehousesRepository(t)
		mockRepository.On("Delete", 1).Return(errors.New("error"))
		err := service.Delete(context.Background(), 1)

		assert.Error(t, err)
	})
}

func InitServerWithWarehousesRepository(t *testing.T) (*mocks.SectionRepositoryMock, section.Service) {
	t.Helper()
	mockRepository := &mocks.SectionRepositoryMock{}
	mockService := section.NewService(mockRepository)
	return mockRepository, mockService
}
