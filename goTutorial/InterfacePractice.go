package main
import "fmt"

type Animal struct{
    name string
    specie string
}

type Birds struct{
    Animal
    region string
}

type mamiferos struct{
    Animal
    Zone string
}

func (B Birds) SayHi(){
    fmt.Printf("Hi I'm a %s and belong to %s specie and live in the %s region \n", B.name, B.specie, B.region)
}

func (M mamiferos) SayHi(){
    fmt.Printf("Hi I'm a %s and belong to %s specie and live in the %s area \n ", M.name, M.specie, M.Zone)
}

type akingdom interface{
    SayHi()
}

func main(){
    birdy:= Birds{Animal{"canario","amarillo"},"Sur"}
    perro:= mamiferos{Animal{"pastor","canino"},"caliente"}
    
    var i akingdom
    
    i= birdy
    i.SayHi()
    
    i= perro
    i.SayHi()
    
    
    
}