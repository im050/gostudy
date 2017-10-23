package main

import "fmt"

type Animal interface {
    eat() string
}

type Tiger struct {
    Name string
}

type Monkey struct {
    Name string
}

func main() {
    var m = new(Monkey)
    var t = new(Tiger)
    m.Name = "老猴子"
    t.Name = "傻狮子"
    m.eat()
    t.eat()
    //目前不清楚New和这样的方式创建对象有没有区别,2017年10月23日21:17:38
    var sm = Monkey{Name:"小猴子"}
    sm.eat()
}

func(t Tiger) eat() {
    fmt.Println(t.Name, "：我吃香蕉")
}

func(m Monkey) eat() {
    fmt.Println(m.Name, "：我吃榴莲")
}