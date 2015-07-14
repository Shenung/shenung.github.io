package Converter

import "fmt"

const (
	f2km = 0.0003048
)

func FeetToKm(x int) {
	fmt.Println(x*f2km, "km")
}
