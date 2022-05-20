package main

import "fmt"

type Circle struct {
	// 结构体嵌套
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {

	var wheel Wheel
	/*
		wheel.Circle.Point.X = 8
		wheel.Circle.Point.Y = 8
		wheel.Circle.Radius = 10
		wheel.Spokes = 20
	*/
	wheel.X = 8
	wheel.Y = 8
	wheel.Radius = 10
	wheel.Spokes = 20

	// 结构体字面量
	w1 := Wheel{
		Circle: Circle{
			Point: Point{
				X: 8,
				Y: 8,
			},
			Radius: 10,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w1)

}
