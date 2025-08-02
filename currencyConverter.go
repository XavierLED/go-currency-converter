package main

import (
	"net/http"
	"github.com/charmbracelet/huh"
	//"encoding/json"
	"fmt"
	"strconv"
)

var (
	currency1   string
	currency2	string
	amount		string
)

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose Currency youre converting from").
				Options(//use currency as a filler term
					huh.NewOption("CAD", "CAD"),
					huh.NewOption("USD", "USD"),
					huh.NewOption("JPY", "JPY"),
					huh.NewOption("EUR", "EUR"),
				).
				Value(&currency1),

			huh.NewSelect[string]().
				Title("Choose Currency youre coverting to").
				Options(
					huh.NewOption("CAD", "CAD"),
					huh.NewOption("USD", "USD"),
					huh.NewOption("JPY", "JPY"),
					huh.NewOption("EUR", "EUR"),
				).
				Value(&currency2),

			huh.NewInput().
				Title("How much do you want to convert example: 100.25").
				Value(&amount),
		),
	)
	err := form.Run()
	if err != nil {
		fmt.Println(err)
	}

	startAmount, err := strconv.ParseFloat(amount, 32)
	fmt.Println(startAmount)
	//get api request from api
	resp, err := http.Get("https://openexchangerates.org/api/latest.json?app_id=" + apikey thingy)
	resp.Body.Close()
	fmt.Println(resp)
	//find first currency amount unless usd and divide it by starting amount
	//find second amount and multipy it by starting amount to get final amount
	//or just do amount * (rate to/ rate from)
	//return answer
}
