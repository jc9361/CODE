package main

import (
    "fmt"
) 

func main() {
    names := [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }

    fmt.Println(names)

    for i, s := range names {
        fmt.Println(i, s)
    }
}
