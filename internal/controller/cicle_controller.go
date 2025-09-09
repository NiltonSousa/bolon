package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NiltonSousa/bolon/internal/model"
	"github.com/NiltonSousa/bolon/internal/usecase"
)

type CicleController struct {
	UseCase *usecase.CicleUseCase
}

func NewCicleController(useCase *usecase.CicleUseCase) *CicleController {
	return &CicleController{
		UseCase: useCase,
	}
}

func (c *CicleController) Create(w http.ResponseWriter, ctx context.Context) error {
	//TODO: implement authmiddleware

	//TODO: implement validation request

	request := model.CicleCreateRequest{
		Name:     "teste",
		IsActive: true,
	}

	response, err := c.UseCase.Create(ctx, &request)

	if err != nil {
		fmt.Println("failed to create address")
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	return nil
}
