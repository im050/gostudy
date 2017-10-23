package main

import (
    "fmt"
    "myfunc"
    "encoding/json"
    "bytes"
)

type User struct {
    Name string
    Age  int
    Sex  string
}

func main() {
    var num1 int = 1
    var (
        var1 int = 1
        var2 int = 2
    )
    var hashMap = map[string]int{
        "One":1,
        "Two": 2,
    }
    var b bytes.Buffer
    text := "Hello world"
    user := User{
        "Memory",
        26,
        "Man",
    }
    var arr = [5]int{1,2,3,4,5}
    for i :=0; i<len(arr);i++ {
        fmt.Println(arr[i])
    }
    //字符串连接
    b.WriteString("My Name is ")
    b.WriteString(user.Name)
    s := b.String()
    fmt.Println(s)
    //hash map使用
    fmt.Print(hashMap["One"],hashMap["Two"]);
    //json使用
    userJson, err := json.Marshal(user)
    if (err == nil) {
        fmt.Printf("%s\r\n", userJson)
    }
    //数字与包
    fmt.Printf("Text: %s\r\n", text)
    fmt.Printf("var1: %d, var2: %d\r\n", var1, var2)
    fmt.Printf("myfunc.Test: %d\r\n", myfunc.Test(num1))

    s1, s2 := "你好吗", "我很好"
    fmt.Printf("%s, %s\r\n", s1, s2)
    s1, s2 = swapString(s1, s2)
    fmt.Printf("%s, %s\r\n", s1, s2)
}

func swapString(a, b string) (string, string) {
    //由于可以返回多个值，很方便交换
    //php list($a, $b) = array($b, $a)
    return b, a
}