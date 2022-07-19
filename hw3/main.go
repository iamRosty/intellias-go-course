package main

import "fmt"

type kg float64

const (
	dogFoodPerKgWeight kg = 2.
	catFoodPerKgWeight kg = 7.
	cowFoodPerKgWeight kg = 25.
)

type animal interface {
	getFoodWeight() kg
}

type dog struct {
	name   string
	weight kg
}

func (d *dog) getFoodWeight() kg {
	return d.weight * dogFoodPerKgWeight
}

func newdog(name string, weight kg) *dog {
	return &dog{name, weight}
}

type cat struct {
	name   string
	weight kg
}

func (c *cat) getFoodWeight() kg {
	return c.weight * catFoodPerKgWeight
}

func newcat(name string, weight kg) *cat {
	return &cat{name, weight}
}

type cow struct {
	name   string
	weight kg
}

func (c *cow) getFoodWeight() kg {
	return c.weight * cowFoodPerKgWeight
}

func (c *cow) getInfo() {
	fmt.Printf("Name: %v, Weight: %v kg, Need %v kg food.\n", c.name, c.weight, c.weight*cowFoodPerKgWeight)
}

func newcow(name string, weight kg) *cow {
	return &cow{name, weight}
}

func farmInfo(a ...animal) kg {
	var name string
	var weight, foodWeight, totalFoodweight kg
	for _, v := range a {
		switch x := v.(type) {
		case *dog:
			name = x.name
			weight = x.weight
			foodWeight = x.weight * dogFoodPerKgWeight
		case *cat:
			name = x.name
			weight = x.weight
			foodWeight = x.weight * catFoodPerKgWeight
		case *cow:
			name = x.name
			weight = x.weight
			foodWeight = x.weight * cowFoodPerKgWeight
		}
		fmt.Printf("Name: %v, Weight: %v kg, Need %v kg food.\n", name, weight, foodWeight)
		totalFoodweight += v.getFoodWeight()
	}

	return totalFoodweight
}
func main() {
	c := newcat("Tom", 7)
	c1 := newcat("Lisa", 5)
	d := newdog("Pluto", 15)
	t := farmInfo(c, c1, d)
	fmt.Println("Total food weight(kg):", t)
}
