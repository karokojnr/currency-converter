package utils

import (
	"fmt"
	"math"
	"strconv"
)

func GetCurrencyRates() map[string]map[string]float64 {

	ngn := GoDotEnvVariable("NIGERIAN_NAIRA")
	ghs := GoDotEnvVariable("GHANAIAN_CEDIS")
	ksh := GoDotEnvVariable("KENYAN_SHILLINGS")

	ngnToGhs, err := strconv.ParseFloat(GoDotEnvVariable("NGN_TO_GHS"), 64)
	ngnToKsh, err := strconv.ParseFloat(GoDotEnvVariable("NGN_TO_KSH"), 64)

	ghsToNgn, err := strconv.ParseFloat(GoDotEnvVariable("GHS_TO_NGN"), 64)
	ghsToKhs, err := strconv.ParseFloat(GoDotEnvVariable("GHS_TO_KSH"), 64)

	kshToNgn, err := strconv.ParseFloat(GoDotEnvVariable("KSH_TO_NGN"), 64)
	kshToGhs, err := strconv.ParseFloat(GoDotEnvVariable("KSK_TO_GHS"), 64)
	if err != nil {
		fmt.Println(err)
	}

	currencyRates := map[string]map[string]float64{
		ngn: {
			ghs: ngnToGhs,
			ksh: ngnToKsh,
		},
		ghs: {
			ngn: ghsToNgn,
			ksh: ghsToKhs,
		},
		ksh: {
			ngn: kshToNgn,
			ghs: kshToGhs,
		},
	}
	return currencyRates
}

func Convert(amount float64, fromCurrency string, toCurrency string) float64 {
	currencyRates := GetCurrencyRates()
	if fromCurrency == toCurrency {
		result := amount
		return math.Floor(result*100) / 100
	} else {
		result := amount * (currencyRates[fromCurrency][toCurrency])
		return math.Floor(result*100) / 100
	}
}
