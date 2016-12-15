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
func main(){
    var a = Salutation{"Paul","hola"}
    fmt.Println( a.greeting, a.name)
    
    var s = Othersample{msg:"this is a msg for:", name:"Paul"}
    fmt.Println( s.msg,s.name)
    
}
