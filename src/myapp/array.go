package main

import (
    "fmt"
)

func main() {
    //普通数组
    var arr = [5]int{1,2,3,4,5}
    for key, val := range arr {
        fmt.Println(key , ":" , val)
    }

    //指针数组，和C类似
    var ptr [5]*int
    for key, _ := range arr {
        ptr[key] = &arr[key]
    }
    for i := 0; i<5; i++ {
        fmt.Printf("%d\r\n", *ptr[i])
    }

    var numbers []int
    //创建一个长度5, 容量5的切片
    numbers = make([]int, 5, 5)
    //切片的上限下限
    var numbers2 []int = numbers[2:]
    var numbers3 []int = numbers[:3]
    var numbers4 []int = numbers[1:2]
    //追加numbers2元素，同时会扩充capacity
    numbers2 = append(numbers, 2,3,4)
    //拷贝元素numbers到numbers2
    copy(numbers2, numbers)

    printSlices(numbers)
    printSlices(numbers2)
    printSlices(numbers3)
    printSlices(numbers4)

}

func printSlices(x []int) {
    fmt.Printf("len=%d, cap=%d, array=%x\r\n", len(x), cap(x), x)
}