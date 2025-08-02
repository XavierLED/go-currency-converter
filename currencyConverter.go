package main

import (
	"encoding/json"
	"net/http"
	"github.com/charmbracelet/huh"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	currency1   string
	currency2	string
	amount		string
)

type jsonData struct{
	Rates []string `json:"rates"`
}

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
	resp, err := http.Get("https://openexchangerates.org/api/latest.json?app_id=" + os.Getenv("API_KEY"))

	data, err := ioutil.ReadAll(resp.Body)
	
	conversion1, conversion2 := DecodeJson(data)

	intAmount, err := strconv.ParseFloat(amount, 64)
	newAmount := intAmount * (conversion2 / conversion1)
	fmt.Printf("Your new amount is: %.2f\n", newAmount)
}

func DecodeJson(data []byte) (float64, float64){
	var newData map[string]interface{}	
	json.Unmarshal(data, &newData)
	rates := newData["rates"].(map[string]interface{})

	conversion1 := rates[currency1]
	conversion2 := rates[currency2]

	return conversion1.(float64), conversion2.(float64)
}
