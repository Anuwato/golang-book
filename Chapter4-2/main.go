package main

import "fmt"

func main () {
	fmt.Println("Enter a number: ")
	var F float64
	fmt.Scanf("%f", &F)
	//output := input * 2
	C := (F-32) * 5/9
	fmt.Printf("Celcious: %.2f\n",C)
}