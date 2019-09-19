package Adapter

import "fmt"

type RiverWater int
type PureWater float32

// river
type River struct {
}

// irrigate
func (r *River) Irrigate() RiverWater {
	fmt.Println("Irrigating")
	return 1
}

// get some water from river
func (r *River) Little() RiverWater {
	fmt.Println("Get water from river")
	return 1
}

// get some pure water to drink
func Distill(r RiverWater) PureWater {
	fmt.Println("Distill river water")
	return PureWater(r)
}

type Cup struct {
}

// get some water from cup
func (c *Cup) Little() PureWater {
	fmt.Println("Get water from cup")
	return 1.1
}

func Drink(i interface{}) {
	switch v := i.(type) {
	case Cup:
		v.Little()
		fmt.Println("Drinking")
	case River:
		r := v.Little()
		Distill(r)
		fmt.Println("Drinking")
	default:
		fmt.Println("Water is not able to drink")
	}
}

func Test() {
	c := Cup{}
	Drink(c)
	r := River{}
	Drink(r)
}

/*
river water -> irrigate
pure water -> drink
river water -> distill -> pure water -> drink
*/
