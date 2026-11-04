package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "";
	for i:= 1; i <= 20; i = i+1 {		
		switch {
		case  i % 3 == 0 && i % 5 != 0:
			s = "fizz"
		case i % 5 == 0 && i % 3 != 0:
			s = "buzz"
		case i % 5 == 0 && i % 3 == 0:
			s = "fizzbuzz"						
		default:
			s =strconv.Itoa(i)
		}
		fmt.Println(s)
	}

	for i:= 0;;i=i+1 {
		fmt.Println(i)
		if  i == 5 {
			break
		}
	}
}
