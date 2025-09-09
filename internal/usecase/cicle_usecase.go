package usecase

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/NiltonSousa/bolon/domain"
	"github.com/NiltonSousa/bolon/internal/model"
	"github.com/NiltonSousa/bolon/internal/repository"
	"gorm.io/gorm"
)

type CicleUseCase struct {
	DB              *gorm.DB
	CicleRepository *repository.CicleRepository
}

func NewCicleUseCase(db *gorm.DB, cicleRepository *repository.CicleRepository) *CicleUseCase {
	return &CicleUseCase{
		DB:              db,
		CicleRepository: cicleRepository,
	}
}

func (c *CicleUseCase) Create(ctx context.Context, request *model.CicleCreateRequest) (*domain.Cicle, error) {
	tx := c.DB.WithContext(ctx).Begin()
	cicleId := rand.Int64()

	defer tx.Rollback()

	cicle := domain.NewCicle(cicleId, request.Name, request.IsActive)

	if err := c.CicleRepository.Create(tx, cicle); err != nil {
		fmt.Println("Error creating cicle")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println("Error commiting cicle")
		return nil, err
	}

	return cicle, nil
}
