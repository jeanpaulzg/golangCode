package main

import "fmt"

type Salutation struct{
    name string
    greeting string
}

func Greet(salutation Salutation){
    msg,alt := Createmsg(salutation.name, salutation.greeting)
    fmt.Println(msg)
    fmt.Println(alt)
    }
func Createmsg(name,greeting string)(string,string){
    return greeting +" "+ name, "HEY"+ name
    
    }
func main(){
    var a = Salutation{"Paul","Hola"}
    
    Greet(a)
}