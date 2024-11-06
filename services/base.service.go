package services

import "product-service-app/repositories"

type BaseService[T any] struct {
	Repository *repositories.BaseRepository[T]
}

func (s *BaseService[T]) Create(model *T) error {
	return s.Repository.Create(model)
}

func (s *BaseService[T]) FindByID(id uint, model *T) error {
	return s.Repository.FindByID(id, model)
}

func (s *BaseService[T]) FindByModelID(id string) (*T, error) {
	return s.Repository.FindByModelID(id)
}

func (s *BaseService[T]) Update(model *T) error {
	return s.Repository.Update(model)
}

func (s *BaseService[T]) Delete(model *T) error {
	return s.Repository.Delete(model)
}

func (s *BaseService[T]) DeleteByModelID(id string) error {
	return s.Repository.DeleteByModelID(id)
}

func (s *BaseService[T]) GetAll() ([]T, error) {
	return s.Repository.FindAll()
}
