package main

import (
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"-", radius}
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Println(err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println(area)
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f : %s", e.radius, e.err)
}
