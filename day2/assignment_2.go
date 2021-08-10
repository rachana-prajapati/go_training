package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

type FilterCriteria struct {
	AttributeName string
	Operator      string
	Value         string
}

func main() {
	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Stationary"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	testProduct := Product{101, "Kettle", 2500, 10, "Stationary"}
	indexOfTestProduct := IndexOf(products, testProduct)
	fmt.Println(testProduct, " | indexOfTestProduct -> ", indexOfTestProduct)

	doesTestProdExist := Includes(products, testProduct)
	fmt.Println(testProduct, " | doesTestProdExist -> ", doesTestProdExist)

	stationaryProductPredicate := func(p Product) bool {
		if p.Category == "Stationary" {
			return true
		}
		return false
	}
	filteredProds := Filter(products, stationaryProductPredicate)
	fmt.Println("stationaryProds -> ", filteredProds)

	costlyProductPredicate := func(p Product) bool {
		if p.Cost > 100 {
			return true
		}
		return false
	}
	filteredProds = Filter(products, costlyProductPredicate)
	fmt.Println("costlyProducts -> ", filteredProds)

	fmt.Println()
	fmt.Println("Any function")
	// Are thery any stationary products?
	fmt.Println("Are there any stationary products ? => ", Any(products, stationaryProductPredicate))
	// Are there any costly product?
	fmt.Println("Are there any costly products ? => ", Any(products, costlyProductPredicate))

	fmt.Println()
	fmt.Println("All function")
	//Are all the products in the collection are costly products?
	fmt.Println("Are all the products in the collection are costly products ? => ", All(products, costlyProductPredicate))
	//Are all the products in the collection are stationary products?
	fmt.Println("Are all the products in the collection are stationary products ? => ", All(products, stationaryProductPredicate))
}

//IndexOf => returns the index of the given product in the products collections (-1 if it doesnt exist)
func IndexOf(collectionOfProducts []Product, prod Product) int {

	for i, val := range collectionOfProducts {
		if prod == val {
			return i
		}
	}
	return -1
}

//Includes => return true/false based on the existence of the product in the products collection
func Includes(collectionOfProducts []Product, prod Product) bool {

	return IndexOf(collectionOfProducts, prod) != -1
}

//Filter => returns a new collection of products that match the given criteria
//use cases
//filter all stationary products
//filter costly products (cost > 100)
func Filter(collectionOfProducts []Product, predicate func(Product) bool) (filteredProducts []Product) {
	for _, prod := range collectionOfProducts {
		if predicate(prod) {
			filteredProducts = append(filteredProducts, prod)
		}
	}
	return filteredProducts
}

//Any => return true/false based on the existence of the given product that satisfies the given criteria in the products collection
//Use Cases
// Are there any costly product?
// Are thery any stationary products?
func Any(collectionOfProducts []Product, predicate func(Product) bool) (isThereAnyMatchingProduct bool) {

	isThereAnyMatchingProduct = false

	for _, prod := range collectionOfProducts {
		if predicate(prod) {
			isThereAnyMatchingProduct = true
			break
		}
	}
	return isThereAnyMatchingProduct
}

//All => return true/false based on the existence of all the given products that satisfy the given criteria in the products collection
//Use Cases
//Are all the products in the collection are costly products?
//Are all the products in the collection are stationary products?
func All(collectionOfProducts []Product, predicate func(Product) bool) (doAllProdMatchTheCriteria bool) {

	doAllProdMatchTheCriteria = true

	for _, prod := range collectionOfProducts {
		if !predicate(prod) {
			doAllProdMatchTheCriteria = false
			break
		}
	}
	return doAllProdMatchTheCriteria
}
