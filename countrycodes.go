package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
{
	"location": "kuala lumpur",
	"weather": "sunny"
	"temparature": 30,
	"celsius": true,
	"temp_forecast": [25, 28, 24, 28, 30, 30, 29],
	"wind": {
		"direction": "NW",
		"speed": 120

	}
}
//which can be automatically converted to json in the website called jsontogo
type AutoGenerated struct {
	Location     string `json:"location"`
	Weather      string `json:"weather"`
	Temparature  int    `json:"temparature"`
	Celsius      bool   `json:"celsius"`
	TempForecast []int  `json:"temp_forecast"`
	Wind         struct {
		Direction string `json:"direction"`
		Speed     int    `json:"speed"`
	} `json:"wind"`
}

*/

type Country struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type weatherInfo struct {
	LocationName        string    `json: "locationName"`
	Weather             string    `json: "weather"`
	Temparature         float32   `json: "temperature"`
	TemparatureForecast []float32 `json: "temparatureForecast"`
	Celsius             bool      `json: "celsius"`
	Wind                windInfo  `json: "wind"`
}

type windInfo struct {
	Direction string  `json: "direction"`
	Speed     float32 `json: "speed"`
}

type locationInfo struct {
	Lon float32 `json: "lon"`
	Lat float32 `json: "lat"`
}

func getWeatherInfo(response http.ResponseWriter, request *http.Request) {
	location := locationInfo{}
	jsn, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Fatal("Error reading the request body!", err)
	}

	err = json.Unmarshal(jsn, &location)

	if err != nil {
		log.Fatal("decoding error! something wrong!", err)
	}
	log.Print("Received : %v\n", location)

	weather1 := weatherInfo{
		LocationName:        "seri kembangan",
		Weather:             "More or Less Cloudy",
		Temparature:         28.23,
		Celsius:             true,
		TemparatureForecast: []float32{23.3, 33.3, 34.456, 343},
		Wind: windInfo{
			Direction: "SE",
			Speed:     23.3,
		},
	}
	weather2 := weatherInfo{
		LocationName:        "seri serdang",
		Weather:             "More or Less Cloudy",
		Temparature:         28.23,
		Celsius:             true,
		TemparatureForecast: []float32{23.3, 33.3, 34.456, 343},
		Wind: windInfo{
			Direction: "SE",
			Speed:     23.3,
		},
	}
	weather3 := weatherInfo{
		LocationName:        "puchong",
		Weather:             "More or Less Cloudy",
		Temparature:         28.23,
		Celsius:             true,
		TemparatureForecast: []float32{23.3, 33.3, 34.456, 343},
		Wind: windInfo{
			Direction: "SE",
			Speed:     23.3,
		},
	}

	weatherSlice := []weatherInfo{weather1, weather2, weather3}

	weatherToBeServed, err := json.Marshal(weatherSlice)

	if err != nil {
		fmt.Fprintf(response, "Error: %s", err)
		log.Fatal("Error", err)

	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(weatherToBeServed)
}

func server() {
	http.HandleFunc("/", setCountryInfo)
	http.ListenAndServe(":8080", nil)
}

func setCountryInfo(w http.ResponseWriter, r *http.Request) {

	country := Country{}
	jsn, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("Error reading the request body!", err)
	}

	err = json.Unmarshal(jsn, &country)

	if err != nil {
		log.Fatal("decoding error! something wrong!", err)
	}
	log.Println("Received : %v", country)
}

func client() {
	//locJson, err := json.Marshal(locationInfo{Latitude:45.44, Longititude: -55.66})
	locJSON, err := json.Marshal(Country{
		Code:  "008",
		Name:  "Bangladesh",
		Phone: "0176961080",
	})

	req, err := http.NewRequest("POST", "http://10.0.2.15:8080", bytes.NewBuffer(locJSON))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Response: ", string(body))
	resp.Body.Close()
}

func main() {
	server()
	client()
}
