package main

import("fmt")

const usixteenbitmax float64 = 65553
const kmh_multiple float64 = 1.6042

type car struct{
	gas_pedal uint16
	break_pedal uint16
	streeing_wheel int16
	top_speed_kmp float64
}

func (c car) kmh() float64{
	return float64(c.gas_pedal)* (c.top_speed_kmp/usixteenbitmax)
}
func (c car) mph() float64{
	return float64(c.gas_pedal)* (c.top_speed_kmp/usixteenbitmax/kmh_multiple)
}

func (c* car) new_top_speed(newspeed float64){
	c.top_speed_kmp=newspeed
}
func main(){
	aCar := car{   //car object
		gas_pedal:22314,
		break_pedal:0,
		streeing_wheel:12561,
		top_speed_kmp:225.0}
		fmt.Println(aCar.kmh())
		fmt.Println(aCar.mph())
		aCar.new_top_speed(344)
		fmt.Println(aCar.kmh())
		fmt.Println(aCar.mph())
	}
