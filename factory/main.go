package main

import (
	"fmt"

	"factory/products"
)

func main() {
	factory := products.Product{}

	product := factory.New()

	fmt.Println("my product was created at", product.CreatedAt.UTC())
}
