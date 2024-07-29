package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/macreai/todo-app-clean-architecture/internal/domain"
	"github.com/macreai/todo-app-clean-architecture/internal/usecase"
	"github.com/macreai/todo-app-clean-architecture/pkg/auth"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(u *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	user := new(domain.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.usecase.Register(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    user.Username,
		"message": "Register Success!",
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, errLogin := h.usecase.Login(req.Username, req.Password)

	if errLogin != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errLogin.Error(),
		})
	}

	token, errToken := auth.GenerateJWT(user.ID, user.Username)

	if errToken != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errToken.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token":   token,
		"message": "Login Success!",
	})
}
