package main

import "fmt"

func main() {
    s := "你好Golang"
    s2 := Reverse(s)
    fmt.Print(s2)
    fmt.Print(Reverse2(s))
}

func Reverse(s string) string {
    str := []rune(s)
    strlen := len(str)
    str2 := make([]rune, strlen)
    for i := 0; i < strlen; i++ {
        str2[i] = str[strlen - i - 1];
    }
    return string(str2)
}

func Reverse2(s string) string {
    r := []rune(s)
    for i, j := 0, len(r) - 1; i < len(r) / 2; i, j = i + 1, j - 1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
