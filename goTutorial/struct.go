package main

import "fmt"

type Kitchen struct{
    numOfPlates int
}

type House struct{
    Kitchen 
    numOfRooms int
}

func main(){
    h:= House{Kitchen{10}, 3}
    fmt.Println ("House h has this many rooms:", h.numOfRooms)
    fmt.Println ("House h has this many plates:", h.numOfPlates)
    fmt.Println ("The kitchen contents of this house is:", h.Kitchen)
}