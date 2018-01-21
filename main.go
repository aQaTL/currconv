package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type CurrencyData struct {
	Currencies map[string]Currency `json:"results"`
}

type Currency struct {
	Name string `json:"currencyName"`
}

func loadCurrencies() CurrencyData {
	var data CurrencyData

	// List of all currencies from http://free.currencyconverterapi.com/api/v3/currencies
	file, _ := Asset("data/currencies.json")
	json.Unmarshal(file, &data)

	return data
}

func getCurrencyNames(from string, to string) (string, string) {
	data := loadCurrencies()
	fromCurrencyName := data.Currencies[from].Name
	toCurrencyName := data.Currencies[to].Name

	return fromCurrencyName, toCurrencyName
}

func isValidCurrency(currencyID string) bool {
	data := loadCurrencies()
	currencyName := data.Currencies[currencyID].Name

	return currencyName != ""
}

func handleError(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func getRate(rateID string) float64 {
	url := "http://free.currencyconverterapi.com/api/v3/convert?q=" + rateID + "&compact=ultra"

	response, err := http.Get(url)
	if err != nil {
		handleError("Error getting data")
	}
	defer response.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		handleError("Error fetching data")
	}

	rate := data[rateID]
	if rate == nil {
		handleError("No results for currency rate " + rateID)
	}

	return rate.(float64)
}

func validArgs(args []string) (float64, string, string) {
	if len(args) < 4 {
		handleError("Insufficient arguments")
	}
	args = args[1:] //Remove program name

	from, to := strings.ToUpper(args[1]), strings.ToUpper(args[2])
	if !isValidCurrency(from) || !isValidCurrency(to) {
		handleError("Invalid currency code")
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		handleError("Invalid value for conversion")
	}

	return amount, from, to
}

func main() {
	amount, from, to := validArgs(os.Args)

	result := amount * getRate(from+"_"+to)
	fromCurrency, toCurrency := getCurrencyNames(from, to)

	fmt.Println(amount, fromCurrency, "=", result, toCurrency)
}
