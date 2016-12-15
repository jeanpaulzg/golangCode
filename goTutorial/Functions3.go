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
func Createmsg(name,greeting string)(message string,alternate string){
    message= greeting +" "+ name
    alternate="HEY"+ name
    
    return
    
    }
func main(){
    var a = Salutation{"Paul","Hola"}
    
    Greet(a)
}