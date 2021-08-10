package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Units    int
	Cost     float64
	Category string
}

func main() {
	//var pen Product = Product{}
	//var pen Product = Product{100, "OCTAGEL PEN", 100, 5, "Stationary"}
	//var pen Product = Product{Id: 100, Name: "OCTAGEL PEN", Units: 100, Cost: 5, Category: "Stationary"}
	pen := Product{Id: 100, Name: "OCTAGEL PEN", Units: 100, Cost: 5, Category: "Stationary"}
	fmt.Println(pen)

	//the dot operator
	fmt.Printf("Product %s costs %v\n", pen.Name, pen.Cost)

	applyDiscount(&pen) //apply a discount of 10% on the product
	fmt.Println("After applying 10% discount, pen --> ", pen)
}

func applyDiscount(penPtr *Product) {
	//(*penPtr).Cost = (*penPtr).Cost - (*penPtr).Cost*0.10
	//fmt.Println("func -> ", (*penPtr).Cost)
	penPtr.Cost = penPtr.Cost - penPtr.Cost*0.10
	//fmt.Println("func -> ", (penPtr).Cost)
}
