package main

import "fmt"

type BaseSupply struct {
	name string
}

type Freezer struct {
	BaseSupply
	latitude, longitude float64
}
type Beer struct {
	alcohol_percent_min, alcohol_percent_max float64
}

func put_in_freezer(beer Beer) {}

func main() {
	fmt.Println("nanani nanana")
}
