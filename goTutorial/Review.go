package main

import ("fmt"
        "strconv"
)


type car interface{
    start() 
    speeds()
    slows()
    stops()
    }

type car1 struct{
    name string
    description string
    price int
    quantity int
    }

func (c1 car1) start() string{
    compute := c1.price*c1.quantity
    return "hola " + c1.name + " \n esta es la description :" + c1.description + "\n cost=" + strconv.Itoa(compute)
    }

func (c2 car1) test() string{
    return "car price: " + strconv.Itoa(c2.price)
    }

func  display() string{
    return "is displaying"
    }

func main(){
 bm:= car1{"Ruben","todo la decricion esra disponible", 1000,3}
 t:= car1{price:100}
 fmt.Println (bm.start())
 fmt.Println(display())
 fmt.Println(t.test())
}