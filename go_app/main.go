// Created by: Ryan Duffett
// Created on: Sep 2020
//
// This program finds the area of a triangle

package main

import "fmt"

func main() {
	// This function finds the area and perimeter of a rectangle
	var base int
	var height int
	var area int

	// input
	fmt.Println("This program finds the area of a triangle.")
	fmt.Println()
	fmt.Print("Enter the base (mm): ")
	fmt.Scanln(&base)
	fmt.Print("Enter the height (mm): ")
	fmt.Scanln(&height)

	// process
	area = base * height / 2

	// output
	fmt.Println()
	fmt.Println("The area is:", area, "mm².")
	fmt.Println("\nDone.")
}