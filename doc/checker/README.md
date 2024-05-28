# 介绍

`检查器`主要用来检查输出（标准输出流、标准错误流）当满足条件时就返回，不阻塞当前`goroutine`，举个例子

> Drone的[Git插件](https://github.com/dronestock/git)
> 功能是，使用Drone持续集成时，拉代码和推代码

之所以要写这么一个插件是因为：天朝网络Github经常抽风。而这里面其它一个功能就是Github加速
> Github加速使用的是[FastGithub](https://github.com/dotnetcore/FastGithub)

当插件启动时，首先判断是否要启动`FastGithub`，如果需要启动，则调用FastGithub命令，等待FastGithub可以正常服务时执行拉/推代码的操作。
而判断FastGithub是否正常启动就是使用的`检查器`来完成

`Gex`内置了常见检查器，包括

- 字符串包含`Contains`
- 字符串全包含`ContainsAll`
- 字符串任何包含`ContainsAny`
- 字符串相等`Equal`
- 路径模式匹配`PathMatch`
- 正则匹配`Regexp`

::: warning 注意

使用`检查器`默认开启了异步执行，除非检查器不满足，否则将在任意时间点返回
:::

## 使用

所有的检查器都是使用以`Checker`结尾的方法来做配置，包括

- `ContainsChecker`
- `ContainsAllChecker`
- `ContainsAnyChecker`
- `EqualChecker`
- `PathMatchChecker`
- `RegexpChecker`

## 配置

检查器同样也支持配置，目前支持的配置有

- 缓存`Cache`

