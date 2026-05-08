package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func scalePoint(p *Point, factor float64) {
	p.X = int(float64(p.X) * factor)
	p.Y = int(float64(p.Y) * factor)
}

func main() {
	Point1 := Point{X: 1, Y: 2}
	Point2 := Point{X: 3, Y: 4}
	Point3 := Point{X: 5, Y: 6}

	fmt.Printf("Point1: (%d, %d)\n", Point1.X, Point1.Y)
	scalePoint(&Point1, 2.0)
	fmt.Printf("缩放后2.0 Point1: (%d, %d)\n", Point1.X, Point1.Y)

	fmt.Printf("Point2: (%d, %d)\n", Point2.X, Point2.Y)
	scalePoint(&Point2, 3.0)
	fmt.Printf("缩放后3.0 Point2: (%d, %d)\n", Point2.X, Point2.Y)

	fmt.Printf("Point3: (%d, %d)\n", Point1.X, Point1.Y)
	scalePoint(&Point3, 0.5)
	fmt.Printf("缩放后0.5 Point3: (%d, %d)\n", Point3.X, Point3.Y)
}
