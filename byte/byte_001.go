package main 

import (
    _ "bytes"
    "fmt"
)

func byte_to_string() {

    /*
    当 string 放入 []byte 中，返回 [1 2 3 4 5] 属于 ascii 码
    要转换回来 string 则需要 string(byte) 
    */

    b := "justtest"
    a := []byte(b)
    // 输出 byte 
    fmt.Println(a)
    // 输出 string
    fmt.Println(string(a))
}


func main() {

	// []byte 输出字符方法

	byte_to_string()
}
