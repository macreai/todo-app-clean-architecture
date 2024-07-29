package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/macreai/todo-app-clean-architecture/internal/domain"
	"github.com/macreai/todo-app-clean-architecture/internal/usecase"
)

type ActivityUserHandler struct {
	usecase *usecase.ActivityUserUsecase
}

func NewActivityUserHandler(u *usecase.ActivityUserUsecase) *ActivityUserHandler {
	return &ActivityUserHandler{usecase: u}
}

func (h *ActivityUserHandler) Create(c *fiber.Ctx) error {
	activityUser := new(domain.ActivityUser)

	if err := c.BodyParser(activityUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.usecase.Create(activityUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(activityUser)
}

func (h *ActivityUserHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	activityUser, err1 := h.usecase.GetByID(uint(id))
	if err1 != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err1.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": activityUser,
	})
}

func (h *ActivityUserHandler) GetAll(c *fiber.Ctx) error {
	activityUser, err := h.usecase.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&activityUser)
}

func (h *ActivityUserHandler) Update(c *fiber.Ctx) error {
	activityUser := new(domain.ActivityUser)
	if err := c.BodyParser(activityUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.usecase.Update(activityUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(activityUser)
}

func (h *ActivityUserHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.usecase.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
