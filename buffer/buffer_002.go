package main

import (
   "fmt"
   "bytes"
)

//   buffer 中 write func 如下
// func (b *Buffer) Write(p []byte) (n int,err error)


//  buffer 中 read func 如下
// func (b *Buffer) Read(p []byte)(n int,err error)


func btest(){

    // 使用 bytes.Buffer
    t1 := new(bytes.Buffer)
    testbytes := []byte{1,2,3,4,5,6,7,8,9,10}
    t1.Write(testbytes)
    fmt.Println(t1.Bytes())    // 输出方法 必须带 () 否则返回地址

    newbytes := []byte{33, 44, 55}
    t1.Write(newbytes)          // 这里也会变成追加信息到变量中
    fmt.Println(t1.Bytes()) 

    s3 := make([]byte,3)
    t1.Read(s3)
    fmt.Println(s3)
    fmt.Println(t1.Bytes())    
    // 发生了读之后， 就会从 buffer 中踢走了已经 read 的数据
    // 这个就是最重要的 buffer 用法， 读完就像消费完成， 方便链路继续使用

}

func ctest() {
    // 使用 bytes.NewBufferString

    t1 := bytes.NewBufferString("my hello")

    fmt.Println(t1.String())

    chars := "new character"
    t1.WriteString(chars)        // 这里会变成追加， 那么就系 t1 包含了两组字符

    fmt.Println(t1.String())


}


func dtest() {
     // ReadBytes 作为分隔符进行取值， 如果不给分隔符， 默认为空
    var d byte = 'f'
    buf := bytes.NewBufferString("xuxiaofeng")
    fmt.Println(buf.String())

    b,_ :=buf.ReadBytes(d)
    fmt.Println(string(b))              // 取走了 x...f 所有字符， 包括 f
    fmt.Println(buf.String())
}


func etest() {
    // Next 去下一个字符
    t1 := new(bytes.Buffer)
    testbytes := []byte{1,2,3,4,5,6,7,8,9,10}
    t1.Write(testbytes)
    fmt.Println(t1.Bytes())    // 输出方法 必须带 () 否则返回地址

    c := t1.Next(1)
    fmt.Println(c)
    c = t1.Next(3)
    fmt.Println(c)


   t2 := bytes.NewBufferString("my hello")
   u := t2.Next(3)           // 空格也会当前一个字符
   fmt.Println(t2)
   fmt.Println(string(u))    // 默认返回 []byte 需要使用 string() 转换一下为字符输出



}


func ftest(){

    //Reset() 会清空了当前 buffer 所有东西

    t1 := new(bytes.Buffer)
    testbytes := []byte{1,2,3,4,5,6,7,8,9,10}
    t1.Write(testbytes)
    fmt.Println(t1.Bytes())
    t1.Reset()
    fmt.Println(t1.Bytes())


}




func main() {
   //btest()
   //ctest()
   // dtest()
   // st()
   ftest()
}
