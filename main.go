package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Weather struct {
	Temp struct {
	Temperature    float32 `json:"temp"`
	FeelsTemp float32 `json:"feels_like"`
    } `json:"main"`
	Wind struct {
		WindSpeed float32 `json:"speed"`
	} `json:"wind"`
}

func main() {
	var city string
	var apiUrl string
	var apiKey = "671d70a43a5245f64a966d8f8b0bb989"
	var weather Weather
	fmt.Println("Введите город: ")
	fmt.Scan(&city)
	apiUrl = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%v&units=metric&appid=%v&lang={ru}",city, apiKey)
	resp , err := http.Get(apiUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Текущая температура в городе %v : %v Градусов Цельсия, ощущается как : %v Градусов цельсия\n",city,int(weather.Temp.Temperature),int(weather.Temp.FeelsTemp))
	fmt.Printf("Текущая скорость ветра в городе %v : %v м/с\n",city,int(weather.Wind.WindSpeed))
}