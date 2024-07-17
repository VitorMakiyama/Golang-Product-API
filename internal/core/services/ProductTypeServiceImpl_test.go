package services

import (
	"api-produtos/internal/core/domain"
	"github.com/stretchr/testify/mock"
	"slices"
	"testing"
)

func getTypes() []domain.ProductType {
	return []domain.ProductType{{
		Id:     0,
		Name:   "T1",
		Active: false,
	},
		{
			Id:     1,
			Name:   "T2",
			Active: true,
		},
		{
			Id:     2,
			Name:   "T3",
			Active: true,
		},
	}
}

func getCreatedType(t domain.ProductType) []domain.ProductType {
	var types = getTypes()
	return append(types, t)
}

type MockedRepository struct {
	mock.Mock
}

func (m *MockedRepository) CreateType(newType domain.ProductType) ([]domain.ProductType, error) {
	args := m.Called()
	return args.Get(0).([]domain.ProductType), args.Error(1)
}

func (m *MockedRepository) GetType(id int) (*domain.ProductType, error) {
	args := m.Called()
	return args.Get(0).(*domain.ProductType), args.Error(1)
}

func (m *MockedRepository) GetAllTypes() ([]domain.ProductType, error) {
	args := m.Called()
	return args.Get(0).([]domain.ProductType), args.Error(1)
}

func (m *MockedRepository) UpdateType(id int, update domain.ProductType) (*domain.ProductType, error) {
	args := m.Called()
	return args.Get(0).(*domain.ProductType), args.Error(1)
}

func (m *MockedRepository) DeleteType(id int, active bool) error {
	args := m.Called()
	return args.Error(1)
}

func (m *MockedRepository) CheckExistence(name string) bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockedRepository) ValidateType(id int) bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockedRepository) GetTypeByName(name string) (*domain.ProductType, error) {
	args := m.Called()
	return args.Get(0).(*domain.ProductType), args.Error(1)
}

func TestShouldGetAllTypes(t *testing.T) {
	mockedRepo := new(MockedRepository)
	mockedRepo.On("GetAllTypes").Return(getTypes(), nil)

	service := NewProductTypeService(mockedRepo)

	ts, err := service.GetAllTypes()
	if err != nil {
		t.Fatalf(`Error inside GetAllTypes (repo) - %v`, err) // we can use "" inside ``
	}
	if !slices.Equal(getTypes(), ts) {
		t.Fatalf(`Returned types are wrong! - %v`, ts)
	}

	mockedRepo.AssertExpectations(t)
}

func TestCreateType_ShouldCreate(t *testing.T) {
	mockedRepo := new(MockedRepository)
	newT := domain.ProductType{
		Id:   3,
		Name: "T4",
	}
	mockedRepo.On("CreateType").Return(getCreatedType(newT), nil)
	mockedRepo.On("CheckExistence").Return(false)

	service := NewProductTypeService(mockedRepo)

	ts, err := service.CreateType(newT)
	if err != nil {
		t.Fatalf(`Error inside CreateTypes (repo) - %v`, err)
	}
	if len(ts) == 3 {
		t.Fatalf(`Returned len is wrong! - %v`, len(ts))
	}
	if ts[3] != newT {
		t.Fatalf(`Created type was not returned - %v e %v`, ts[3], newT)
	}
	mockedRepo.AssertExpectations(t)
}
