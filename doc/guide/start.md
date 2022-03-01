# 快速下手

::: warning 前提条件
<!--@formatter:off-->
Golang需要1.16及以上版本
<!--@formatter:on-->
:::

本文会帮助你从头搭建一个简单的Golang应用

## 新建项目

创建普通的Golang项目，容纳项目代码

``` shell
mkdir exec && cd exec
```

## 初始化`Go Module`

`Gex`建议使用`Go Module`安装，其它方式不在支持之列

``` shell
go mod init
```

## 创建启动文件

`Gex`推行瘦启动器，只有少量的代码（大部分是配置）

<<< @/../example/main.go

## 编译

编译Golang代码为可执行程序

``` shell
go build
```

## 运行

运行可执行程序，你就能看到程序的输出啦

``` shell
./exec
```
