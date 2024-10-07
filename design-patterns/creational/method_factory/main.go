package main

import "method_factory/factory"

func main() {
	factoryProduct1 := &factory.FactoryProduct1{}
	product1, _ := factoryProduct1.CreateProduct()
	println(product1.Info())
}
