package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Currency struct {
	Result float64 `json:"result"`
}

func changeCurrency(from string, to string, amount float64) float64 {
	uri := "https://api.currencylayer.com/convert?access_key=f40ae3edc86968a7dee03bca0ffc7043&from=" + from + "&to=" + to + "&amount=" + fmt.Sprintf("%f", amount)

	resp, err := http.Get(uri)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	respData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	var objeto Currency
	json.Unmarshal(respData, &objeto)
	return objeto.Result
}
