package handler

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sikoba-tm/api/core/domain"
	"github.com/sikoba-tm/api/core/service"
)

type bencanaHandler struct {
	service service.BencanaService
}

func NewBencanaHandler(service service.BencanaService) *bencanaHandler {
	return &bencanaHandler{service: service}
}

func (h *bencanaHandler) GetAll(c *fiber.Ctx) error {
	ctx := context.Background()
	result := h.service.FindAll(ctx)

	return c.Status(http.StatusOK).JSON(result)
}

func (h *bencanaHandler) GetById(c *fiber.Ctx) error {
	ctx := context.Background()
	idBencana := c.Params("id_bencana")
	result, err := h.service.FindById(ctx, idBencana)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(ObjectNotFound)
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h *bencanaHandler) Create(c *fiber.Ctx) error {
	ctx := context.Background()
	var bencana domain.Bencana
	if err := c.BodyParser(&bencana); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	created, err := h.service.Create(ctx, bencana)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(created)
}

func (h *bencanaHandler) UpdateById(c *fiber.Ctx) error {
	ctx := context.Background()
	idBencana := c.Params("id_bencana")

	var bencana domain.Bencana
	if err := c.BodyParser(&bencana); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updated, err := h.service.Update(ctx, idBencana, bencana)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(updated)
}

func (h *bencanaHandler) DeleteById(c *fiber.Ctx) error {
	ctx := context.Background()
	idBencana := c.Params("id_bencana")

	err := h.service.Delete(ctx, idBencana)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(http.StatusNoContent)
}
