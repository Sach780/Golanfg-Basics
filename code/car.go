package main

import "fmt"
const usixteenbitmax float64 = 65535
const kmph float64 = 1.532

type car struct{
	pedal uint16
	brake uint16
	fuel uint16
	speed float64
}

func (c car) kmh() float64 {
	return float64 (c.pedal) * (c.speed/usixteenbitmax)

}

func (c car) mmh() float64 {
	return float64 (c.pedal) * (c.speed/usixteenbitmax/kmph)

}

func (c *car) newspeed(newspeed float64) {
	c.speed = newspeed
}

func main() {
	a_car := car{pedal: 65000,
			brake: 0,
			fuel: 12356,
			speed: 225}
	fmt.Println(a_car.pedal,a_car.speed)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mmh())
	a_car.newspeed(500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mmh())
}
