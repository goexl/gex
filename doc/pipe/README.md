# 介绍

`管道`主要用来连接多个命令，实现类型于如下效果

```shell
last | cut -d ' ' -f 1 | sort | uniq -c
```

::: warning 注意

上面的例子纯粹是演示使用，命令的作用是：使用last取出历史登录信息的账号，然后排序去重
:::

## 使用

`Gex`非常管道非常方便，利益于`语义化`设计，能够直接使用`Gex`本身的配置选项，比如

- 参数
- 运行时目录
- 环境变量
- 上下文
- 其它`Gex`支持的配置选项

<<< @/../example/pipe/usage.go
