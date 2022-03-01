# 介绍

`收集器`主要用来收集输出，其中输出包括

- 标准输出
- 标准错误

而`收集器`的作用大概有

- 命令执行结束后分析
- 进度显示
- 其它待补充

`Gex`内置了常见收集器，包括

- 字符串`String`
- 文件`File`
- 文件名`Filename`
- 写入者`Writer`

## 使用

所有的收集器都是使用以`Collector`结尾的方法来做配置，包括

- `StringChecker`
- `WriterChecker`
