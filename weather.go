package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Gradus struct {
	temp float64 `json:"temp"`
}

type Vlagnost struct {
	humidity float64 `json:"humidity"`
}

func (cp* Vlagnost) Vlagnost () float64   {
	fmt.Println("Vlagnost :" cp.humidity)
	return cp.humidity
}

func (cp* Gradus) gradus () float64   {
	fmt.Println("Gradusy :" cp.gradus)
	return cp.gradus
}


func.main () {
	url := "https://api.openweathermap.org/data/2.5/weather?id=1496747&appid=39967919ab50999a0969c0942b3bc0b5"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body)
	if err != nil {
		fmt.Println(err)
	}

	Result := Gradus
	fmt.Println("Gradusy" Result)

	Vla := Vlagnost
	fmt.Println("vlagnost vozduha" Vlagnost)

}
