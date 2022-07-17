package main

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

type Weather struct {
    Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int     `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
    } `json:"location"`
    Current struct {
	TempC     float64 `json:"temp_c"`
	Condition struct {
	    Text string `json:"text"`
	    Icon string `json:"icon"`
	} `json:"condition"`
	PressureMb float64 `json:"pressure_mb"`
	Humidity   int     `json:"humidity"`
	Cloud      int     `json:"cloud"`
	FeelslikeC float64 `json:"feelslike_c"`
	Uv         float64 `json:"uv"`
    } `json:"current"`
}

func main() {
    // Here your endpoint with your own API (weather api)
    url := ""
    res, err := http.Get(url)
    if err != nil {
	log.Fatal(err)
    }

    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
	log.Fatal(err)
    }

    var weather Weather
    json.Unmarshal(body, &weather)
    fmt.Printf("%s City \n", weather.Location.Name)
    fmt.Printf("Temp: %g °C \n", weather.Current.TempC)
    fmt.Printf("Region: %s \n", weather.Location.TzID)
    fmt.Printf("Feels like %g °C \n", weather.Current.FeelslikeC)
}
