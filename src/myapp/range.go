package main

import "fmt"

func main() {


    //range有两个返回值，range不同类型的变量，返回值不同
    //array or slice , first return: index(int) second return: a[i]E, it means key => val
    //Strubg s字符串, fr: the same as array, sr: 符文整数 ???
    //map m map[K]V  fr: key k K, sr: value m[k]V , hashmap key => val

    //example 1:
    numbers := []int{0,1,2,3,4,5,6,7,8,9}
    fmt.Printf("%d, %d, %x\r\n", len(numbers), cap(numbers), numbers)
    for key := range numbers {
        fmt.Print( key, ",");
    }

    //example 2:
    //定义map , map[键类型][值类型]
    age := make(map[string]int)
    age["小红"] = 29
    age["小明"] = 30
    age["小强"] = 18
    for key, val := range age {
        fmt.Printf("%s 年龄： %d\r\n", key, val)
    }


}