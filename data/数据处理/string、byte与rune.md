# 关于引号
` "" `:双引号，包裹的类型是string  
` '' `:单引号，包裹的类型是rune,rune是int32的别名，只能是单个字符  
` `` `:反引号，包裹的类型是string  
1. ` `` `不会解析转义符，`""`会解析转义符
2. 在编辑器里` `` `中 可以回车换行
3. 使用`""`时不管有没有转义符都会遍历字符数组一遍寻找转义符，而使用` `` `则是直接输出，所以` `` `性能好一点
```go
    char1 := "A"
    char2 := 'A'
    char3 := `A`
    char4 := "asd\nfghjkl"
    //在编辑器里这串字符换行后有两个缩进
	char5 := `asd\nfg
		nhjkl`
	// char4:='AB' 不行，只能单个字符
	fmt.Printf("char1:%v(类型:%T)\n", char1, char1)
	fmt.Printf("char2:%v(类型:%T)\n", char2, char2)
	fmt.Printf("char3:%v(类型:%T)\n", char3, char3)
	fmt.Printf("char4:%v(类型:%T)\n", char4, char4)
	fmt.Printf("char5:%v(类型:%T)\n", char5, char5)
```
```go
char1:A(类型:string)
char2:65(类型:int32)
char3:A(类型:string)
char4:asd
fghjkl(类型:string)
char5:asd\nfg
		nhjkl(类型:string)
```
# byte 与 rune 
byte: unit8的别名。代表 ASCII 码的一个字符（占一个字节）
rune: int32的别名。用于标识 Unicode 字符，代表一个 UTF-8 字符(占多个字节，例如英文字母占一个字节，中文汉字一般占三个字节)
因为识别的字节数不一样，所以使用不同的处理方式结果会有所不同，请看下一小节
```go
    var char1 byte = 'A'
	char2 := 'A'	// 默认是 int32/rune
	var char3 byte = 65
	var char4 byte = '\x41'
	var char5 byte = '\101'
    // %c:字符的unicode码值; %T：变量的类型; %v:值的默认格式; %U:表示为Unicode格式
	fmt.Printf("char1:%c(类型:%T)(%v)(%U)\n", char1, char1,char1,char1)
	fmt.Printf("char2:%c(类型:%T)(%v)\n", char2, char2,char2)
	fmt.Printf("char3:%c(类型:%T)(%v)\n", char3, char3,char3)
	fmt.Printf("char4:%c(类型:%T)(%v)\n", char4, char4,char4)
	fmt.Printf("char5:%c(类型:%T)(%v)\n", char5, char5,char5)
	var char6 rune ='\u0041'
	var char7 int64='\U00000041'
	fmt.Printf("char6:%c(类型:%T)(%v)(%U)\n", char6, char6,char6,char6)
	fmt.Printf("char7:%c(类型:%T)(%v)(%U)\n", char7, char7,char7,char7)
```
```go
char1:A(类型:uint8)(65)(U+0041)
char2:A(类型:int32)(65)
char3:A(类型:uint8)(65)
char4:A(类型:uint8)(65)
char5:A(类型:uint8)(65)
char6:A(类型:int32)(65)(U+0041)
char7:A(类型:int64)(65)(U+0041)
```
# string
string的底层是字符数组

使用传统的`for-i++`循环时会按ASCII处理，按unit8的格式进行解析,得到的是[]byte;

使用`for-range`循环时则是按int32的格式进行解析，得到的是[]rune;
```go
    str := "goの世界"
	for i := 0; i < len(str); i++ {
		// byte/unit8类型，按 ASCII 处理,会有乱码
        // %c:字符的unicode码值; %T：变量的类型; %v:值的默认格式; %U:表示为Unicode格式
		fmt.Printf("str[%v]:%c(类型:%T)(%U)\n", i, str[i], str[i], str[i])
	}
	fmt.Println("--------")
	for i, v := range str {
		// rune/int32类型,按 UTF-8 处理
		fmt.Printf("str[%v]:%c(类型:%T)(%U)\n", i, v, v, v)
	}

	str1 := []byte(str)
	str2 := []rune(str)
	fmt.Printf("str1:%c,%v,%U\n", str1, str1, str1)
	fmt.Printf("str2:%c,%v,%U\n", str2, str2, str2)
	fmt.Printf("str的ASCII长度:%v\n", len(str))
	fmt.Printf("str的UTF-8长度:%v\n", utf8.RuneCountInString(str))
```
```go
str[0]:g(类型:uint8)(U+0067)
str[1]:o(类型:uint8)(U+006F)
str[2]:ã(类型:uint8)(U+00E3)
str[3]:(类型:uint8)(U+0081)
str[4]:®(类型:uint8)(U+00AE)
str[5]:ä(类型:uint8)(U+00E4)
str[6]:¸(类型:uint8)(U+00B8)
str[7]:(类型:uint8)(U+0096)
str[8]:ç(类型:uint8)(U+00E7)
str[9]:(类型:uint8)(U+0095)
str[10]:(类型:uint8)(U+008C)
--------
str[0]:g(类型:int32)(U+0067)
str[1]:o(类型:int32)(U+006F)
str[2]:の(类型:int32)(U+306E)
str[5]:世(类型:int32)(U+4E16)
str[8]:界(类型:int32)(U+754C)
str1:[g o ã  ® ä ¸  ç  ],[103 111 227 129 174 228 184 150 231 149 140],[U+0067 U+006F U+00E3 U+0081 U+00AE U+00E4 U+00B8 U+0096 U+00E7 U+0095 U+008C]
str2:[g o の 世 界],[103 111 12398 19990 30028],[U+0067 U+006F U+306E U+4E16 U+754C]
str的ASCII长度:11
str的UTF-8长度:5
```
使用str[i]这种方式得到的是按ASCII解析的byte字符，显示不了UTF-8编码的字符，自然会出现乱码。
`g`和`o`分别由str[0]和str[1]表示;
`の`由str[2]、str[3]和str[4]表示;
`世`由str[5]、str[6]和str[7]表示；
`界`由str[8]、str[9]和str[10]表示。
