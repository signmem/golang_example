package main

import (
    "io"
    "os"
)

func main() {
    // open input file
    fcontroll, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }

    // close fi on exit and check for its returned error
    defer func() {
        if err := fcontroll.Close(); err != nil {
            panic(err)
        }
    }()


    // open output file
    fopen, err := os.Create("output.txt")
    if err != nil {
        panic(err)
    }

    // close fo on exit and check for its returned error
    defer func() {
        if err := fopen.Close(); err != nil {
            panic(err)
        }
    }()


    // make a buffer to keep chunks that are read
    buf := make([]byte, 1024)
    for {
        // read a chunk
        n, err := fcontroll.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if n == 0 {
            break
        }

        // write a chunk
        if _, err := fopen.Write(buf[:n]); err != nil {
            panic(err)
        }
    }
}
