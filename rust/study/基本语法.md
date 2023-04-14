# 变量绑定
不可变变量:在变量赋值或对象创建结束之后就不能再改变它的值或状态的变量
可变变量:同理，在变量赋值或对象创建结束之后就可以再改变它的值或状态的变量

在rust中使用`let`关键字声明变量绑定
在rust中默认的都是不可变的，即不可以使用`变量=值`的方式修改绑定的值，如需修改则通过shadowing(掩盖/重影/隐藏)创建一个新的不可变变量绑定
变量如果需要可变需要使用`mut`修饰，可以直接修改绑定的值
```rust
let x = 1;      // 创建一个不可变变量，将 1 与变量绑定
x = 2;          // 错误，不能修改不可变变量绑定的值
let x = 2;      // 正确，通过shadowing创建一个新的不可变绑定
let mut y = 1;  // 创建一个可变变量，将 1 与变量绑定
y = 2;          // 正确，可以修改可变变量
let z: i16 = 3  // 显式声明，指定类型，默认的隐式声明会自动推断类型
```
rust的报错信息
```bash
 --> main.rs:3:5
  |
2 |     let x=1;
  |         -
  |         |
  |         first assignment to `x`
  |         help: consider making this binding mutable: `mut x`
3 |     x=2;
  |     ^^^ cannot assign twice to immutable variable
```
你会看到cannot assign twice to immutable variable（不能对不可变变量 x 二次赋值）
# shadowing(隐藏/掩盖/重影)
上面提到不可变变量不能二次赋值，但可以使用shadowing，将原变量“shadowing(隐藏)”
shadowing是使用同名的新变量隐藏掉了原变量，直到新变量也被隐藏或者其作用域结束为止，使用该变量名的行为都视为使用这个新变量。
可以多次隐藏。
```rust
fn main() {
    let x = 5;// x=5

    let x = x + 1;// x=6

    {
        let x = x * 2;
        println!("The value of x in the inner scope is: {x}");
        // x=12
        // x=12的作用域结束
    }

    println!("The value of x is: {x}");
    // x=6
}
```
# 常量
常量只能被设置为常量表达式，而不可以是其他任何只能在运行时计算出的值
可以出现在任何作用域（包括全局作用域）
```rust
const PI: f64 = 3.14159;
const MAGIC: i32 = 42;
const A: i32 =x*y;//错误， x*y是运行时才能确定
```
# 数据类型
数值类型：分为整型与浮点型
+ i8、i16、i32、i64、i128、isize
+ u8、u16、u32、u64、u128、usize
+ f32、f64
+ 其中 isize 和 usize 是指针大小的整数，因此它们的大小与机器架构相关。
+ 字面值 (literals) 写为 10i8、10u16、10.0f32、10usize 等。
+ 字面值如果不指定类型，则默认整数为 i32，浮点数为 f64。

布尔类型(bool)：true/false
字符类型：用单引号`''`包裹，unicode编码
复合类型：数组 (array)、切片 (slice)、字符串 (string)、元组 (tuple)
函数：函数也是一种类型
## 数组
`[T;N]`,T为元素类型，N为元素个数
+ N 是编译时常数 (compile-time constant)，也就是说数组的长度是固定的。
+ 运行时 (runtime) 访问数组元素会检查是否越界。
```rust
let arr1 = [1, 2, 3]; // 3个元素的数组
let arr2 = [2; 3];  // 3个元素的数组，元素都是 2
```
## 切片
切片类型的形式为 `&[T]`，例如 `&[i32]`
切片可以是可变的，也可以是不可变的。
切片不能直接创建，需要从别的变量借用 (borrow),这里暂时直接用go的切片来理解
```rust
let arr = [0, 1, 2, 3, 4, 5];
let total_slice = &arr; // 数组arr的切片
let total_slice = &arr[..]; // 与上面的效果一样
let partial_slice = &arr[2..5]; // 索引2到4的的元素：[2, 3, 4]
```
## 字符串
`String` 是在堆上分配空间、可以增长的字符序列。
`&str` 是 String 的切片类型
`str` 是没有大小的类型，编译时不知道大小，因此无法独立存在
```rust
let s: &str = "galaxy";
let s2: String = "galaxy".to_string();
let s3: String = String::from("galaxy");
let s4: &str = &s3;
```

## 元组
+ 元组是固定大小的、有序的、异构的列表类型。
+ 可以通过下标来访问元组的分量，例如 foo.0。
+ 可以使用 let 绑定来解构。
```rust
let foo: (i32, char, f64) = (72, 'H', 5.1);
let (x, y, z) = (72, 'H', 5.1);
let (a, b, c) = foo; // a = 72, b = 'H', c = 5.1
```

## 向量
`Vec`是分配在堆上的、可增长的数组
`<T>` 表示泛型，使用时代入实际的类型
使用 `Vec::new()` 或 `vec!` 宏来创建 Vec
```rust
// Explicit typing
let v0: Vec<i32> = Vec::new();
// v1 and v2 are equal
let mut v1 = Vec::new();
v1.push(1);
v1.push(2);
v1.push(3);
let v2 = vec![1, 2, 3];
// v3 and v4 are equal
let v3 = vec![0; 4];
let v4 = vec![0, 0, 0, 0];

let v2 = vec![1, 2, 3];
let x = v2[2]; // 3
```
+ 向量可以像数组一样使用 [] 来访问元素。
  - 在 Rust 中不能用 i32/i64 等类型的值作为下标访问元素。
  - 必须使用 usize 类型的值，因为 usize 保证和指针是一样长度的。
  - 其他类型要显式转换成 usize：
    let i: i8 = 2;
    let y = v2[i as usize];

## 类型转换
```rust
let x: i32 = 100;
let y: u32 = x as u32;
```
一般来说，只能在可以安全转换的类型之间进行转换操作
例如，[u8; 4] 不能转换为 char 类型
# 引用
`&` 取引用
`*` 解引用
与 c++ 一样，在 Rust 中，引用保证是合法的
```rust
let x = 12;
let ref_x = &x;
println!("{}", *ref_x); // 12
```
# 语句
## if-else 条件语句
```rust
let x = 1;
if x > 0 {
    println!("1")
} else {
    println!("0")
}
```

## while 循环
```rust
let mut x = 0;
while x < 100 {
    x += 1;
    println!("x: {}", x);
}
```

## loop 循环
无限循环，相当于while true和for(;;),直到遇到break语句或程序终止
loop 循环中的 break 语句可以返回一个值，作为整个循环的求值结果（另外两种循环没
有这个功能）
```rust
let mut x = 0;
    let y = loop {
        x += 1;
        if x * x >= 100 {
            break x;
        }
    };
    print!("{},{}",x,y)//10,10
```
## for 循环
```rust

for x in 0..10 {
    println!("{}", x);
    //输出0到9
}

let xs = [0, 1, 2, 3, 4];
for x in &xs {
    println!("{}", x);
    // 循环输出数组的元素值
}
```
## match 匹配语句
类似于switch语句
```rust
let x = 3;
// 匹配 x 的值
match x {   
    // 匹配到 1
    1 => println!("one fish"), 
    // 匹配到 2
    2 => {  
        println!("two fish");
        println!("two fish");
    } 
    // 其他情况使用下划线 `_`
    _ => println!("no fish for you"), 
}
//  no fish for you
```
可以匹配任意表达式
```rust
let x = 3;
let y = -3;
match (x, y) {
    (1, 1) => println!("one"),
    (2, j) => println!("two, {}", j),
    (_, 3) => println!("three"),
    (i, j) if i > 5 && j < 0 => println!("On guard!"),
    (_, _) => println!(":<"),
}
```
# 模式
# 函数

```rust
fn 函数名称(参数列表(参数名:参数类型,...))-> 返回值类型{
    //函数体
}

fn foo(x: T, y: U, z: V) -> T {
// ...
}
```
Rust 必须显式定义函数的参数和返回值的类型。
实际上编译器是可以推断函数的参数和返回值的类型的，但是 Rust 的设计者认为显式指定是
一种更好的实践。
```rust
fn add(x: i32, y: i32) -> i32 {
    x + y // 不带分号，返回x+y的值
}

fn sub(x: i32, y: i32) -> i32 {
    x - y; // 带分号，没有返回值，编译错误
}

fn mul(x: i32, y: i32) -> i32 {
    return x * y; // 使用return关键字，可以加分号
}
```
因为Rust的函数也是表达式，它们的返回值是由函数体最后一个表达式的值决定的。如果最后一个表达式后面加了分号，那么它就变成了一个语句，而语句没有值，所以函数就没有返回值了。

# 宏

# print! 和 println!
```rust
let x = "foo";
print!("{}, {}, {}", x, 3, true);
// => foo, 3, true
println!("{:?}, {:?}", x, [1, 2, 3]);
// => "foo", [1, 2, 3]
let y = 1;
println!("{0}, {y}, {0}", x);
// => foo, 1, foo
```
使用 `{}` 来做字符串插入，`{:?}` 做调试输出。
有些类型，例如数组和向量，只能用调试输出的方式来打印
`{}` 里可以加数字，表示第几个参数，新版本还可以把变量名写在 `{}` 里
