# Golang 环境变量
## GOROOT
### Golang enviroment variable
```
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=/Users/linyulin/golang/workspace
```

# Project目录结构

## GOPATH的说明 

* (a) GOPATH允许多个目录，当有多个目录时，请注意分隔符，*NIX下是冒号，Windows下是分号。当有多个GOPATH是，默认将go get的内容放在第一个目录下。

* (b) GOPATH目录约定有三个子目录：

    * src 存放源代码（比如：.go / .c / .h / .s 等）
    * pkg 编译后生成的文件（比如：.a）
    * bin 编译后生成的可执行文件 （不建议把这个目录放在PATH下，虽然方便！）

# Golang命名规范与访问权限
* 1、golang的命名需要使用`驼峰命名法`，且**不能出现下划线**
* 2、golang中根据首字母的大小写来确定可以访问的权限。无论是方法名、常量、变量名还是结构体的名称，如果首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用
  可以简单的理解成，首字母大写是公有的，首字母小写是私有的
* 3、结构体中属性名的大写
如果属性名小写则在数据解析（如json解析,或将结构体作为请求或访问参数）时无法解析
```
type User struct {
        name string
        age  int
 }
func main() {
         user:=User{"Tom",18}
         if userJSON,err:=json.Marshal(user);err==nil{
       　　  fmt.Println(string(userJSON))   //数据无法解析
        }
}
```
如上面的例子，如果结构体中的字段名为小写，则无法数据解析。所以一般建议结构体中的字段大写