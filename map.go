package main

import (
	"fmt"
)

func main() {
	population := map[string]int{
		"Islamabad":  600000,
		"Taxila":     400000,
		"HasanAbdal": 250000,
		"Karachi":    980000,
		"Wah Cant":   150000,
	}
	// population["Lahore"] = 850000
	// delete(population, "Islamabad")

	// pop, ok := population["Taxila"] //to check whether key is present or not
	// fmt.Println(pop, ok)            // ok return false if not present otherwise true and with the value

	// comma ok ???????
	// ok := population["Taxila"] //to check whether key is present or not
	// fmt.Println(ok)            // ok return false if not present otherwise true and with the value

	//fmt.Println(len(population)) //lenth fuction to find length of map

	x := population
	delete(x, "Wah Cant")
	fmt.Println(x)
	fmt.Println(population)

}
