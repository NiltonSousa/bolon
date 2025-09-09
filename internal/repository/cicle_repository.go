package repository

import (
	"github.com/NiltonSousa/bolon/domain"
	"gorm.io/gorm"
)

type CicleRepository struct {
	Repository[domain.Cicle]
}

func NewCicleRepository() *CicleRepository {
	return &CicleRepository{}
}

func (r *CicleRepository) FindById(tx *gorm.DB, id string) (domain.Cicle, error) {
	var cicles domain.Cicle

	result := tx.Where("id = ?", id).Find(&cicles)

	return cicles, result.Error
}
