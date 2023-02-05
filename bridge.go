package main

import "fmt"

type Shape interface {
	Draw()
}

// Circle is a concrete implementation of Shape
type Circle struct {
	x, y, radius int
	drawAPI      DrawAPI
}

func (c *Circle) Draw() {
	c.drawAPI.DrawCircle(c.radius, c.x, c.y)
}

// DrawAPI is the abstraction for the low-level drawing API
type DrawAPI interface {
	DrawCircle(radius, x, y int)
}

// RedCircle is a concrete implementation of DrawAPI
type RedCircle struct{}

func (r *RedCircle) DrawCircle(radius, x, y int) {
	fmt.Println("Drawing Circle[ color: red, radius:", radius, "x:", x, "y:", y, "]")
}

// GreenCircle is a concrete implementation of DrawAPI
type GreenCircle struct{}

func (g *GreenCircle) DrawCircle(radius, x, y int) {
	fmt.Println("Drawing Circle[ color: green, radius:", radius, "x:", x, "y:", y, "]")
}

func main() {
	redCircle := Circle{x: 100, y: 100, radius: 10, drawAPI: &RedCircle{}}
	redCircle.Draw()

	greenCircle := Circle{x: 200, y: 200, radius: 20, drawAPI: &GreenCircle{}}
	greenCircle.Draw()
}

//Шаблон проектирования «Мост» — это структурный шаблон, который отделяет интерфейс объекта от его реализации,
//позволяя им изменяться независимо друг от друга.
//Этот шаблон позволяет отделить объекты и повышает гибкость, позволяя переключаться между различными реализациями абстракции, не затрагивая клиентский код.

//Интерфейс Shape действует как абстракция, а Circle — как его конкретная реализация.
//Интерфейс DrawAPI действует как низкоуровневая абстракция, а RedCircle и GreenCircle являются его конкретными реализациями.
//Используя шаблон Bridge, мы можем разделить абстракции Shape и DrawAPI и изменить низкоуровневый API рисования, не влияя на реализацию Circle.
