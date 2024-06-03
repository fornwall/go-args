package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    fmt.Println("# os.Args: ", os.Args)

    flag.Parse()
    fmt.Println("# flag.Args(): ", flag.Args())
}
