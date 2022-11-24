package main
  
import "fmt"

type Querys struct {
    field string
    comparison  string
    value string
}

func main() {
  
    // taking an array
    // arr := [5]string{"Ronaldo", "Messi", "Kaka", "James", "Casillas"}
    

    a := Querys{"name", "==", "conall"}

    q := [3]Querys{a,a,a}
    fmt.Println(q)
    fmt.Println("The elements of the array are:")
    // using for loop
    for index, element := range q {
        // fmt.Println("At index", index, "value is", element)
      fmt.Println(index, " _ ", element)
    }
}
