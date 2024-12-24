package main

import (
    "fmt"
)

func main () {
    str := [] string{"HELLOZ", "abc"}
    var bstr [][] byte

    for _, line := range str {
        bstr = append(bstr, [] byte(line))
    }

    fmt.Println(str)
    fmt.Println(bstr)

    for i, line := range bstr {
        for j, chr := range line {
            bstr[i][j] = 'h'
            fmt.Println(chr)
        }
    }

    fmt.Println(str)
    fmt.Println(bstr)

}
