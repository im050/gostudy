package main

import (
    "errors"
    "fmt"
)

func main() {
    x, err := divide(5, 0)
    if (err != nil) {
        fmt.Println(x," err:",err)
    }
}


func divide(x int, y int) (int, error) {
    var result int
    var err error
    if (y == 0) {
        result = 0
        err = errors.New("error!!!!!!!!!!")
    } else {
        result = x / y
    }
    return result, err
}