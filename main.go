package main

import (
	"github.com/NiltonSousa/bolon/domain"
	"github.com/NiltonSousa/bolon/internal/config"
	"github.com/NiltonSousa/bolon/internal/controller"
	"github.com/NiltonSousa/bolon/internal/repository"
	"github.com/NiltonSousa/bolon/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := config.SQLOpenConnection()
	db.AutoMigrate(&domain.Cicle{})

	cicleRepository := repository.NewCicleRepository()
	ciclesUseCase := usecase.NewCicleUseCase(db, cicleRepository)
	ciclesController := controller.NewCicleController(ciclesUseCase)
	router.POST("/cicles", func(c *gin.Context) {
		// You may need to adapt this part based on your actual Create method's requirements
		// For example, you might extract data from gin.Context and pass it to Create
		// Here is a basic example assuming Create does not use gin.Context directly

		// If your Create method expects http.ResponseWriter and context.Context:
		err := ciclesController.Create(c.Writer, c.Request.Context())
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.Status(201)
	})

	err := router.Run("localhost:8000")

	if err != nil {
		panic(err)
	}
}
