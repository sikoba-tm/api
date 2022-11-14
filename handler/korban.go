package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/sikoba-tm/api/core/domain"
	"github.com/sikoba-tm/api/core/service"
	"github.com/sikoba-tm/api/core/service/external"
	"github.com/sikoba-tm/api/utils"
	"net/http"
)

type korbanHandler struct {
	service service.KorbanService
	gcs     external.CloudStorageService
}

func NewKorbanHandler(service service.KorbanService, gcs external.CloudStorageService) *korbanHandler {
	return &korbanHandler{service: service, gcs: gcs}
}

func (h *korbanHandler) GetAll(c *fiber.Ctx) error {
	idBencana := c.Params("id_bencana")
	result := h.service.FindAll(c.Context(), idBencana)

	return c.Status(http.StatusOK).JSON(result)
}

func (h *korbanHandler) GetById(c *fiber.Ctx) error {
	paramsIdKorban := c.Params("id_korban")

	idKorban, err := utils.ParseUUIDFromString(paramsIdKorban)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(IdNotValid)
	}

	result, err := h.service.FindById(c.Context(), idKorban)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(ObjectNotFound)
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h *korbanHandler) Create(c *fiber.Ctx) error {
	ctx := context.Background()
	idBencana := c.Params("id_bencana")
	idPosko := c.Params("id_posko")

	var korban domain.Korban
	if err := c.BodyParser(&korban); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	created, err := h.service.Create(ctx, idBencana, idPosko, korban)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	foto, _ := c.FormFile("foto")
	if foto != nil {
		f, _ := foto.Open()

		objectPath := "korban/"
		objectName := created.ID.String()

		err := h.gcs.UploadFile(ctx, objectPath, objectName, f)
		if err != nil {
			return err
		}

		PUBLIC_URL := "https://storage.googleapis.com/sikoba-dev/"
		fileURL := PUBLIC_URL + objectPath + objectName
		created.Foto = fileURL
		//	Update with foto
		//UpdateFoto(ctx, created)
		created, err = h.service.Update(ctx, created.ID, *created)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return c.Status(http.StatusCreated).JSON(created)
}

func (h *korbanHandler) UpdateById(c *fiber.Ctx) error {
	ctx := context.Background()
	paramsIdKorban := c.Params("id_korban")

	idKorban, err := utils.ParseUUIDFromString(paramsIdKorban)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(IdNotValid)
	}

	var korban domain.Korban
	if err := c.BodyParser(&korban); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	foto, _ := c.FormFile("foto")
	if foto != nil {
		f, _ := foto.Open()

		objectPath := "korban/"
		objectName := paramsIdKorban

		err := h.gcs.UploadFile(ctx, objectPath, objectName, f)
		if err != nil {
			return err
		}

		PUBLIC_URL := "https://storage.googleapis.com/sikoba-dev/"
		fileURL := PUBLIC_URL + objectPath + objectName
		korban.Foto = fileURL
	}

	updated, err := h.service.Update(c.Context(), idKorban, korban)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusAccepted).JSON(updated)
}

func (h *korbanHandler) DeleteById(c *fiber.Ctx) error {
	paramsIdKorban := c.Params("id_korban")

	idKorban, err := utils.ParseUUIDFromString(paramsIdKorban)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(IdNotValid)
	}

	err = h.service.Delete(c.Context(), idKorban)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(http.StatusNoContent)
}
