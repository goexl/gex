# 路径匹配

此检查器是判断给定的模式是否是路径匹配的，举个例子

- 模式是*，那么能匹配的是
    - aaa
    - bbb
    - ccc
    - 任意字符
- 模式是a*e，那么能匹配的是
    - abe
    - ace
    - apple
- 举一反三，谢谢

## 上代码

<<< @/../example/checker/path_match.go
