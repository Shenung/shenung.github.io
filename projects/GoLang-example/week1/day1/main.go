package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// import (
//   "fmt"
// ) //for declaring multiple imports

// var name string = "Shenung" //another way to make a variable

// var name = "Shenung"
const (
	miTkm = 1.60934
	miTft = 5280
	miTm  = 1609.34
	mTmi  = 0.000621371
	mTft  = 3.28084
	mTkm  = 0.001
)

// const pTkg = 0.453592

func main() {
	from := os.Args[1]
	to := os.Args[2]

	switch {
	case strings.HasSuffix(from, "mi"):
		miles, _ := strconv.ParseFloat(from[0:len(from)-2], 64)
		switch to {
		case "km":
			fmt.Println(miles * miTkm)
		case "ft":
			fmt.Println(miles * miTft)
		case "m":
			fmt.Println(miles * miTm)
		}
	case strings.HasSuffix(from, "km"):
		miles, _ := strconv.ParseFloat(from[0:len(from)-2], 64)
		switch to {
		case "km":
			fmt.Println(miles * miTkm)
		case "ft":

		case "m":
		}
	case strings.HasSuffix(from, "ft"):
		miles, _ := strconv.ParseFloat(from[0:len(from)-2], 64)
		switch to {
		case "km":
			fmt.Println(miles * miTkm)
		case "ft":

		case "m":
		}
	case strings.HasSuffix(from, "m"):
		miles, _ := strconv.ParseFloat(from[0:len(from)-2], 64)
		switch to {
		case "km":
			fmt.Println(miles * mTkm)
		case "ft":
			fmt.Println(miles * mTkm)
		case "m":
			fmt.Println(miles * mTkm)
		}
	}
	fmt.Println(from, to)

	//----------------------------------------------------------------
	// name := "Shenung"
	//
	// var name string
	// fmt.Println("Enter your name: ")
	// fmt.Scanf("%s", &name)
	// fmt.Println("Hello World my name is ", name)
	//--------------------------------------------------------
	// var miles float64
	// fmt.Print("Enter number(miles): ")
	// fmt.Scanln(&miles)
	// km := miles * miTkm
	// fmt.Println("+-----------------------+")
	// fmt.Printf("| Miles: %15.2f", miles)
	// fmt.Printf("|\n")
	// fmt.Println("+-----------------------+")
	// fmt.Printf("| Kilometers: %10.2f", km)
	// fmt.Printf("|\n")
	// fmt.Println("+-----------------------+")
	//
	// var fahrenheit float64
	// fmt.Print("Enter number(fahrenheits): ")
	// fmt.Scanln(&fahrenheit)
	// celcius := (fahrenheit - 32) * (5.0 / 9)
	// fmt.Println(celcius, "C")
	//
	// var weight float64
	// fmt.Print("Enter number(pounds): ")
	// fmt.Scanln(&weight)
	// lbs := weight * pTkg
	// fmt.Println(lbs, "kg")
	//--------------------------------------------------------
	//
	// miles, err := strconv.ParseFloat(os.Args[1], 64)
	// if err != nil {
	// 	panic(err)
	// }
	// km := miles * miTkm
	// fmt.Println("<!DOCTYPE HTML>")
	// fmt.Println("<html><head><title>GoProgram</title></head>")
	// fmt.Println("<body>")
	// fmt.Println("Miles:", miles)
	// fmt.Println("Kilometers:", km)
	// fmt.Println("</body>")
	// fmt.Println("</html>")
}
