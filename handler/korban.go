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
	id_bencana := c.Params("id_bencana")
	result := h.service.FindAll(c.Context(), id_bencana)

	return c.Status(http.StatusOK).JSON(result)
}

func (h *korbanHandler) GetById(c *fiber.Ctx) error {
	//id_bencana := c.Params("id_bencana")
	id_korban := c.Params("id_korban")
	result, err := h.service.FindById(c.Context(), id_korban)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(ObjectNotFound)
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h *korbanHandler) Create(c *fiber.Ctx) error {
	id_bencana := c.Params("id_bencana")
	id_posko := c.Params("id_posko")
	var korban domain.Korban
	if err := c.BodyParser(&korban); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	created, err := h.service.Create(c.Context(), id_bencana, id_posko, korban)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(created)
}
