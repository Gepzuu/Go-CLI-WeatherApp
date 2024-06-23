package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC    float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch    int64   `json:"time_epoch"`
				TempC        float64 `json:"temp_c"`
				Condition    struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	q := "Iasi"
	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=aa50cc9825904adb811112917242306&q=" + q + "&days=1&aqi=no&alerts=no")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
