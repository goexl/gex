# gex

[![Sourcegraph](https://sourcegraph.com/github.com/storezhang/gex/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/storezhang/gex?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/storezhang/gex)
[![Go Report Card](https://goreportcard.com/badge/github.com/storezhang/gex?style=flat-square)](https://goreportcard.com/report/github.com/storezhang/gex)
[![Build Status](https://github.ruijc.com:20443/api/badges/storezhang/gex/status.svg)](https://github.ruijc.com:20443/storezhang/gex)
[![Codecov](https://img.shields.io/codecov/c/github/storezhang/gex.svg?style=flat-square)](https://codecov.io/gh/storezhang/gex)
[![License](https://img.shields.io/github/license/storezhang/gex)](https://raw.githubusercontent.com/storezhang/gex/master/LICENSE)

Golang exec extension library，Golang外部命令执行扩展库，有如下功能

- 非常容易使用
- 一切皆可配置
- 检查器
  - 字符串包含`Contains`
  - 字符串全包含`ContainsAll`
  - 字符串任何包含`ContainsAny`
  - 字符串相等`Equal`
  - 路径模式匹配`PathMatch`
  - 正则匹配`Regexp`
- 通知器
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
  `github.com/storezhang/gex`
)

func main() {
  _, _ = gex.Exec(`ping`, gex.Args(`www.163.com`, `-c`, 10))
}
```

> `Gex`有非常多的配置项，请参看文档

## 文档

[点击这里查看最新文档](https://gex.storezhang.tech)

## 捐助

![支持宝](doc/docs/.vuepress/public/donate/alipay-small.jpg)
![支持宝](doc/docs/.vuepress/public/donate/weipay-small.jpg)
