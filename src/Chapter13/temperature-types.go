package main

import "fmt"

type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func main() {
	var k kelvin = 394.0
	c1 := kelvinToCelsius(k)
	fmt.Println(k, "Kは、", c1, "Cです")

	c2 := k.celsius()
	fmt.Println(k, "Kは、", c2, "Cです")
}
