# 参数

参数是在命令执行时，命令所需要的参数列表

比如说在执行`ping`命令时，需要指定主机域名，参数可以是任意数据类型，包括

- 数字
- 字符串
- 指针
- 结构体（需要实现`String`接口）

## 参数列表

在`Gex`里面很方便使用参数

<<< @/../example/guide/args_slice.go

## 命令行

为了方便调试用（很多情况下，大家都是在控制台先执行一次命令再写代码，使用列表会人为增加一次转换），`Gex`可以使用命令行来完成参数的传入

<<< @/../example/guide/args_cli.go

## 参数类型

实际上，`Gex`支持任意数据类型，包括

- 整形参数
    - `int`
    - `int8`
    - `int64`
    - `int32`
    - `uint`
    - `uint8`
    - `uint32`
    - `uint64`
    - 整形指针
- 浮点型参数
    - `float32`
    - `float64`
- 字符串参数
    - `string`
    - 字符串指针
- 布尔参数
    - `bool`
    - 布尔指针
- 结构体
    - 实现`String`接口
    - 结构体指针

## `fmt.Stringer`接口

任何实现了`fmt.Stringer`接口的类型，都可以被正确的解析成参数，`fmt.Stringer`接口的定义如下

```go
type Stringer interface {
	String() string
}
```
