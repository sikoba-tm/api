package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sikoba-tm/api/core/service"
	"github.com/sikoba-tm/api/handler"
	"github.com/sikoba-tm/api/repository"
	"log"
	"os"
)

func Run() {
	db := repository.Init(os.Getenv("POSTGRES_URL"))

	port := getPort()
	app := fiber.New()

	bencanaRepository := repository.NewBencanaRepository(db)
	bencanaService := service.NewBencanaService(bencanaRepository)
	bencanaHandler := handler.NewBencanaHandler(bencanaService)

	bencana := app.Group("/bencana")
	bencana.Get("", bencanaHandler.GetAll)
	bencana.Post("", bencanaHandler.Create)
	bencana.Get("/:id", bencanaHandler.GetById)
	bencana.Put("/:id", bencanaHandler.UpdateById)

	log.Fatal(app.Listen(port))

}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
