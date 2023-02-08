https://juejin.cn/post/6968801451132846111

## 数组
数组的结构体在 /cmd/compile/internal/types2.Array
```go

package types2

// Array 表示数组类型.
type Array struct {
	len  int64 // 数组长度，即元素个数
	elem Type  // 元素类型
}

// NewArray 为给定的元素类型和长度返回一个新的数组类型。
// A negative length indicates an unknown length.
func NewArray(elem Type, len int64) *Array { return &Array{len: len, elem: elem} }

// Len 返回数组 a 的长度.
// A negative result indicates an unknown length.
func (a *Array) Len() int64 { return a.len }

// Elem 返回数组 a 的元素类型.
func (a *Array) Elem() Type { return a.elem }

func (a *Array) Underlying() Type { return a }
func (a *Array) String() string   { return TypeString(a, nil) }
```
## 切片
切片的结构体在 /runtime/slice.go
```go
type slice struct {
	array unsafe.Pointer  // 指向数组的指针
	len   int             // 长度
	cap   int             // 容量
}
```
创建slice
```go
func makeslice(et *_type, len, cap int) unsafe.Pointer {
	mem, overflow := math.MulUintptr(et.size, uintptr(cap))
	if overflow || mem > maxAlloc || len < 0 || len > cap {
		// NOTE: Produce a 'len out of range' error instead of a
		// 'cap out of range' error when someone does make([]T, bignumber).
		// 'cap out of range' is true too, but since the cap is only being
		// supplied implicitly, saying len is clearer.
		// See golang.org/issue/4085.
		mem, overflow := math.MulUintptr(et.size, uintptr(len))
		if overflow || mem > maxAlloc || len < 0 {
			panicmakeslicelen()
		}
		panicmakeslicecap()
	}

	return mallocgc(mem, et, true)
}
```
看着很复杂，其实前面全是判断错误，例如长度小于0，容量小于长度。最重要的是最后一句话，分配内存。
