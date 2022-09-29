package main

import (
	"github.com/joho/godotenv"
	app "github.com/sikoba-tm/api/cmd"
)

func main() {
	godotenv.Load()
	app.Run()
}
