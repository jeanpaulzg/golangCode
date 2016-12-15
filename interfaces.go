package main

import "fmt"

func main(){
fmt.Println("Interfaces are two things: A set of methods, \n but it is also a type")
animals := []Animal{Dog{},Cat{}}
for _,animal := range animals{
    fmt.Println(animal.Speak())
    }
 
}

type Animal interface{
    Speak() string
    }
    
type Dog struct{
    }

func (d Dog) Speak() string{
    return "Woof"
    }
    
type Cat struct{
    
    }

func (c Cat) Speak() string{
    return "meaww"
    }