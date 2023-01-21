package main

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	jq := gojsonq.New().File("./sample-data.json")
	res := jq.From("vendor.items").Where("price", ">", 1200).OrWhere("id", "=", nil).Get()
	fmt.Println(res)
	// output: [map[price:1350 id:1 name:MacBook Pro 13 inch retina] map[id:2 name:MacBook Pro 15 inch retina price:1700] map[id:<nil> name:HP core i3 SSD price:850]]
}
