package repositories

import (
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func (r *BaseRepository[T]) Create(model *T) error {
	return r.DB.Create(model).Error
}

func (r *BaseRepository[T]) FindByID(id uint, model *T) error {
	return r.DB.First(model, id).Error
}

// FindByID finds a record by its ID
func (r *BaseRepository[T]) FindByModelID(id string) (*T, error) {
	var entity T
	if err := r.DB.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete removes a record by its ID
func (r *BaseRepository[T]) DeleteByModelID(id string) error {
	if err := r.DB.Delete(new(T), id).Error; err != nil {
		return err
	}
	return nil
}

func (r *BaseRepository[T]) Update(model *T) error {
	return r.DB.Save(model).Error
}

func (r *BaseRepository[T]) Delete(model *T) error {
	return r.DB.Delete(model).Error
}

func (r *BaseRepository[T]) FindAll() ([]T, error) {
	var entities []T
	err := r.DB.Find(&entities).Error
	return entities, err
}
