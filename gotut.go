package main

import ("fmt";
		//"net/http"
	)
type car struct{
	gas_pedal uint16
	brake_pedal uint16
}

func (c car) kmh() float64{
	return float64 (c.brake_pedal*2)
}

func main(){
	a_car := car{
		gas_pedal:696,
		brake_pedal:0}
	fmt.Println(a_car.kmh())
}