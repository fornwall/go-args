package main

import (
    "flag"
    "fmt"
    "os"
    "syscall"
)

func main() {
    fmt.Println("# os.Args: ", os.Args)

    flag.Parse()
    fmt.Println("# flag.Args(): ", flag.Args())

    path, err := os.Executable()
    if err == nil {
        fmt.Println("# os.Executable(): ", path)
    } else {
        fmt.Println("# os.Executable(): Error")
    }
}
