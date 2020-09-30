package main
import (
    "fmt"
    "strconv"
)
type Human struct {
    name string
    age int
    phone string
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human) String() string  {
    return "*** " + h.name+ " **  " + strconv.Itoa(h.age) + " years - " + h.phone + " ***"
}

func main() {
    Bob := Human{"Bob", 39, "000-7777-XXX"}

    fmt.Println(Bob.name + " *** " + strconv.Itoa(Bob.age))
    fmt.Println(Bob.age)
    fmt.Println(Bob)
    fmt.Println("This Human is : ", Bob)

    // 调用 String() 方法
    newString := Bob.String()
    fmt.Println(newString)
}
