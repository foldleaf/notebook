https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/07.6.md
## slice操作

### 插入
1. 在切片a后追加元素x
```go
a=append(a, x)
```
2. 在切片a后追加切片b的元素:
```go
a=append(a, b...)
```
3. 将切片a的元素复制到切片b:
```go
b:=make([]int,len(a))
copy(b,a)
```
4. 切片a扩展m个长度
```go
a=append(a,make([]int,m))
```
5. 在索引i处插入元素x
```go
a = append(a[:i], append([]T{x}, a[i:]...)...)
```
6. 在索引i处插入j长度的切片
```go
a=append(a[:i],append(make([]T,j),a[i:]...)...)
```
7. 在索引i处插入切片b的所有元素
```go
a=append(a[:i],append(b,a[i:]...)...)
```
### 删除
1. 删除切片索引i的元素
```go
a=append(a[:i],a[i+1:]...)
```
2. 删除切片索引i到j的元素:
```go
a=append(a[:i],a[j+1:]...)
```
### 查询
```go
// 索引i的元素
a[i]
// 切片的的长度
len(a)
// 切片a的容量
cap(a)
```
### 排序
```go
// 均为快速排序
// a为int切片、b为float64切片、c为string切片
// int类型的排序
sort.Ints(a)
// float64类型的排序
sort.Float64s(b)
// string类型的排序
sort.Strings(c)
```
```go
// 倒序
sort.Sort(sort.Reverse(sort.IntSlice(a)))
sort.Sort(sort.Reverse(sort.Float64Slice(b)))
sort.Sort(sort.Reverse(sort.StringSlice(c))))
```
### 遍历
```go
for i,v:=range slice{
   fmt.Printf("%d:%v\n",i,v)
}
```
## map操作
### 插入和更新
```go
map[key]=value
```
### 删除
```go
delete(mapName,key)
```
### 单个查询
```go
value,ok:=map[key]
if ok{
    fmt.Printf("存在/%v,值为%v",key,value)
}else{
    fmt.Printf("不存在/%v",key)
}
```
### 遍历
```go
for i,v:=range map{
    fmt.Printf("%d:%v\n",i,v)
}
```
## 关于string
string的结构本质上有两个字段，一个是指向底层字符数组的指针，一个是长度。我们可以将string认为是字符数组，可以使用切片的操作对string进行处理，在strings与strconv包中也有更方便的专门处理string的函数，详见字符串处理。
不过go中字符串是不可变的，如`str[i] = 'D'`是不允许的  
必须先将字符串转换成字节数组，然后再通过修改数组中的元素值来达到修改字符串的目的，最后将字节数组转换回字符串格式。
```go
s := "hello"
c := []byte(s)
c[0] = 'c'
s2 := string(c) // s2 == "cello"
```
