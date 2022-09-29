package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sikoba-tm/api/core/domain"
	"github.com/sikoba-tm/api/core/service"
	"net/http"
)

type bencanaHandler struct {
	service service.BencanaService
}

func NewBencanaHandler(service service.BencanaService) *bencanaHandler {
	return &bencanaHandler{service: service}
}

func (h *bencanaHandler) GetAll(c *fiber.Ctx) error {
	result := h.service.FindAll(c.Context())

	return c.Status(http.StatusOK).JSON(result)
}

func (h *bencanaHandler) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.service.FindById(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(ObjectNotFound)
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h *bencanaHandler) Create(c *fiber.Ctx) error {
	var bencana domain.Bencana
	if err := c.BodyParser(&bencana); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	created, err := h.service.Create(c.Context(), bencana)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(created)
}

func (h *bencanaHandler) UpdateById(c *fiber.Ctx) error {
	id := c.Params("id")

	var bencana domain.Bencana
	if err := c.BodyParser(&bencana); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updated, err := h.service.Update(c.Context(), id, bencana)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(updated)
}