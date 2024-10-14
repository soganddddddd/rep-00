package main

import (
	"fmt"
	"math"
)

func main() {

	var area, radius float32
	const pi = 3.14
	fmt.Printf("input radius in meter \n")
	fmt.Scanf("%f", &radius)

	area = math.Pi * radius * radius

	fmt.Printf("radius of %f meters : area is %f square meter \n", radius, area)

}
