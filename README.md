# go-dns

本地缓存


```go
net.LookupIP("www.baidu.com")
```
利用的是本机上53DNS服务器，通过他代理访问请求，拿到答案

两种方式
* 只实现和客户端报文解析
    - https://blog.cyeam.com/network/2015/02/03/dns
* 实现服务端与客户端










---
参考  
[字节序及 Go encoding/binary 库](https://huangwenwei.com/blogs/endian-and-encoding-binary-package)  
[dns之GitHub代码](https://github.com/changjixiong/goNotes/tree/master/dnsnotes)