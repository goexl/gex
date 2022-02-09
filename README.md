# gex

Golang exec extension library，Golang外部命令执行扩展库，有如下功能

- 非常容易使用
- 一切皆可配置
- 检查器
- 生命周期
- 收集器
- 自动解决乱码

## 快速开始

`Gex`使用非常简单

```go
package main

import (
  `github.com/storezhang/gex`
)

func main() {
  _, _ = gex.Exec(`ping`, gex.Args(`www.163.com`))
}
```

> `Gex`有非常多的配置项，请参看文档

## 文档

[点击这里查看最新文档](https://gex.storezhang.tech)

## 捐助

![支持宝](doc/docs/.vuepress/public/donate/alipay-small.jpg)
![支持宝](doc/docs/.vuepress/public/donate/weipay-small.jpg)
