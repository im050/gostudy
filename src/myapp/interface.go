package main

import "fmt"

type Animal interface {
    eat()
}

type Tiger struct {
    Name string
}

type Monkey struct {
    Name string
}

type Flower struct {
    Name string
}

func main() {
    var m = new(Monkey) //new 返回指针
    m.Name = "老猴子"


    var t = Tiger{Name: "老狮子"} //返回的是struct值类型

    //返回指针类型
    var sm = &Monkey{Name:"小猴子"}


    //必须将sm转换成interface类型
    if _, ok := interface{}(sm).(Animal); !ok {
        fmt.Println(m.Name, " not animal")
    } else {
        sm.eat()
    }

    if _, ok := interface{}(m).(Animal); !ok {
        fmt.Println(m.Name, " not animal")
    } else {
        m.eat()
    }

    //struct类型不属于Animal
    if _, ok := interface{}(t).(Animal); !ok {
        fmt.Println(t.Name, " not animal")
    } else {
        t.eat()
    }

    eat(m)
    eat(sm)
    //eat(t) //error cannot use t (type Tiger) as type Animal in argument to eat:
             //Tiger does not implement Animal (eat method has pointer receiver)
}

func(t *Tiger) eat() {
    fmt.Println(t.Name, "：我吃香蕉")
}

func(m *Monkey) eat() {
    fmt.Println(m.Name, "：我吃榴莲")
}

func eat(a Animal) {
    a.eat()
}