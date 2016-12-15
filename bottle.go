package main

import (
    "fmt"
     "strconv"
    )


type BottleWork interface{
    Fill()
    Pour()
    Gallons()
    Material()
    Use()
    }
    
    
type getWater struct{
    }
    
func (f getWater) Fill() string{
    return "get water"
    }
    
type pourWater struct{
    
    }

func (p pourWater) Pour() string{
    return "Pour water"
    }
    
type bottleInfo struct{
    ounces float64
    material string
    uses string
    }
    
func (g bottleInfo) Gallons() float64{
    return g.ounces*0.0078125
    }
func (m bottleInfo) Material() string{
    return "Material is: " + m.material
    }
    
func (u bottleInfo) Use() string{
    return "Use:"+ u.uses
    }
    
type vaso struct{
    bottleInfo
    size string
    }

    
func main (){
    b1:= getWater{}
    fmt.Println(b1.Fill())
    
    b2:= pourWater{}
    fmt.Println(b2.Pour())
    b3:= bottleInfo{8,"plastic", "sport"}
    fmt.Println(b3.Gallons())
    fmt.Println(b3.Material()+" "+ b3.Use())
    var s string="medium"
    v:= vaso{bottleInfo{5,"glass","dinner"},s}
    fmt.Println("ounces: "+strconv.FormatFloat(v.ounces, 'f', -1, 64)+" material: "+v.material+" uses: "+v.uses+" size: "+v.size)
    fmt.Printf("\n The vase material is  %s.\n and is used in  %s\n and its size is %s\n which has this many ounces %f", v.material, v.uses, v.size,v.ounces)
  
}