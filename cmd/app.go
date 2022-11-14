package cmd

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sikoba-tm/api/core/service"
	"github.com/sikoba-tm/api/core/service/external"
	"github.com/sikoba-tm/api/handler"
	"github.com/sikoba-tm/api/repository"
	"log"
	"os"
)

func Run() {
	db := repository.InitConnection(os.Getenv("POSTGRES_URL"))

	port := getPort()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("storage.NewClient: %v\n", err)
	}
	defer client.Close()

	//mySess := session.Must(session.NewSession())
	//rkg := rekognition.New(mySess, aws.NewConfig().WithRegion("ap-southeast-1"))
	//rekognitionSvc := external.NewRekognitionClient(rkg)

	bencanaRepository := repository.NewBencanaRepository(db)
	bencanaService := service.NewBencanaService(bencanaRepository)
	bencanaHandler := handler.NewBencanaHandler(bencanaService)

	poskoRepository := repository.NewPoskoRepository(db)
	poskoService := service.NewPoskoService(poskoRepository)
	poskoHandler := handler.NewPoskoHandler(poskoService)

	cloudStorageService := external.NewGoogleCloudStorage(client)

	korbanRepository := repository.NewKorbanRepository(db)
	korbanService := service.NewKorbanService(korbanRepository, poskoRepository)
	korbanHandler := handler.NewKorbanHandler(korbanService, cloudStorageService)

	bencana := app.Group("/bencana")
	bencana.Get("", bencanaHandler.GetAll)
	bencana.Post("", bencanaHandler.Create)
	bencana.Get("/:id_bencana", bencanaHandler.GetById)
	bencana.Put("/:id_bencana", bencanaHandler.UpdateById)
	bencana.Delete("/:id_bencana", bencanaHandler.DeleteById)

	posko := app.Group("/bencana/:id_bencana/posko")
	posko.Get("", poskoHandler.GetAll)
	posko.Post("", poskoHandler.Create)
	posko.Get("/:id_posko", poskoHandler.GetById)
	posko.Put("/:id_posko", poskoHandler.UpdateById)
	posko.Delete("/:id_posko", poskoHandler.DeleteById)
	posko.Post("/:id_posko", korbanHandler.Create)

	korban := app.Group("/bencana/:id_bencana/korban")
	korban.Get("", korbanHandler.GetAll)
	korban.Get("/:id_korban", korbanHandler.GetById)
	korban.Put("/:id_korban", korbanHandler.UpdateById)
	korban.Delete("/:id_korban", korbanHandler.DeleteById)

	log.Fatal(app.Listen(port))

}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
