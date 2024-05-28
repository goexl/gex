# Gex

[![Sourcegraph](https://sourcegraph.com/github.com/goexl/gex/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/goexl/gex?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/goexl/gex)
[![Go Report Card](https://goreportcard.com/badge/github.com/goexl/gex?style=flat-square)](https://goreportcard.com/report/github.com/goexl/gex)
[![Build Status](https://github.ruijc.com:20443/api/badges/goexl/gex/status.svg)](https://github.ruijc.com:20443/goexl/gex)
[![Codecov](https://img.shields.io/codecov/c/github/golangex/gex.svg?style=flat-square)](https://codecov.io/gh/goexl/gex)
[![License](https://img.shields.io/github/license/goexl/gex)](https://raw.githubusercontent.com/goexl/gex/master/LICENSE)

Golang exec extension library，Golang外部命令执行扩展库，有如下功能

- 非常容易使用
- 一切皆可配置
- 可变参数
- 运行时目录
- 环境变量
- 管道
- 检查器
    - 字符串包含`Contains`
    - 字符串全包含`ContainsAll`
    - 字符串任何包含`ContainsAny`
    - 字符串相等`Equal`
    - 路径模式匹配`PathMatch`
    - 正则匹配`Regexp`
- 通知器
    - 应用内方法通知`Func`
    - 其它通知器，通过额外代码库实现，不增加`Gex`的复杂度
- 收集器
    - 字符串`String`
    - 文件`File`
    - 文件名`Filename`
    - 写入者`Writer`
- 自动解决乱码

## 快速开始

`Gex`使用非常简单

```go
package main

import (
  `github.com/goexl/gex`
)

func main() {
  _, _ = gex.Exec(`ping`, gex.Args(`www.163.com`, `-c`, 10))
}
```

or

```go
package main

import (
  `github.com/goexl/gex`
)

func main() {
  _, _ = gex.Exec(`ping`, gex.Cli(`www.163.com -c 10`))
}
```

> `Gex`有非常多的配置项，请参看[**使用文档**](https://gex.storezhang.tech)

## 文档

[点击这里查看最新文档](https://gex.storezhang.tech)

## 使用示例

[点击这里查看最新的代码示例](example)

## 交流

## 捐助

![支持宝](https://github.com/storezhang/donate/raw/master/alipay-small.jpg)
![微信](https://github.com/storezhang/donate/raw/master/weipay-small.jpg)
