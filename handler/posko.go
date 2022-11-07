package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sikoba-tm/api/core/domain"
	"github.com/sikoba-tm/api/core/service"
	"net/http"
)

type poskoHandler struct {
	service service.PoskoService
}

func NewPoskoHandler(service service.PoskoService) *poskoHandler {
	return &poskoHandler{service: service}
}

func (h *poskoHandler) GetAll(c *fiber.Ctx) error {
	id_bencana := c.Params("id_bencana")
	result := h.service.FindAll(c.Context(), id_bencana)

	return c.Status(http.StatusOK).JSON(result)
}

func (h *poskoHandler) GetById(c *fiber.Ctx) error {
	id_bencana := c.Params("id_bencana")
	id_posko := c.Params("id_posko")
	result, err := h.service.FindById(c.Context(), id_bencana, id_posko)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(ObjectNotFound)
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h *poskoHandler) Create(c *fiber.Ctx) error {
	id_bencana := c.Params("id_bencana")
	var posko domain.Posko
	if err := c.BodyParser(&posko); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	created, err := h.service.Create(c.Context(), id_bencana, posko)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(created)
}
