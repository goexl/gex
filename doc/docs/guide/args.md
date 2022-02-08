# 参数列表

参数是在命令执行时，命令所需要的参数列表

比如说在执行`ping`命令时，需要指定主机域名，参数可以是任意数据类型，包括

- 数字
- 字符串
- 指针
- 结构体（需要实现`String`接口）

## 定义一个参数

在`Gex`里面很方便使用参数

<<< @/../example/guide/args.go

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