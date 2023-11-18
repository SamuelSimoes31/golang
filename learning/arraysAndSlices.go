package main
import "fmt"

func main() {
    var a4 [4]int
    // var s []int
    s4 := make([]int, 4)

    fmt.Print("a4 ", a4)
    fmt.Printf(" %T\n",a4)
    // fmt.Print("s ",s)
    // fmt.Printf(" %T\n",s)
    fmt.Print("s4 ",s4)
    fmt.Printf(" %T\n",s4)

    a4_c := a4
    a4_c[0] = 5
    fmt.Println(a4_c, a4)

    s4_c := s4
    s4_c[0] = 5
    fmt.Println(s4_c, s4)
}