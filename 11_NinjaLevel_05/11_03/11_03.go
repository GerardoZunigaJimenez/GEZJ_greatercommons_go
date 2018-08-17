package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	a5 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}

	r28 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white & black",
		},
		fourWheel: false,
	}

	fmt.Printf("Sedan:\n\tDoors: %v\n\tColor: %v\n\tLuxury Car: %v\n", a5.doors, a5.color, a5.luxury)
	fmt.Printf("Truck:\n\tDoors: %v\n\tColor: %v\n\tFour Wheele Truck: %v\n", r28.doors, r28.color, r28.fourWheel)
}
