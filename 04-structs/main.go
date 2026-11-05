package main

import "fmt"

type Rectangle struct {
	Width, Height float32
}

func (r *Rectangle) Area () float32 {
	return r.Height * r.Width
}

func (r *Rectangle) Scale (factor float32) {
	r.Width = r.Width * factor
	r.Height = r.Height * factor
}



func main (){
	r := Rectangle{3,2}
	fmt.Println(r.Area())

	r.Scale(2)

	fmt.Println(r.Area())

	r2 := NewRectangle(5,4)
	fmt.Println(r2.Area())
	
}

func NewRectangle(w, h float32) *Rectangle {
	return &Rectangle{Width: w, Height: h}
}
