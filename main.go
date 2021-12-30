package main

import "fmt"

var x int
var y string
var z bool

func main(){
	//atv1()
	atv2()
}

func atv1(){
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, " ", y, " ", z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func atv2(){
	fmt.Println(x, y, z)
}

func atv3(){
	
}