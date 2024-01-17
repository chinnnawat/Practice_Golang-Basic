package main

import (
	"fmt"
	"math"
)

// ใช้เก็บฟังก์ชัน
type geometry interface {
	// เก็บ func 2 ตัว
	area() float64
	perim() float64
}

type reactant struct {
	width, height float64
}

type circle struct {
	radius	float64
}
// **************Reactant****************
// r => var r float
// (r reactant) => คือการกำหนดค่า r ให้เข้าถึง struc ของ reactant {width, height}
// (r reactant) area() เมื่อทำงานเสร็จ จะรีเทินไปยัง area
func (r reactant) area()  float64{
	return r.width * r.width
}

func (r reactant) perim()  float64{
	return (r.width * r.width)*2
}

// ****************Circle*******************
func (c circle) perim() float64{
	// math.Pi = ค่า pi
	return 2*math.Pi*c.radius
}
func (c circle) area() float64{
	// math.Pi = ค่า pi
	return math.Pi*(c.radius*c.radius)
}

//	*****************Measure****************
func measure(g geometry){
	area := g.area()
	perim := g.perim()
	fmt.Println("Area = ", area)
	fmt.Println("Perim = ", perim)
}

func main() {
	r := reactant{width:5, height: 10}
	c := circle{radius: 15}

	measure(r)
	measure(c)
}