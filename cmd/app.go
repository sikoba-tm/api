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
	db := repository.InitConnection(os.Getenv("POSTGRES_URL"))

	port := getPort()
	app := fiber.New()

	bencanaRepository := repository.NewBencanaRepository(db)
	bencanaService := service.NewBencanaService(bencanaRepository)
	bencanaHandler := handler.NewBencanaHandler(bencanaService)

	poskoRepository := repository.NewPoskoRepository(db)
	poskoService := service.NewPoskoService(poskoRepository)
	poskoHandler := handler.NewPoskoHandler(poskoService)

	korbanRepository := repository.NewKorbanRepository(db)
	korbanService := service.NewKorbanService(korbanRepository, poskoRepository)
	korbanHandler := handler.NewKorbanHandler(korbanService)

	bencana := app.Group("/bencana")
	bencana.Get("", bencanaHandler.GetAll)
	bencana.Post("", bencanaHandler.Create)
	bencana.Get("/:id_bencana", bencanaHandler.GetById)
	bencana.Put("/:id_bencana", bencanaHandler.UpdateById)

	posko := app.Group("/bencana/:id_bencana/posko")
	posko.Get("", poskoHandler.GetAll)
	posko.Post("", poskoHandler.Create)
	posko.Get("/:id_posko", poskoHandler.GetById)
	posko.Post("/:id_posko", korbanHandler.Create)

	korban := app.Group("/bencana/:id_bencana/korban")
	korban.Get("", korbanHandler.GetAll)
	korban.Get("/:id_korban", korbanHandler.GetById)

	log.Fatal(app.Listen(port))

}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
