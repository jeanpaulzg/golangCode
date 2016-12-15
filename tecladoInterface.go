package main
import "fmt"

type Teclado interface{
    teclasNumeros() int
    teclasAlpha() int
    }
    
type tecladoSp struct{
    tS int
    }

func (t tecladoSp) teclasNumeros() int{
    return 12
    }
    
func (t tecladoSp) teclasAlpha() int{
    return 30
    }
    
type tecladoEn struct{
    tE int
    }

func (t tecladoEn) teclasNumeros() int{
     return 8
    }
func (t tecladoEn) teclasAlpha() int{
    return 40
    }
    

func main(){
   // c:= new(tecladoSp)
    c:= tecladoSp{}
    c1:= new(tecladoEn)
    fmt.Println("numeros:",c.teclasNumeros())
    fmt.Println("letras:",c.teclasAlpha())
 
        fmt.Println("numeros:",c1.teclasNumeros())
    fmt.Println("letras:",c1.teclasAlpha())
    }