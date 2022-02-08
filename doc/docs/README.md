---

# @formatter:off
home: true
heroImage: /hero.png
heroText: Gex
tagline: Golang外部程序调用扩展库
actionText: 快速下手 →
actionLink: /guide/
features:
  - title: 简单的API封装
    details: 对外只暴露了Exec这一个方法就可以完成最简单的命令调用
  - title: 丰富的内置特性
    details: 内置支持输出编码转换以及自动解决乱码问题
  - title: 强大的扩展性
    details: 可以很方便的扩展整个框架而不侵入代码
  - title: 非常易于配置
    details: 框架的配置变得很容易，甚至可以无感知的配置系统
  - title: 线程安全
    details: 应用内单例，可以很方便的任何地方获得整个框架入口并作出配置而不用担心线程安全
  - title: 语义化
    details: 不需要记太多命令和参数以及配置项，语义化设计
  - title: 检查器
    details: 方便调用命令时在合适的时机继续往下执行而不阻塞
  - title: 收集器
    details: 方便对命令执行结果进行分析
  - title: 生命周期
    details: 方便在命令执行过期中接管生命周期回调
footer: MIT Licensed | Copyright © 2021-present Storezhang
# @formatter:on
---

# 简单得不能再简单

<<< @/../example/main.go

::: warning 警告
<!--@formatter:off-->
Golang最好是使用1.16及以上版本
<!--@formatter:on-->
:::
