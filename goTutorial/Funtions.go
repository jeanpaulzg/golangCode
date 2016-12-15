package main

import "fmt"

type Salutation struct {
    name string
    greeting string
}

type Othersample struct{
    name string
    msg string
}

func Greet(salutation Salutation) {
    fmt.Println(salutation.name)
    fmt.Println(salutation.greeting)
    }
    
func Greet2(salutation Salutation) {
    fmt.Println(CreateMsg(salutation.name,salutation.greeting))
}

/*Funtions can retrurn*/
func CreateMsg(name, greeting string) string{
    
    return greeting + " " + name
}
    
 /*Funtions returning 2 values*/   
 
 func Greet3(salutation Salutation) {
    msg,alt := CreateMsg1(salutation.name,salutation.greeting)
    fmt.Println(msg)
    fmt.Println(alt)
}
func CreateMsg1(name, greeting string) (string,string){
    
    return greeting + " " + name, "HEY!" +  name
}
    
func main(){
    var a = Salutation{"Paul","hola"}
    
    Greet(a)
    Greet2(a)
    Greet3(a)
}
