package services

import (
	"api-produtos/internal/core/domain"
	"api-produtos/internal/core/ports"
	"errors"
	"fmt"
)

type productTypeServiceImpl struct {
	repo ports.ProductTypeRepository
}

func NewProductTypeService(repo ports.ProductTypeRepository) ports.ProductTypeService {
	return &productTypeServiceImpl{
		repo: repo,
	}
}

func (s *productTypeServiceImpl) CreateType(newType domain.ProductType) ([]domain.ProductType, error) {
	if s.repo.CheckExistence(newType.Name) {
		return nil, errors.New(fmt.Sprintf("%stype name %s already exists", errorMessage, newType.Name))
	}

	t, err := s.repo.CreateType(newType)
	if err != nil {
		errLog := fmt.Sprintf("%s repository error creating type %s%s", errorMessage, errorStr, err.Error())
		log.Error(errLog)
		return nil, errors.New(errLog)
	}

	return t, nil
}

func (s *productTypeServiceImpl) GetType(id int) (*domain.ProductType, error) {
	t, err := s.repo.GetType(id)
	if err != nil {
		errLog := fmt.Sprintf("%s repository error getting type with id: %d %s %s", errorMessage, id, errorStr, err.Error())
		log.Error(errLog)
		return nil, errors.New(errLog)
	}

	return t, nil
}

func (s *productTypeServiceImpl) GetAllTypes() ([]domain.ProductType, error) {
	ts, err := s.repo.GetAllTypes()
	if err != nil {
		errLog := fmt.Sprintf("%s repository error getting all types %s %s", errorMessage, errorStr, err.Error())
		log.Error(errLog)
		return nil, errors.New(errLog)
	}

	return ts, nil
}

func (s *productTypeServiceImpl) UpdateType(id int, update domain.ProductType) (*domain.ProductType, error) {
	t, err := s.repo.UpdateType(id, update)
	if err != nil {
		errLog := fmt.Sprintf("%s repository error updating type with id: %d %s %s", errorMessage, id, errorStr, err.Error())
		log.Error(errLog)
		return nil, errors.New(errLog)
	}

	return t, nil
}

func (s *productTypeServiceImpl) DeleteType(id int, active bool) error {
	err := s.repo.DeleteType(id, active)
	if err != nil {
		errLog := fmt.Sprintf("%s repository error (logic) deleting type with id: %d %s %s", errorMessage, id, errorStr, err.Error())
		log.Error(errLog)
		return errors.New(errLog)
	}

	return nil
}
