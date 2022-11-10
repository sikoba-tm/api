package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sikoba-tm/api/core/domain"
	"github.com/sikoba-tm/api/core/service"
	"net/http"
)

type korbanHandler struct {
	service service.KorbanService
}

func NewKorbanHandler(service service.KorbanService) *korbanHandler {
	return &korbanHandler{service: service}
}

func (h *korbanHandler) GetAll(c *fiber.Ctx) error {
	idBencana := c.Params("id_bencana")
	result := h.service.FindAll(c.Context(), idBencana)

	return c.Status(http.StatusOK).JSON(result)
}

func (h *korbanHandler) GetById(c *fiber.Ctx) error {
	idKorban := c.Params("id_korban")
	result, err := h.service.FindById(c.Context(), idKorban)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(ObjectNotFound)
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h *korbanHandler) Create(c *fiber.Ctx) error {
	idBencana := c.Params("id_bencana")
	idPosko := c.Params("id_posko")
	var korban domain.Korban
	if err := c.BodyParser(&korban); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	created, err := h.service.Create(c.Context(), idBencana, idPosko, korban)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(created)
}

func (h *korbanHandler) UpdateById(c *fiber.Ctx) error {
	idKorban := c.Params("id_korban")
	var korban domain.Korban
	if err := c.BodyParser(&korban); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	created, err := h.service.Update(c.Context(), idKorban, korban)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(created)
}

func (h *korbanHandler) DeleteById(c *fiber.Ctx) error {
	idKorban := c.Params("id_korban")

	err := h.service.Delete(c.Context(), idKorban)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(http.StatusNoContent)
}
