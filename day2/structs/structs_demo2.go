package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Units    int
	Cost     float64
	Category string
}

/*
type PerishableProduct struct {
	Prod   Product
	Expiry string
}
*/

type PerishableProduct struct {
	Product
	Expiry string
}

//utility method to construct an instance of the struct
func NewPerishableProduct(id int, name string, units int, cost float64, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{Product{id, name, units, cost, category}, expiry}
}

func main() {

	/*
		grape := PerishableProduct{Prod: Product{201, "Green Grape", 60, 100, "Food"}, Expiry: "2 days"}
		fmt.Println(grape)
		fmt.Println("Cost of product -> ", grape.Prod.Cost)
	*/

	// grape := PerishableProduct{Product{201, "Green Grape", 60, 100, "Food"}, "2 days"}
	grape := PerishableProduct{Product: Product{201, "Green Grape", 60, 100, "Food"}, Expiry: "2 days"}
	fmt.Println("Cost of product -> ", grape.Cost)

	mango := NewPerishableProduct(202, "Badam Mango", 200, 50, "Food", "5 days")
	fmt.Println(mango)
	fmt.Println("Cost of mango --> ", mango.Cost)
}
