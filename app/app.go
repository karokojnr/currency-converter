package app

import (
	"github.com/joho/godotenv"
	"github.com/karokojnr/go-currency-converter/app/routes"
)

func Run() {
	godotenv.Load()
	routes.Run()
}
