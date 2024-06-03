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

    id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
    if id == 0 {
        fmt.Println("In child, args=", os.Args)
    } else {
        fmt.Println("In parent, args=", os.Args)
    }
}
