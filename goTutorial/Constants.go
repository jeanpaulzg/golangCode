package main

import "fmt"

const (
   PI= 3.14
   Language = "GO"
)

const (
    A= iota
    B= iota
    C=iota
    )    
 /*Only declares one const and the rest take that assigment */   
const (
    D=iota
    E
    F
    )
    
func main(){
    
    fmt.Println(PI, Language)
    fmt.Println(A,B,C)
     fmt.Println(D,E,F)
}