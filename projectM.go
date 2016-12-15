package main

import(
    "fmt"
    //"strconv"
    )
    
type project interface{
    estimatedHrs()
    completionHrs()
    }
    
type projectD struct{
    name string
    stakeholders []string
    hours int
    workDays int
    }

func (pName projectD) pN() string{
    return pName.name
    }
    
func (sHolders projectD) sH() string{
    return "Hola"
    }
    
func (calc projectD) estimatedHrs() int{
    return calc.hours*calc.workDays
}

func main(){
    var names = []string{"leto", "paul", "teg","uno mas", "one more"}
    pr:= projectD {name:"roland",stakeholders:names}
    fmt.Println(pr.name)
    fmt.Println(pr.pN())
    fmt.Println(len(pr.stakeholders))
   
    for i:=0; i<= len(pr.stakeholders)-1;i++{
        fmt.Println(pr.stakeholders[i])
        
    }
    
    est:= projectD{hours:2,workDays:20}
    
    fmt.Println(est.estimatedHrs())
    
    }




