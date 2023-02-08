string 是所有8位字节字符串的集合，通常但不一定表示 UTF-8编码的文本。字符串可以为empty，但不能为null。
字符串类型的值是不可变的。
```go
type string string
```
字符串实际上是一片连续的内存空间，我们也可以将它理解成一个由字符组成的数组。内存空间存储的字节共同组成了字符串，go 语言中的字符串只是一个只读的字节数组。