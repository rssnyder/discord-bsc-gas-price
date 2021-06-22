package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

const (
	GasURL = "https://bscgas.info/gas"
)

type GasPrices struct {
	Slow     int `json:"slow"`
	Standard int `json:"standard"`
	Fast     int `json:"fast"`
	Instant  int `json:"instant"`
}

func GetGasPrices() (GasPrices, error) {

	var prices GasPrices

	req, err := http.NewRequest("GET", GasURL, nil)
	if err != nil {
		return prices, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0")
	req.Header.Add("accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return prices, err
	}

	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return prices, err
	}

	err = json.Unmarshal(results, &prices)
	if err != nil {
		fmt.Printf(resp.Status)
		return prices, err
	}

	return prices, nil
}
