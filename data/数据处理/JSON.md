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

### 嵌套结构体
#### 解析为单层json
```go
type User struct {
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Hobby   []string `json:"hobby"`
	Profile 
}

type Profile struct {
	Website string `json:"site"`
	Slogan  string `json:"slogan"`
}

u1 := User{
		Name:  "小明",
		Hobby: []string{"足球", "篮球"},
	}
```
```json
{
    "name":"小明",
    "email":"",
    "hobby":["足球","篮球"],
    "site":"",
    "slogan":""
}
```
#### 解析为多层嵌套json
只需要添加给嵌套的字段json标签
```go
type User struct {
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Hobby   []string `json:"hobby"`
	Profile `json:"profile"`
}
```
```json
{
    "name":"小明",
    "email":"",
    "hobby":["足球","篮球"],
    "profile":
        {
            "site":"",
            "slogan":""
        }
}
```
#### 忽略嵌套结构体的空值
嵌套字段只使用 omitempty 标签是不行的，还需要加指针才能忽略空值情况
```go
type User struct {
	Name    string   `json:"name"`
	Email   string   `json:"email,omitempty"`
	Hobby   []string `json:"hobby,omitempty"`
	*Profile `json:"profile,omitempty"`
}
```
#### 不修改原结构体忽略空值字段
```go
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PublicUser struct {
	*User             // 匿名嵌套
	Password *struct{} `json:"password,omitempty"`
}
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PublicUser struct {
	*User             // 匿名嵌套
	Password *struct{} `json:"password,omitempty"`
}

func omitPasswordDemo() {
	u1 := User{
		Name:     "七米",
		Password: "123456",
	}
	b, err := json.Marshal(PublicUser{User: &u1})
	if err != nil {
		fmt.Printf("json.Marshal u1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)  // str:{"name":"七米"}
}
```
### 字符串类型的数字解析
因为数据是string类型，所以不能对应相应数据类型的字段。使用string标签。
```go
type Card struct {
	ID    int64   `json:"id,string"`    // 添加string tag
	Score float64 `json:"score,string"` // 添加string tag
}
jsonStr1 := `{"id": "1234567","score": "88.50"}`
```
