package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-3/internal/domain"
	"github.com/stretchr/testify/mock"
)

type SectionServiceMock struct {
	mock.Mock
}

type SectionRepositoryMock struct {
	mock.Mock
}

func (m *SectionServiceMock) GetAll(ctx context.Context) ([]domain.Section, error) {
	args := m.Called()

	return args.Get(0).([]domain.Section), args.Error(1)
}

func (m *SectionRepositoryMock) GetAll() ([]domain.Section, error){

	args := m.Called()

	return args.Get(0).([]domain.Section), args.Error(1)
}

func (m *SectionServiceMock) Save(ctx context.Context, s domain.Section) (int, error) {
	return 0, nil
}

func (m *SectionServiceMock) Delete(ctx context.Context, id int) error {
	return nil
}

func (m *SectionServiceMock) 	Get(ctx context.Context, id int) (domain.Section, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Section), args.Error(1)
}

func (m *SectionRepositoryMock) Get(ctx context.Context, id int) (domain.Section, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Section), args.Error(1)
}

func (m *SectionServiceMock) Update(ctx context.Context, s domain.Section) error {
	return nil
}