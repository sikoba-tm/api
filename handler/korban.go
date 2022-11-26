package handler

import (
	"bytes"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/sikoba-tm/api/core/domain"
	"github.com/sikoba-tm/api/core/service"
	"github.com/sikoba-tm/api/core/service/external"
	"github.com/sikoba-tm/api/utils"
	"io"
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
	ctx := context.Background()
	idBencana := c.Params("id_bencana")
	result := h.service.FindAll(ctx, idBencana)

	return c.Status(http.StatusOK).JSON(result)
}

func (h *korbanHandler) GetById(c *fiber.Ctx) error {
	ctx := context.Background()
	paramsIdKorban := c.Params("id_korban")

	idKorban, err := utils.ParseUUIDFromString(paramsIdKorban)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(IdNotValid)
	}

	result, err := h.service.FindById(ctx, idKorban)
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
		defer f.Close()

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
		defer f.Close()

		objectPath := "korban/"
		objectName := paramsIdKorban

		err := h.gcs.UploadFile(ctx, objectPath, objectName, f)
		if err != nil {
			return err
		}

		PUBLIC_URL := "https://storage.googleapis.com/sikoba-dev/"
		fileURL := PUBLIC_URL + objectPath + objectName
		korban.Foto = fileURL
	} else {
		// Not replacing with new image
		fotoURL := c.FormValue("foto_url")
		if fotoURL == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "foto can't be empty"})
		}
		korban.Foto = fotoURL
	}

	updated, err := h.service.Update(ctx, idKorban, korban)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusAccepted).JSON(updated)
}

func (h *korbanHandler) DeleteById(c *fiber.Ctx) error {
	ctx := context.Background()
	paramsIdKorban := c.Params("id_korban")

	idKorban, err := utils.ParseUUIDFromString(paramsIdKorban)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(IdNotValid)
	}

	err = h.service.Delete(ctx, idKorban)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(http.StatusNoContent)
}

func (h *korbanHandler) Search(c *fiber.Ctx) error {
	ctx := context.Background()
	idBencana := c.Params("id_bencana")
	reference, _ := c.FormFile("foto")
	if reference == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "image can't be empty"})
	}
	// Process image in bytes
	f, _ := reference.Open()
	defer f.Close()

	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, f)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	filter := make(map[string]interface{})
	namaFromForm := c.FormValue("nama")
	if namaFromForm != "" {
		filter["nama"] = namaFromForm
	}

	korbans, err := h.service.SearchByFoto(ctx, idBencana, buf.Bytes(), filter)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if len(korbans) == 0 {
		return c.Status(http.StatusNotFound).JSON(korbans)
	}
	return c.Status(http.StatusOK).JSON(korbans)
}
