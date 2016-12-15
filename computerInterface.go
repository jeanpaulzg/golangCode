package main

import "fmt"


type Computer interface{
    Print()
    Display()
    Server()
    Navigate()
    }
    
type mac struct{
    }
    
func (p mac) Print() string{
    return "mac is printing"
    }
    
func (d mac) Display() string{
    return "mac is Displaying"
    }
    
type pc struct{
    print string
    }
    
func (p pc) Print() string{
    return p.print
    }

func (p pc) Display() string{
    return "pc is displaying"
    }

func main(){
    compu:= mac{}
    fmt.Println(compu.Print())
    fmt.Println(compu.Display())
    
    compu1:= pc{"Printing"}
    fmt.Println(compu1.Print())
    // fmt.Println(compu1.Display())
    }
    
    
