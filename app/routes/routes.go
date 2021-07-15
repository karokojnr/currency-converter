package routes

import (
	"fmt"
	"github.com/karokojnr/go-currency-converter/app/controllers"
	"github.com/karokojnr/go-currency-converter/app/utils"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", controllers.GetIndex)
	http.HandleFunc("/convert", controllers.ConvertCurrency)

	port := utils.GetPort()
	fmt.Println("App running on port  ✓  :", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
func Run() {
	HandleRequests()
}
