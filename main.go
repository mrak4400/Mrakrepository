package Mrakrepository

import (
	"fmt"
)

type car struct {
	body, wheel, salon, color       string
	cost, age, owners, Enginevolume int32
	Shit                            bool
}
type salon struct {
	salcolor, salmaterial string
}
type cost struct {
	basecomp, maxcpomp int32
}

func main() {
	var Moskvich car
	Moskvich.body = "соетская сталь"
	Moskvich.wheel = "леворукая"
	Moskvich.salon = salon{salcolor: "Красный", salmaterial: "кожа"}
	Moskvich.age = 15
	Moskvich.owners = 4
	Moskvich.Enginevolume = 2.0
	Moskvich.Shit = false
	Moskvich.cost = cost{basecomp: 12000, maxcpomp: 25000}

	var Porshe car
	Porshe.body = "Алюмишка"
	Porshe.salon = salon{salcolor: "Черный", salmaterial: "кожа"}
	Porshe.wheel = "Леврорукая"
	Porshe.color = "черная"
	Porshe.cost = cost{basecomp: 125000, maxcpomp: 25000}
	Porshe.age = 5
	Porshe.owners = 2
	Porshe.Enginevolume = 3.0
	Porshe.Shit = false

	var Result string
	Result = "Зачем переплачивать, если и то и то говно?"

	fmt.Printf(Moskvich)
	fmt.Printf(Porshe)
	fmt.Printf(Result)
}
