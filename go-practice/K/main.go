package main

import "fmt"

type Shape interface {
	name() string
	area() float64
	perimeter() float64
}
type Ractangle struct {
	width  float64
	height float64
}

func (r Ractangle) name() string {
	return "Ractangle"
}

func (r Ractangle) area() float64 {
	return r.height * r.width
}

func (r Ractangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func printShapeStats(s Shape) {
	fmt.Println(s.name())
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

type Namer interface {
	GetName() string
	SetName(name string)
}
type Person struct {
	name string
}

func (p Person) GetName() string {
	return p.name
}
func (p *Person) SetName(name string) {
	p.name = name
}

func testName(p Namer) {
	fmt.Println(p.GetName())
	p.SetName("k")
	fmt.Println(p.GetName())
}

func main() {
	me := &Person{"J"}
	testName(me)
}
