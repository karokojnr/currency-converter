package controllers

import (
	"fmt"
	"github.com/karokojnr/go-currency-converter/app/utils"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := utils.GetPage("")
	t := template.New("index.html")
	t, _ = t.Parse(tmpl)
	err := t.Execute(w, "")
	if err != nil {
		log.Fatal(err)
	}
}

func ConvertCurrency(w http.ResponseWriter, r *http.Request) {
	var result string
	if r.URL.Path != "/convert" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	amount := r.FormValue("amount")
	fromCurrency := r.FormValue("from_currency")
	toCurrency := r.FormValue("to_currency")

	floatAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println(err)
	}

	conversionResult := utils.Convert(floatAmount, fromCurrency, toCurrency)
	if math.Signbit(floatAmount) || floatAmount == 0 {
		result = "Cannot convert a negative or zero value!"
	} else {
		result = fmt.Sprintf("%.2f %s = %.2f %s ", floatAmount, fromCurrency, conversionResult, toCurrency)
	}
	resultContent := `<div>
					   <p>
						   <span>{{ .}}</span>
					   </p>
					   </div>`

	tmpl := utils.GetPage(resultContent)
	t := template.New("index.html")
	t, _ = t.Parse(tmpl)
	err = t.Execute(w, result)
	if err != nil {
		log.Fatal(err)
	}

}
