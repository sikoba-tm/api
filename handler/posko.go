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
	idBencana := c.Params("id_bencana")
	result := h.service.FindAll(c.Context(), idBencana)

	return c.Status(http.StatusOK).JSON(result)
}

func (h *poskoHandler) GetById(c *fiber.Ctx) error {
	idBencana := c.Params("id_bencana")
	idPosko := c.Params("id_posko")
	result, err := h.service.FindById(c.Context(), idBencana, idPosko)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(ObjectNotFound)
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h *poskoHandler) Create(c *fiber.Ctx) error {
	idBencana := c.Params("id_bencana")
	var posko domain.Posko
	if err := c.BodyParser(&posko); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	created, err := h.service.Create(c.Context(), idBencana, posko)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(created)
}

func (h *poskoHandler) UpdateById(c *fiber.Ctx) error {
	idBencana := c.Params("id_bencana")
	idPosko := c.Params("id_posko")

	var posko domain.Posko
	if err := c.BodyParser(&posko); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updated, err := h.service.Update(c.Context(), idBencana, idPosko, posko)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(updated)
}

func (h *poskoHandler) DeleteById(c *fiber.Ctx) error {
	idBencana := c.Params("id_bencana")
	idPosko := c.Params("id_posko")

	err := h.service.Delete(c.Context(), idBencana, idPosko)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(http.StatusNoContent)
}
