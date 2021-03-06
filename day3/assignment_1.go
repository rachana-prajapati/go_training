package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

type Products []Product

func (products Products) Len() int { return len(products) }
func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}
func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

type ByName struct {
	Products
}

func (p ByName) Less(i, j int) bool {
	return p.Products[i].Name < p.Products[j].Name
}

func main() {
	products := Products{
		{105, "Pen", 5, 50, "Stationary"},
		{107, "Pencil", 2, 100, "Stationary"},
		{103, "Marker", 50, 20, "Utencil"},
		{102, "Stove", 5000, 5, "Utencil"},
		{101, "Kettle", 2500, 10, "Utencil"},
		{104, "Scribble Pad", 20, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(products)

	//Sort the products by any attribute
	//IMPORTANT : Do NOT implement your own sorting logic. Use "sort" package functions
	//HINT : look at sort.Sort() function
	sort.Sort(products)
	fmt.Println("Default Sorting (using id)")
	fmt.Println(products)
	fmt.Println()

	fmt.Println()
	fmt.Println("Sort by name")
	sort.Sort(ByName{products})
	fmt.Println(products)
	fmt.Println()
}

//Method of product
func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

//Method of []Product
func (products Products) String() string {
	var result string
	for _, p := range products {
		result += fmt.Sprintf("%s\n", p)
	}
	return result
}
