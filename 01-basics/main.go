package main

import (
	"fmt"
)

func main(){
	var i int
	var f float32
	var s string
	var b bool

	ii := 1
	ff := float32(11.24)
	ss := "string"
	bb := false 

	fmt.Printf("%d, %.2f, %s, %t\n",i,f,s,b)

	fmt.Printf("%d, %.2done f, %s, %t\n",ii,ff,ss,bb)

	
	
}
