package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Main struct {
	temp     float64 `json:"temp"`
	humidity float64 `json:"humidity"`
}

func (cp *Main) grads() float64 {
	fmt.Println("Gradusy :", cp.temp)
	return cp.temp
}
func (cp *Main) hum() float64 {
	fmt.Println("Vlagnost :", cp.humidity)
	return cp.humidity
}

func main() {
	url := "https://api.openweathermap.org/data/2.5/weather?id=1496747&appid=39967919ab50999a0969c0942b3bc0b5"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	result := &Main{}

	err = json.Unmarshal(body, result)
	if err != nil {
		fmt.Println(err)
	}

	r := result.grads()
	r2 := result.hum()
	fmt.Println(r, r2)

}
