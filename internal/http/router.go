package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/macreai/todo-app-clean-architecture/internal/http/handler"
	"github.com/macreai/todo-app-clean-architecture/internal/repository/postgres"
	"github.com/macreai/todo-app-clean-architecture/internal/usecase"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *fiber.App {
	app := fiber.New()

	activityRepo := postgres.NewPostgresActivityUserRepository(db)
	activityUsecase := usecase.NewActivityUserUsecase(activityRepo)
	activityHandler := handler.NewActivityUserHandler(activityUsecase)

	app.Post("/activity", activityHandler.Create)
	app.Get("/activity/:id", activityHandler.GetByID)
	app.Get("/activity/", activityHandler.GetAll)
	app.Put("/activity/:id", activityHandler.Update)
	app.Delete("/activity/:id", activityHandler.Delete)

	return app
}
