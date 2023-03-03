## struct 与 json 的序列/反序列化
结构体
```go
type Person struct {
	Name   string  
	Age    int     
	Weight float64 
}
```
json
```json
{   
    "Name":"小明",
    "Age":18,
    "Weight":66.6
}
```
将结构体序列化为json
```go
// 将结构体序列化为json
func json.Marshal(v any) ([]byte, error)
// 将json反序列化为结构体
func json.Unmarshal(data []byte, v any) error
```
Unmarshal() 接收的是`[]byte`类型参数，其他类型需要使用`[]byte()`进行转换；同样的， Marshal() 输出的也是`[]byte`类型
## tag 标签
### 指定字段名
将结构体解析后输出的字段名改为标签中的名称，比如go中变量公开要首字母大写，需要将其序列化后的字段为小写。
fmt格式化输出中的 %s 占位符是自动转换为 string 再输出
```go
type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Weight float64 `json:"weight"`
}
func main{
   // 结构体转json
	p1 := Person{
		Name:   "小明",
		Age:    18,
		Weight: 0,
	}
	b,err:=json.Marshal(p1)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("str:%s\n",b)

	// json转结构体
	var p2 Person
	err=json.Unmarshal(b,&p2)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Printf("p2:%#v\n",p2)
}
```
```json
{
    "name":"小明",
    "age":18,
    "weight":66.6
}
```
### 忽略字段
添加tag以忽略某些字段
`"-"`:指定json序列化/反序列化时忽略此字段
`"omitempty"`:字段值为空时(不是零值，是没有该字段)，序列化/反序列化时忽略此字段
```go
type Person struct {
	Name   string  `json:"name"`            // 指定序列化/反序列化的字段名
	Age    int     `json:"age,omitempty"`   // 若字段值为空,序列化/反序列化时忽略此字段
	Weight float64 `json:"-"`               // 序列化/反序列化时忽略此字段
}
```

