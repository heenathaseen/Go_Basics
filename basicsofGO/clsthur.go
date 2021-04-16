package main

import (
    "fmt"
    "math"
)

type Circle struct {
    x, y, r float64
}

type Rectangle struct {
    l, b float64
}

type Shape interface {
    area() float64
}

func totalShapeArea(shapes ...Shape) (total float64) {
    for _,currentShape := range shapes {
        total += currentShape.area()
    }
    return
}

func totalArea(c []Circle, r []Rectangle) (total float64) {
    for _, a := range c {
        total += a.area()
    }
    for _, a := range r {
        total += a.area()
    }
    return
}

/*func totalArea(c Circle, r Rectangle) float64 {
    return c.area() + r.area()
}*/

func (c Circle) area() float64 {
    return math.Pi * c.r * c.r
}

func (r Rectangle) area() float64 {
    return r.l * r.b
}

func circleArea(c Circle) float64 {
    return math.Pi * c.r * c.r
}

func rectancleArea(r Rectangle) float64 {
    return r.l * r.b
}

func main() {
    c1 := Circle{15, 15, 8}
    fmt.Println(c1, c1.area())

    r1 := Rectangle{35, 40}
    fmt.Println(r1, r1.area())

    //fmt.Println(totalArea(c1, r1))
    fmt.Println(totalArea([]Circle{c1, Circle{10, 10, 6}}, []Rectangle{r1}))
    fmt.Println(totalShapeArea(c1, r1, Circle{10, 10, 6}))
}

func main2() {
    r1 := Rectangle{120, 40}

    fmt.Println(r1)
    fmt.Println(rectancleArea(r1))
}

func main1() {
    c1 := Circle{60, 60, 20}

    fmt.Println(c1)
    fmt.Println(circleArea(c1))
}