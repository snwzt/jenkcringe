package main

import (
	"fmt"
	"simple_factory/factory"
)

func main() {
	airbus, err := factory.NewPlane("AirbusA380")
	if err != nil {
		fmt.Println(err)
		return
	}
	airbus.SetSpeed(900)
	fmt.Printf("Airbus A380 Speed: %f\n", airbus.GetSpeed())

	f35, err := factory.NewPlane("F35")
	if err != nil {
		fmt.Println(err)
		return
	}
	f35.SetSpeed(1200)
	fmt.Printf("F35 Speed: %f\n", f35.GetSpeed())

	_, err = factory.NewPlane("UnknownPlane")
	if err != nil {
		fmt.Println(err)
	}
}
