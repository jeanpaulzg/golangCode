package main

import (
    "fmt"
// "strconv"
)


type Books interface{

    AboutB() string
    }
    
type historyd struct{
    description string
    pages int
    name string
    }
    

    
func (b historyd) AboutB() string{
    return   b.description+" "+b.name
    }
    
type fiction struct{
    historyd
    }
// type fiction struct{
//     Pages int
//     Name string
//     }

// func (f fiction) Description() string{
//     return  strconv.Itoa(f.Pages) +" "+ f.Name
//     }
    
func main() {
    d:=historyd{description:"holaaa", name:"prueba"}
    
    f:=fiction{historyd{description:"descript",name:"hh"}}
    fmt.Println(d.description)
    fmt.Println(d.name)
    fmt.Println(d.AboutB())
    
    c:= fiction{}
    c.description="hola fiction"
    fmt.Println(c.description)
    fmt.Println(f.name)
    }