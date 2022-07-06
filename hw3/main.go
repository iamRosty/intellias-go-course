package main

import "fmt"

type kg float64

const (
	dogFoodPerKgWeight kg = 2.
	catFoodPerKgWeight kg = 7.
	cowFoodPerKgWeight kg = 25.
)

type Animal interface {
	getFoodWeight() kg
	getInfo()
}

type Dog struct {
	name               string
	weight, foodWeight kg
}

func (d *Dog) getFoodWeight() kg {
	return d.foodWeight
}

func (d *Dog) getInfo() {
	fmt.Printf("Name: %v, Weight: %v kg, Need %v kg food.\n", d.name, d.weight, d.foodWeight)
}

func newDog(name string, weight kg) *Dog {
	return &Dog{name, weight, weight * dogFoodPerKgWeight}
}

type Cat struct {
	name               string
	weight, foodWeight kg
}

func (c *Cat) getFoodWeight() kg {
	return c.foodWeight
}

func (c *Cat) getInfo() {
	fmt.Printf("Name: %v, Weight: %v kg, Need %v kg food.\n", c.name, c.weight, c.foodWeight)
}

func newCat(name string, weight kg) *Cat {
	return &Cat{name, weight, weight * catFoodPerKgWeight}
}

type Cow struct {
	name               string
	weight, foodWeight kg
}

func (c *Cow) getFoodWeight() kg {
	return c.foodWeight
}

func (c *Cow) getInfo() {
	fmt.Printf("Name: %v, Weight: %v kg, Need %v kg food.\n", c.name, c.weight, c.foodWeight)
}

func newCow(name string, weight kg) *Cow {
	return &Cow{name, weight, weight * cowFoodPerKgWeight}
}

func myFunction(a ...Animal) kg {
	var totalFoodweight kg
	for _, v := range a {
		v.getInfo()
		totalFoodweight += v.getFoodWeight()
	}

	return totalFoodweight
}
func main() {
	c := newCat("Tom", 7)
	c1 := newCat("Lisa", 5)
	d := newDog("Pluto", 15)
	t := myFunction(c, c1, d)
	fmt.Println("Total food weight(kg):", t)
}
