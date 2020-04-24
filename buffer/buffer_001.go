package main

import (
    "fmt"
    "bytes"
)

func main() {

    // 是一个变长的 buffer，具有 Read 和Write 方法
    // Buffer 就像一个集装箱容器，可以存东西，取东西（存取数据）

    buf1 := bytes.NewBufferString("hello")
    buf2 := bytes.NewBuffer([]byte("hello"))
    buf3 := bytes.NewBuffer([]byte{'h','e','l','l','o'})

    // 返回格式为 []byte

    fmt.Printf("%v,%v,%v\n",buf1,buf2,buf3)


    fmt.Printf("%v,%v,%v\n",buf1.Bytes(),buf2.Bytes(),buf3.Bytes())

    buf4 := bytes.NewBufferString("")
    buf5 := bytes.NewBuffer([]byte{})
    fmt.Println(buf4.Bytes(),buf5.Bytes())
}
