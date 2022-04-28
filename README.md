# go-endata 静态文件转GO文件
当你想把静态文件一起打包至二进制文件时，通常会使用`//go:embed static/logo.png`这样的注解，但是这样会将内容放置到全局变量，会大大增加运行时内存，`go-endata`通过将二进制流包含至函数中规避这一问题，能有效解决运行时内存使用过大问题。

## 使用方式

1. 不含gin框架路由文件
```bat
go-endata create -s ./tests/test_src -o ./tests/test_dst -a Bearki
```

2. 包含gin框架路由文件
```bat
go-endata create ^
-s ./tests/test_src ^
-o ./tests/test_dst ^
-a Bearki ^
--pack-prefix git.yoyo.link/yoyo/sdk/service/httpservice ^
-g
```