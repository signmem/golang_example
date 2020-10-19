package main

import (
    "io/ioutil"
    "fmt"
)

func main() {
    // read the whole file at once
    b, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    fmt.Println(b)
    fmt.Println(string(b))

    // write the whole body at once
    err = ioutil.WriteFile("output.txt", b, 0644)
    if err != nil {
        panic(err)
    }
}
