# 内存分配
全局对象：事先分配的内存空间段，程序启动时分配，结束时回收。
局部对象：分配在栈上，进入函数时分配，退出函数时回收。
动态对象：分配在堆上，需要时分配，不需要时回收。
# 规则
+ Rust 中的每一个值都有一个 所有者（owner）。
+ 值在任一时刻有且只有一个所有者。
+ 当所有者（变量）离开作用域，这个值将被丢弃。


```rust
{
    let s = String::from("hello"); // 从此处起，s 是有效的
    // 使用 s
}                                  // 此作用域已结束，
                                   // s 不再有效
```
# move(移动/转移)

Rust中的数值类型是实现了Copy trait的(存储在栈上的类型)，它们在赋值或者传递给函数时，会自动地进行复制，而不是转移所有权。例如：
```rust
let x = 42; // x拥有42的所有权
let y = x; // y拥有42的所有权，x也保留了42的所有权，因为42是Copy类型
println!("x is {}, y is {}", x, y); // 可以同时使用x和y
```
而对于没有实现Copy trait的类型，比如String(存储在堆上的类型)，它们在赋值或者传递给函数时，会转移所有权，原来的变量就不能再使用了34。例如：
```rust
let s1 = String::from("hello"); // String::from方法创建字符串"hello",s1拥有"hello"的所有权，内存分配在堆上
let s2 = s1; // s2拥有"hello"的所有权，s1失去了"hello"的所有权，因为String不是Copy类型
println!("s1 is {}", s1); // 编译错误，s1不能再使用
```
> 类似于go的字符串，string类型实质上是指向底层字符数组的指针(指向的底层数组/长度/容量)，s2拷贝了s1的长度和容量但不拷贝数据，s1和s2都指向了同一个字符数组"hello"的内存，s2拷贝了。当变量离开作用域后，Rust 会自动调用 drop 函数并清理变量的堆内存。而s1和s2离开作用域后它们都会尝试释放相同的内存，会导致二次释放（double free）的错误，因此为了确保内存安全，所有权转移后s1就不再有效，s1离开作用域后不会释放内存。

# clone(克隆)
clone会分配一块新的内存，拷贝s1指向的字符数组的数据，此时s1和s2指向的底层数组地址是不同的。又可以叫做深拷贝，上面不拷贝数据的则叫做浅拷贝。
```rust
let s1 = String::from("hello");
let s2 = s1.clone();

println!("s1 = {}, s2 = {}", s1, s2);
// s1 = hello, s2 = hello
```
实现copy trait的类型则不需要clone也能够复制数据，例如前面的整型类型


**copy trait**(标记类型):
+ 所有整数类型，比如 u32。
+ 布尔类型，bool，它的值是 true 和 false。
+ 所有浮点数类型，比如 f64。
+ 字符类型，char。
+ 元组，当且仅当其包含的类型也都实现 Copy 的时候。比如，(i32, + i32) 实现了 Copy，但 (i32, String) 就没有。

# 所有权与函数
```rust
fn main() {
    let s = String::from("hello");  // s 进入作用域

    takes_ownership(s);             // s 的值移动到函数里 ...
                                    // ... 所以到这里不再有效

    let x = 5;                      // x 进入作用域

    makes_copy(x);                  // x 应该移动函数里，
                                    // 但 i32 是 Copy 的，
                                    // 所以在后面可继续使用 x

} // 这里，x 先移出了作用域，然后是 s。但因为 s 的值已被移走，
  // 没有特殊之处

fn takes_ownership(some_string: String) { // some_string 进入作用域
    println!("{}", some_string);
} // 这里，some_string 移出作用域并调用 `drop` 方法。
  // 占用的内存被释放

fn makes_copy(some_integer: i32) { // some_integer 进入作用域
    println!("{}", some_integer);
} // 这里，some_integer 移出作用域。没有特殊之处
```

# 返回值、参数与作用域
转移返回值的所有权
```rust
fn main() {
    let s1 = gives_ownership();         // gives_ownership 将返回值
                                        // 转移给 s1

    let s2 = String::from("hello");     // s2 进入作用域

    let s3 = takes_and_gives_back(s2);  // s2 被移动到
                                        // takes_and_gives_back 中，
                                        // 它也将返回值移给 s3
} // 这里，s3 移出作用域并被丢弃。s2 也移出作用域，但已被移走，
  // 所以什么也不会发生。s1 离开作用域并被丢弃

fn gives_ownership() -> String {             // gives_ownership 会将
                                             // 返回值移动给
                                             // 调用它的函数

    let some_string = String::from("yours"); // some_string 进入作用域。

    some_string                              // 返回 some_string 
                                             // 并移出给调用的函数
                                             // 
}

// takes_and_gives_back 将传入字符串并返回该值
fn takes_and_gives_back(a_string: String) -> String { // a_string 进入作用域
                                                      // 

    a_string  // 返回 a_string 并移出给调用的函数
}
```
返回多个参数的所有权
```rust
fn main() {
    let s1 = String::from("hello");

    let (s2, len) = calculate_length(s1);

    println!("The length of '{}' is {}.", s2, len);
}

fn calculate_length(s: String) -> (String, usize) {
    let length = s.len(); // len() 返回字符串的长度

    (s, length)
}
```
# 引用与借用

与上一小节的代码不同的是calculate_length函数，以一个对象的引用作为参数而不是获取值的所有权
```rust
fn main() {
    let s1 = String::from("hello");

    let len = calculate_length(&s1);

    println!("The length of '{}' is {}.", s1, len);
}

fn calculate_length(s: &String) -> usize {
    s.len()
}// 这里，s 离开了作用域。但因为它并不拥有引用值的所有权，
  // 所以什么也不会发生
``` 
> s是指向s1的引用，但不拥有，当停止引用时所指的值也不会丢弃。

+ 引用不用获取所有权就可以使用值
+ 引用像一个指针，因为它是一个地址，我们可以由此访问储存于该地+ 址的属于其他变量的数据。 与指针不同，引用确保指向某个特定类型的有效值
+ **借用**就是创建一个引用的行为
+ 正如变量默认是不可变的，引用也一样。（默认）不允许修改引用的+ 值。

# 可变引用
类似于可变变量，使用mut修饰，其格式如下例子
```rust
fn main() {
    // let s = String::from("hello");  //默认引用不允许修改
    let mut s = String::from("hello");//可以修改

    change(&mut s);
}

fn change(some_string: &mut String) {
    some_string.push_str(", world");
}
```

但是可变引用存在限制，如果你有一个对该变量的可变引用，你就不能再创建对该变量的引用。这些尝试创建两个 s 的可变引用的代码会失败：
```rust
let mut s = String::from("hello");
let r1 = &mut s;
let r2 = &mut s;
println!("{}, {}", r1, r2);
```
> 这样的限制能够避免数据竞争(data race)

```rust
let mut s = String::from("hello");
let r1 = &s; // 没问题
let r2 = &s; // 没问题
let r3 = &mut s; // 大问题
println!("{}, {}, and {}", r1, r2, r3);
```
允许有多个不可变引用，因为它们都不能对引用进行修改，这是安全的，但是不可变引用不能再创建可变引用，以及不允许同时拥有多个相同的可变引用，这些是不安全的。

注意是否是同时拥有，以下这样是可行
```rust
let mut s = String::from("hello");
{
    let r1 = &mut s;
} // r1 在这里离开了作用域，所以我们完全可以创建一个新的可变引用
let r2 = &mut s;
```

# 悬垂引用
在具有指针的语言中，很容易通过释放内存时保留指向它的指针而错误地生成一个悬垂指针，其指向的内存可能已经被分配给其它持有者。

在 Rust 中编译器确保引用永远也不会变成悬垂状态：当你拥有一些数据的引用，编译器确保数据不会在其引用之前离开作用域。

# 引用规则总结

+ 在任意给定时间，要么只能有一个可变引用，要么只能有多个不可变引用。
+ 引用必须总是有效的。
  

# 切片
切片是特殊的引用，表示引用序列中的一个片段
```rust
let s = String::from("hello world");
let hello = &s[0..5];
let world = &s[6..11];
```
