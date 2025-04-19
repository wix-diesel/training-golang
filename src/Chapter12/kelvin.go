package main

import "fmt"

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "Kは、", celsius, "Cです。")
}

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}
