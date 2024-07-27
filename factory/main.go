package main

import (
	"factory/products"
	"fmt"
)

func main() {
	factory := products.Product{}

	product := factory.New()

	fmt.Println("my product was created at", product.CreatedAt.UTC())
}
