package main

import "fmt"

type Salutation struct{
    name string
    greeting string
}

func Greet(salutation Salutation){
    msg,alt := Createmsg(salutation.name, salutation.greeting,"yo")
    fmt.Println(msg)
    fmt.Println(alt)
    }
func Createmsg(name string,greeting ...string)(message string,alternate string){
    message= greeting[1] +" "+ name
    alternate="HEY"+ name
    
    return
    
    }
func main(){
    var a = Salutation{"Paul","Hola"}
    
    Greet(a)
}