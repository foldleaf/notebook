为了提高编程语言的熟练度，就必须要大量阅读有关资料。但是，如果你不知道它的意思，你怎么能大量阅读呢？

在本文中，我将不会关注于一两个概念，而是试着通过尽可能多的Rust 代码段，解释其中的关键字和符号的含义。

准备好了吗？让我们开始吧！

***
`let`引入了变量绑定(variable binding)
```rust
let x;  // 声明 "x"
x = 42; // 将 42 分配给 "x"
```
也可以写成一行
```rust
let x = 42;
```
您可以显式地指定变量的类型，用冒号表示类型注释:`:`
```rust
let x: i32; // i32是一个有符号的32位整数
x = 42;
// 有符号整数有: i8, i16, i32, i64, i128
// 无符号整数有: u8, u16, u32, u64, u128 
```
也可以写成一行
```rust
let x: i32 = 42;
```
如果您声明一个变量,之后再初始化，编译器将会阻止你初始化之前使用这个变量
```rust
let x;
foobar(x); // error: borrow of possibly-uninitialized variable: `x`
// 借用可能未初始化的变量: `x`
x = 42;
```
正确的写法是这样:
```rust
let x;
x = 42;
foobar(x); //  `x` 的类型将会在这被推断出来
```
下划线是一个特殊的变量名,更确切地说是 “空缺名”,意思就是扔掉一些东西:`_`
```rust
// 42是常数，所以无事发生
let _ = 42;

// 调用 `get_thing` 函数但是舍弃了扔掉了的结果
let _ = get_thing();
```
以下划线开头的变量名是普通名称，只是编译器不会警告它们不被使用:
```rust
// 变量声明但不使用会被编译器警告
// 在变量名前面加下划线可以去掉警告
let _x = 42;
```
可以对同一个变量名多次绑定-你可以掩盖(shadow)变量绑定
```rust
let x = 13;
let x = x + 3;
// 在这之后使用 `x` 指代的都是第二个 `x`
// 第一个  `x` 已不存在.（关于这点本人有补充，见下）
```
> 译者注: 掩盖(shadowing)是在作用域之内掩盖掉原有的变量名，在作用域结束之后掩盖失效，因为此处两个 `x` 作用域相同，所以第二个`x`会一直掩盖第一个 `x`

Rust 有元组(tuples)，您可以将其视为“不同类型的值的固定长度集合”。
```rust
let pair = ('a', 17);
pair.0; // 这是 'a'
pair.1; // 这是 17
```
如果想添加类型注释:
```rust
let pair: (char, i32) = ('a', 17);
```
元组在执行任务时可以被解构(destructured)，这意味着它们被分解为独立的字段:
```rust
let (some_char, some_int) = ('a', 17);
// 现在 some_char是 'a'，some_int是 17
```
这点在函数返回元组时特别有用:
```rust
let (left, right) = slice.split_at(middle);
```
当然，在解构元组时，可以用`_`来丢弃它的一部分
```rust
let (_, right) = slice.split_at(middle);
```
分号表示语句的结尾:
```rust
let x = 3;
let y = 5;
let z = y + x;
```
这意味着语句可以跨越多行
```rust
let x = vec![1, 2, 3, 4, 5, 6, 7, 8]
    .iter()
    .map(|x| x + 3)
    .fold(0, |x, y| x + y);
```
(我们待会儿再讨论它们的意思)
`fn`声明一个没有返回值的函数
```rust
fn greet() {
    println!("Hi there!");
}
```
下面的函数返回32位有符号整数，箭头`->`表示它的返回类型:
```rust
fn fair_dice_roll() -> i32 {
    4
}
```
一对花括号`{}`声明一个代码块(block)，它有自己的作用域
```rust
// This prints "in", then "out"
fn main() {
    let x = "out";
    {
        // this is a different `x`
        let x = "in";
        println!("{}", x);
    }
    println!("{}", x);
}
```

代码块也是表达式，可以计算为(evaluate)一个值
```rust
let x = 42;
// 两者等价
let x = { 42 };
```
在一个代码块中，可以有多个语句

```rust
let x = {
    let y = 1; // first statement
    let z = 2; // second statement
    y + z // 这是代码块的结尾 - 整个代码块返回的值
};
```
所以省略函数末尾语句的分号等价于return语句
```rust
fn fair_dice_roll() -> i32 {
    return 4;
}

fn fair_dice_roll() -> i32 {
    4
}
```
`if` 条件语句也是表达式

```rust
fn fair_dice_roll() -> i32 {
    if feeling_lucky {
        6
    } else {
        4
    }
}
```
`match`语句也是表达式

```rust
fn fair_dice_roll() -> i32 {
    match feeling_lucky {
        true => 6,
        false => 4,
    }
}
```
符号点`.`通常用于访问值的字段:

```rust
let a = (10, 20);
a.0; // this is 10

let amos = get_some_struct();
amos.nickname; // this is "fasterthanlime"
```
或者调用值的某个方法:

```rust
let nick = "fasterthanlime";
nick.len(); // this is 14
```
双冒号`::`类似，但是它用于命名空间

在这个例子中，`std`是一个crate,`cmd`是一个模块(module),`min`是一个函数
> 译者注:package、crate、module概念不要以其他语言的概念先入为主，这里简要提一下，不作展开
> 1. package为项目，就是cargo new创建的项目就称为package
> 2. crate是一个独立的编译单元，又分为binary(二进制) crate和library crate，硬要翻译这个名词的话叫做分隔箱比较恰当。
    binary crate的充要条件是有src/main.rs
    library crate的充要条件是有src/lib.rs
> 3. module是模块，用于内部代码的组织，可以控制作用域及私有性


```rust
let least = std::cmp::min(3, 8); // this is 3
```
use directives can be used to "bring in scope" names from other namespace:
`use`关键字用来引入其他命名空间的作用域里的内容，例如函数

```rust
use std::cmp::min;
let least = min(7, 1); // this is 1
```
在 use 关键字中，花括号还有另一个意思: 它们来源相同
如果我们想同时导入 min 和 max，我们可以执行以下任何操作:
```rust
// this works:
use std::cmp::min;
use std::cmp::max;

// this also works:
use std::cmp::{min, max};

// this also works!
use std::{cmp::min, cmp::max};
```
> 类似于javascript的`import {min,max} from 'xxx'`

A wildcard () lets you import every symbol from a namespace:*
通配符`*`允许你从命名空间导入所有内容

```rust
// this brings `min` and `max` in scope, and many other things
use std::cmp::*;
```
类型也命名空间，类型的方法可以作为常规函数调用:
```rust
let x = "amos".len(); // this is 4
let x = str::len("amos"); // this is also 4
```
str is a primitive type, but many non-primitive types are also in scope by default.
`str`是原始类型，但也有许多非原始类型

```rust
// `Vec` is a regular struct, not a primitive type
let v = Vec::new();

// this is exactly the same code, but with the *full* path to `Vec`
let v = std::vec::Vec::new();
```

这会正常编译，因为Rust会在每个模块的开头插入:

```rust
use std::prelude::v1::*;
```
(这引入了很多东西，`Vec`, `String`, `Option` ,`Result` 等等)

`struct`关键字用于声明结构体

```rust
struct Vec2 {
    x: f64, // 64-bit floating point, aka "double precision"
    y: f64,
}
```
使用结构体字面量初始化
```rust
Rust code
let v1 = Vec2 { x: 1.0, y: 3.0 };
let v2 = Vec2 { y: 2.0, x: 4.0 };
// the order does not matter, only the names do
```

用另一个结构体初始化剩余字段的快捷方式:
```rust
let v3 = Vec2 {
    x: 14.0,
    ..v2
    // 只能在最后使用，后面不能有逗号
};
```
注意，剩余字段前面不加任何东西时可以表示所有字段:
```rust
let v4 = Vec2 { ..v3 };
```
结构体可以像元组一样被解构

```rust
let (left, right) = slice.split_at(middle);
```
像这样

```rust
let v = Vec2 { x: 3.0, y: 6.0 };
let Vec2 { x, y } = v;
// `x` 为 3.0, `y` 为 `6.0`
```
还有这样

```rust
let Vec2 { x, .. } = v;
// 舍弃了 `v.y`
```

`let`模式匹配可以作为`if`的判断条件
```rust
struct Number {
    odd: bool,
    value: i32,
}

fn main() {
    let one = Number { odd: true, value: 1 };
    let two = Number { odd: false, value: 2 };
    print_number(one);
    print_number(two);
}

fn print_number(n: Number) {
    if let Number { odd: true, value } = n {
        println!("Odd number: {}", value);
    } else if let Number { odd: false, value } = n {
        println!("Even number: {}", value);
    }
}

// this prints:
// Odd number: 1
// Even number: 2
```
match选择支也是模式匹配，类似与 if let
```rust
fn print_number(n: Number) {
    match n {
        Number { odd: true, value } => println!("Odd number: {}", value),
        Number { odd: false, value } => println!("Even number: {}", value),
    }
}

// this prints the same as before
```
match必须是详尽的，即必须至少有一个选择支能够匹配上

```rust
fn print_number(n: Number) {
    match n {
        Number { value: 1, .. } => println!("One"),
        Number { value: 2, .. } => println!("Two"),
        Number { value, .. } => println!("{}", value),
        // if that last arm didn't exist, we would get a compile-time error
    }
}
```
如果很难穷尽所有情形, 可以使用下划线`_`匹配剩下的所有情形

```rust
fn print_number(n: Number) {
    match n.value {
        1 => println!("One"),
        2 => println!("Two"),
        _ => println!("{}", n.value),
    }
}
```

你可以声明你自己的类型的方法
```rust
struct Number {
    odd: bool,
    value: i32,
}

impl Number {
    fn is_strictly_positive(self) -> bool {
        self.value > 0
    }
}
```
然后普通地使用
```rust
fn main() {
    let minus_two = Number {
        odd: false,
        value: -2,
    };
    println!("positive? {}", minus_two.is_strictly_positive());
    // this prints "positive? false"
}
```
Variable bindings are immutable by default, which means their interior can't be mutated:
变量绑定在默认情况下是不可变的，这意味着它们的内部不能变化

```rust
fn main() {
    let n = Number {
        odd: true,
        value: 17,
    };
    n.odd = false; // error: cannot assign to `n.odd`,
                   // as `n` is not declared to be mutable
}
```
这也意味着该变量不能被再次分配

```rust
fn main() {
    let n = Number {
        odd: true,
        value: 17,
    };
    n = Number {
        odd: false,
        value: 22,
    }; // error: cannot assign twice to immutable variable `n`
}
```
`mut` 可以使变量绑定为可变的
```rust
fn main() {
    let mut n = Number {
        odd: true,
        value: 17,
    }
    n.value = 19; // all good
}
```
`Traits`(特型) 是多种类型共有的东西(类似于其他编程语言的接口)

```rust
trait Signed {
    fn is_strictly_negative(self) -> bool;
}
```


“孤儿规则”(orphan rules):
如果要实现某个trait，那么该trait和要实现该trait的类型至少有一个要在当前crate中定义

在上面已经定义了
`trait`:Signed
`struct`:Number

在自定义的的类型中实现自定义的trait
```rust
impl Signed for Number {
    fn is_strictly_negative(self) -> bool {
        self.value < 0
    }
}

fn main() {
    let n = Number { odd: false, value: -44 };
    println!("{}", n.is_strictly_negative()); // prints "true"
}
```
在其他类型中实现自定义trait(i32是一个原始类型):

```rust
impl Signed for i32 {
    fn is_strictly_negative(self) -> bool {
        self < 0
    }
}

fn main() {
    let n: i32 = -44;
    println!("{}", n.is_strictly_negative()); // prints "true"
}
```
在自定义类型中实现其他trait


```rust
//  `Neg` trait 用于一元运算符 `-`的重载

impl std::ops::Neg for Number {
    type Output = Number;

    fn neg(self) -> Number {
        Number {
            value: -self.value,
            odd: self.odd,
        }        
    }
}

fn main() {
    let n = Number { odd: true, value: 987 };
    let m = -n; // this is only possible because we implemented `Neg`
    println!("{}", m.value); // prints "-987"
}
```
An block is always for a type, so, inside that block, means that type:implSelf
一个代码块用于一个类型，所以在这个代码块的内部意味着这个类型的自我实现
```rust
impl std::ops::Neg for Number {
    type Output = Self;

    fn neg(self) -> Self {
        Self {
            value: -self.value,
            odd: self.odd,
        }        
    }
}
```
Some traits are markers - they don't say that a type implements some methods, they say that certain things can be done with a type.
有些trait是marker(标记)——它们不是说类型实现了某些方法，而是说某些事情可以用类型来完成。


举个例子，`i32` 实现了trait `Copy`(i32可以认为是Copy类型)，所以
```rust
fn main() {
    let a: i32 = 15;
    let b = a; // `a` 被拷贝
    let c = a; // `a` 再次被拷贝
}
```
这同样也可以
```rust
fn print_i32(x: i32) {
    println!("x = {}", x);
}

fn main() {
    let a: i32 = 15;
    print_i32(a); // `a` is copied
    print_i32(a); // `a` is copied again
}
```
但结构体`Number`不是`Copy`,所以这样不行:

```rust
fn main() {
    let n = Number { odd: true, value: 51 };
    let m = n; // `n` is moved into `m`
    let o = n; // error: use of moved value: `n`
}
```
这同样不行:
```rust
fn print_number(n: Number) {
    println!("{} number {}", if n.odd { "odd" } else { "even" }, n.value);
}

fn main() {
    let n = Number { odd: true, value: 51 };
    print_number(n); // `n` is moved
    print_number(n); // error: use of moved value: `n`
}
```
但`print_number`使用不可变引用(immutable reference)后则可以

```rust
fn print_number(n: &Number) {
    println!("{} number {}", if n.odd { "odd" } else { "even" }, n.value);
}

fn main() {
    let n = Number { odd: true, value: 51 };
    print_number(&n); // `n` is borrowed for the time of the call
    print_number(&n); // `n` is borrowed again
}
```
如果一个函数接受一个可变的引用，那么它也可以工作——但是只有当我们的变量绑定也是 mut 的时候

```rust
fn invert(n: &mut Number) {
    n.value = -n.value;
}

fn print_number(n: &Number) {
    println!("{} number {}", if n.odd { "odd" } else { "even" }, n.value);
}

fn main() {
    // this time, `n` is mutable
    let mut n = Number { odd: true, value: 51 };
    print_number(&n);
    invert(&mut n); // `n is borrowed mutably - everything is explicit
    print_number(&n);
}
```
Trait methods can also take by reference or mutable reference:self
Trait 方法可以通过引用或可变引用接收`self`:

```rust
impl std::clone::Clone for Number {
    fn clone(&self) -> Self {
        Self { ..*self }
    }
}
```
当调用 trait 方法时，接收方是隐式借用的:

```rust
fn main() {
    let n = Number { odd: true, value: 51 };
    let mut m = n.clone();
    m.value += 100;
    
    print_number(&n);
    print_number(&m);
}
```
注意:这两种写法是等价的

```rust
let m = n.clone();
let m = std::clone::Clone::clone(&n);
```
Marker traits like have no methods:Copy
像 Copy 这样的 Marker trait没有方法:

```rust
// note: `Copy` requires that `Clone` is implemented too
impl std::clone::Clone for Number {
    fn clone(&self) -> Self {
        Self { ..*self }
    }
}

impl std::marker::Copy for Number {}
```
现在Clone仍然可用

```rust
fn main() {
    let n = Number { odd: true, value: 51 };
    let m = n.clone();
    let o = n.clone();
}
```
但是 Number的值不再移动(move)

```rust
fn main() {
    let n = Number { odd: true, value: 51 };
    let m = n; // `m` is a copy of `n`
    let o = n; // same. `n` is neither moved nor borrowed.
}
```
有些trait非常常见，它们可以使用`derive`属性自动实现:
```rust
#[derive(Clone, Copy)]
struct Number {
    odd: bool,
    value: i32,
}
// 扩展了 `impl Clone for Number` 以及 `impl Copy for Number` 的代码块.
```
函数可以是泛型的:

```rust
fn foobar<T>(arg: T) {
    // do something with `arg`
}
```
它们可以有多个类型参数，然后可以在函数的声明和函数体中使用这些参数，而不是具体的类型:


```rust
fn foobar<L, R>(left: L, right: R) {
    // do something with `left` and `right`
}
```
类型参数通常也有约束，因此您可以对它们进行一些实际操作。
最简单的是通过 `trait` 来约束

```rust
fn print<T: Display>(value: T) {
    println!("value = {}", value);
}

fn print<T: Debug>(value: T) {
    println!("value = {:?}", value);
}
```
> Display 和 Debug都是一种trait，用于格式化输出

类型参数约束有更长的语法:

```rust
fn print<T>(value: T)
where
    T: Display,
{
    println!("value = {}", value);
}
```

约束可能更复杂: 可能需要一个类型参数来实现多个 trait:

```rust
use std::fmt::Debug;

fn compare<T>(left: T, right: T)
where
    T: Debug + PartialEq,
{
    println!("{:?} {} {:?}", left, if left == right { "==" } else { "!=" }, right);
}

fn main() {
    compare("tea", "coffee");
    // prints: "tea" != "coffee"
}
```
Generic functions can be thought of as namespaces, containing an infinity of functions with different concrete types.
泛型函数可以被看作是命名空间，包含无数具有不同具体类型的函数

Same as with crates, and modules, and types, generic functions can be "explored" (navigated?) using ::
与crate、module和type一样，泛型函数也可以使用双冒号`::`进行访问 

```rust
fn main() {
    use std::any::type_name;
    println!("{}", type_name::<i32>()); // prints "i32"
    println!("{}", type_name::<(f64, char)>()); // prints "(f64, char)"
}
```
This is lovingly called turbofish syntax, because looks like a fish.::<>
这被亲切地称为 turbofish(涡轮鱼?) 语法，因为`::<>` 看起来像一条鱼。

结构体也可以是泛型的:
```rust
struct Pair<T> {
    a: T,
    b: T,
}

fn print_type_name<T>(_val: &T) {
    println!("{}", std::any::type_name::<T>());
}

fn main() {
    let p1 = Pair { a: 3, b: 9 };
    let p2 = Pair { a: true, b: false };
    print_type_name(&p1); // prints "Pair<i32>"
    print_type_name(&p2); // prints "Pair<bool>"
}
```
标准库类型 Vec (分配在堆上的数组)是泛型的的
The standard library type (~ a heap-allocated array), is generic:Vec
```rust
fn main() {
    let mut v1 = Vec::new();
    v1.push(1);
    let mut v2 = Vec::new();
    v2.push(false);
    print_type_name(&v1); // prints "Vec<i32>"
    print_type_name(&v2); // prints "Vec<bool>"
}
```
Speaking of , it comes with a macro that gives more or less "vec literals":Vec
说到 Vec，它或多或少都会用到宏来提供“ vec 字面值”:
```rust
fn main() {
    let v1 = vec![1, 2, 3];
    let v2 = vec![true, false, true];
    print_type_name(&v1); // prints "Vec<i32>"
    print_type_name(&v2); // prints "Vec<bool>"
}
```
All of , or invoke a macro. Macros just expand to regular code.name!()name![]name!{}
`name!()` ，`name![]`或`name!{}`都调用了宏。宏只是扩展为常规代码

其实`println` 就是一个宏
```rust
fn main() {
    println!("{}", "Hello there!");
}
```

这段代码会扩展为以下等效的代码:
```rust
fn main() {
    use std::io::{self, Write};
    io::stdout().lock().write_all(b"Hello there!\n").unwrap();
}
```
`panic`(恐慌/异常) 也是一个宏。它通过一条错误消息和错误的文件名/行号(如果启用的话)强制停止执行:
```rust
fn main() {
    panic!("This panics");
}
// output: thread 'main' panicked at 'This panics', src/main.rs:3:5
```
Some methods also panic. For example, the type can contain something, or it can contain nothing. If is called on it, and it contains nothing, it panics:Option.unwrap()
有些方法也会引起`panic`。例如，Option 类型可以包含某些内容，也可以不包含。如果。在它上面调用 `.unpack()`但不包含任何东西的话，就会发生`panic`:

```rust
fn main() {
    let o1: Option<i32> = Some(128);
    o1.unwrap(); // this is fine

    let o2: Option<i32> = None;
    o2.unwrap(); // this panics!
}
// output: thread 'main' panicked at 'called `Option::unwrap()` on a `None` value', src/libcore/option.rs:378:21
```
`Option` 不是一个结构体 - 而是一个enum(枚举类)，它有两种变体(variant)

```rust
enum Option<T> {
    None,
    Some(T),
}

impl<T> Option<T> {
    fn unwrap(self) -> T {
        // enums variants 可以用于模式匹配:
        match self {
            Self::Some(t) => t,
            Self::None => panic!(".unwrap() called on a None option"),
        }
    }
}

use self::Option::{None, Some};

fn main() {
    let o1: Option<i32> = Some(128);
    o1.unwrap(); // this is fine

    let o2: Option<i32> = None;
    o2.unwrap(); // this panics!
}
// output: thread 'main' panicked at '.unwrap() called on a None option', src/main.rs:11:27
```
`Result` 也是一个 enum, 包含一个内容和一个错误
```rust
enum Result<T, E> {
    Ok(T),
    Err(E),
}
```
It also panics when unwrapped and containing an error.
当被解析出一个错误(error)时，它也会panic


变量绑定具有生命周期 (lifetime):

```rust
fn main() {
    // `x` doesn't exist yet
    {
        let x = 42; // `x` starts existing
        println!("x = {}", x);
        // `x` stops existing
    }
    // `x` no longer exists
}
```

引用同样也有生命周期
```rust
fn main() {
    // `x` doesn't exist yet
    {
        let x = 42; // `x` starts existing
        let x_ref = &x; // `x_ref` starts existing - it borrows `x`
        println!("x_ref = {}", x_ref);
        // `x_ref` stops existing
        // `x` stops existing
    }
    // `x` no longer exists
}
```
引用的生命周期不能超过其借用的变量绑定的生命周期期:
```rust
fn main() {
    let x_ref = {
        let x = 42;
        &x
    };
    println!("x_ref = {}", x_ref);
    // error: `x` does not live long enough
}
```
A variable binding can be immutably borrowed multiple times:
可以多次不可变地借用变量绑定:
```rust
fn main() {
    let x = 42;
    let x_ref1 = &x;
    let x_ref2 = &x;
    let x_ref3 = &x;
    println!("{} {} {}", x_ref1, x_ref2, x_ref3);
}
```

借用时，变量绑定不能发生变化
```rust
fn main() {
    let mut x = 42;
    let x_ref = &x;
    x = 13;
    println!("x_ref = {}", x_ref);
    // error: cannot assign to `x` because it is borrowed
}
```
不可变借用时，变量不能再被可变借用
```rust
fn main() {
    let mut x = 42;
    let x_ref1 = &x;
    let x_ref2 = &mut x;
    // error: cannot borrow `x` as mutable because it is also borrowed as immutable
    println!("x_ref1 = {}", x_ref1);
}
```

+ 在任意给定时间，要么只能有一个可变引用，要么只能有多个不可变引用。
+ 引用必须总是有效的。

函数参数中的引用也有生命周期:

```rust
fn print(x: &i32) {
    // `x` is borrowed (from the outside) for the
    // entire time this function is called.
}
```

具有引用参数的函数可以通过具有不同生命周期的借用来调用，因此:

+ 所有接受引用参数的函数都是泛型的
+ 生命周期是一个泛型的参数

生命周期的名称以单引号`'`开头

```rust
// elided (non-named) lifetimes:
fn print(x: &i32) {}

// named lifetimes:
fn print<'a>(x: &'a i32) {}
```
This allows returning references whose lifetime depend on the lifetime of the arguments:
谁的生命周期依赖于参数的生命周期，就返回谁的引用
```rust
struct Number {
    value: i32,
}

fn number_value<'a>(num: &'a Number) -> &'a i32 {
    &num.value
}

fn main() {
    let n = Number { value: 47 };
    let v = number_value(&n);
    // `v` borrows `n` (immutably), thus: `v` cannot outlive `n`.
    // While `v` exists, `n` cannot be mutably borrowed, mutated, moved, etc.
}
```
当输入只有一个生命周期时，不需要命名，并且所有东西都有相同的生命周期，因此下面的两个函数是等价的

```rust
fn number_value<'a>(num: &'a Number) -> &'a i32 {
    &num.value
}

fn number_value(num: &Number) -> &i32 {
    &num.value
}
```
Structs can also be generic over lifetimes, which allows them to hold references:
结构体可以在生命周期里都是泛型的，这使其可以接受引用
```rust
struct NumRef<'a> {
    x: &'a i32,
}

fn main() {
    let x: i32 = 99;
    let x_ref = NumRef { x: &x };
    // `x_ref` cannot outlive `x`, etc.
}
```
同样的代码，额外添加了函数
```rust
struct NumRef<'a> {
    x: &'a i32,
}

fn as_num_ref<'a>(x: &'a i32) -> NumRef<'a> {
    NumRef { x: &x }
}

fn main() {
    let x: i32 = 99;
    let x_ref = as_num_ref(&x);
    // `x_ref` cannot outlive `x`, etc.
}
```
The same code, but with "elided" lifetimes:
同样的代码，省略了声明周期
```rust
struct NumRef<'a> {
    x: &'a i32,
}

fn as_num_ref(x: &i32) -> NumRef<'_> {
    NumRef { x: &x }
}

fn main() {
    let x: i32 = 99;
    let x_ref = as_num_ref(&x);
    // `x_ref` cannot outlive `x`, etc.
}
```
`impl` 代码块在生命周期里也是泛型的:

```rust
impl<'a> NumRef<'a> {
    fn as_i32_ref(&'a self) -> &'a i32 {
        self.x
    }
}

fn main() {
    let x: i32 = 99;
    let x_num_ref = NumRef { x: &x };
    let x_i32_ref = x_num_ref.as_i32_ref();
    // neither ref can outlive `x`
}
```
但你也可以省略
```rust
impl<'a> NumRef<'a> {
    fn as_i32_ref(&self) -> &i32 {
        self.x
    }
}
```

如果不需要这个生命周期的名字，你可以狠狠地省略:

```rust
impl NumRef<'_> {
    fn as_i32_ref(&self) -> &i32 {
        self.x
    }
}
```
有一个特殊的生命周期:`static`，它对整个程序的生命周期都有效。(类比其他语言里的静态全局变量)

String 是`'static`(静态全局)的

```rust
struct Person {
    name: &'static str,
}

fn main() {
    let p = Person {
        name: "fasterthanlime",
    };
}
```
但是 string的所有者不是`static`的:
```rust
struct Person {
    name: &'static str,
}

fn main() {
    let name = format!("fasterthan{}", "lime");
    let p = Person { name: &name };
    // error: `name` does not live long enough
}
```

在最后一个示例中，变量`name`不是 `&'static str`，而是 `String`。它是动态分配的，会被释放。它的生命周期小于整个程序(即使它在 main 中)。

To store a non- string in , it needs to either:'staticPerson
要想在结构体`Person`中存储非`'static`的string，需要满足
A) 在生命周期里是泛型:

```rust
struct Person<'a> {
    name: &'a str,
}

fn main() {
    let name = format!("fasterthan{}", "lime");
    let p = Person { name: &name };
    // `p` cannot outlive `name`
}
```
或者
B) 获得string的所有权

```rust
struct Person {
    name: String,
}

fn main() {
    let name = format!("fasterthan{}", "lime");
    let p = Person { name: name };
    // `name` was moved into `p`, their lifetimes are no longer tied.
}
```

在结构体中，当字段与变量绑定同名时

```rust
    let p = Person { name: name };
```
可以简写成:
```rust
    let p = Person { name };
```
Rust 中的许多类型，有`所有者`(owned)的和`非所有者`(nom-owned)的变体:
> 引用可以在没有所有权的情况下使用被绑定的变量

+ Strings: String 是所有者, &str 是引用(reference)
+ Paths: PathBuf 是所有者, &Path 是引用
+ Collections: `Vec<T>` 所有者, `&[T]`是引用

Rust 有切片(slice)-切片是对多个连续元素的引用
可以借用 vector 的切片，例如
```rust
fn main() {
    let v = vec![1, 2, 3, 4, 5];
    let v2 = &v[2..4];
    println!("v2 = {:?}", v2);
}

// output:
// v2 = [3, 4]
```
这不足为奇。利用`Index`和`IndexMut`的`trait`就可以重载索引运算(`foo[index]`)。

(两个点)`..`语法只是表示范围(range),range只是在标准库里定义的少数结构体


范围是开放式的，一般是左闭右开区间，如果最右端使用等号`=`，右端就是闭区间

```rust
fn main() {
    // 0 or greater
    println!("{:?}", (0..).contains(&100)); // true
    // strictly less than 20
    println!("{:?}", (..20).contains(&20)); // false
    // 20 or less than 20
    println!("{:?}", (..=20).contains(&20)); // true
    // only 3, 4, 5
    println!("{:?}", (3..6).contains(&4)); // true
}
```
借用规则也适用于切片
```rust
fn tail(s: &[u8]) -> &[u8] {
  &s[1..] 
}

fn main() {
    let x = &[1, 2, 3, 4, 5];
    let y = tail(x);
    println!("y = {:?}", y);
}
```
等价于
```rust
fn tail<'a>(s: &'a [u8]) -> &'a [u8] {
  &s[1..] 
}
```
这是合法的

```rust
fn main() {
    let y = {
        let x = &[1, 2, 3, 4, 5];
        tail(x)
    };
    println!("y = {:?}", y);
}
```
但这因为`[1,2,3,4,5]`是一个静态数组

而这是不合法的:

```rust
fn main() {
    let y = {
        let v = vec![1, 2, 3, 4, 5];
        tail(&v)
        // error: `v` does not live long enough
    };
    println!("y = {:?}", y);
}
```

因为vector是分配在堆上的，没有`'static`的生命周期

&str实际上是切片:

```rust
fn file_ext(name: &str) -> Option<&str> {
    // this does not create a new string - it returns
    // a slice of the argument.
    name.split(".").last()
}

fn main() {
    let name = "Read me. Or don't.txt";
    if let Some(ext) = file_ext(name) {
        println!("file extension: {}", ext);
    } else {
        println!("no file extension");
    }
}
```
所以借用规则也适用于此:
```rust
fn main() {
    let ext = {
        let name = String::from("Read me. Or don't.txt");
        file_ext(&name).unwrap_or("")
        // error: `name` does not live long enough
    };
    println!("extension: {:?}", ext);
}
```
函数可以失败并特别地返回一个`Result`
```rust
fn main() {
    let s = std::str::from_utf8(&[240, 159, 141, 137]);
    println!("{:?}", s);
    // prints: Ok("🍉")

    let s = std::str::from_utf8(&[195, 40]);
    println!("{:?}", s);
    // prints: Err(Utf8Error { valid_up_to: 0, error_len: Some(1) })
}
```
如果你想在失败的情况下`panic`,可以使用`.unwrap()`

```rust
fn main() {
    let s = std::str::from_utf8(&[240, 159, 141, 137]).unwrap();
    println!("{:?}", s);
    // prints: "🍉"

    let s = std::str::from_utf8(&[195, 40]).unwrap();
    // prints: thread 'main' panicked at 'called `Result::unwrap()`
    // on an `Err` value: Utf8Error { valid_up_to: 0, error_len: Some(1) }',
    // src/libcore/result.rs:1165:5
}
```
或者用 .expect(), 可以自定义信息

```rust
fn main() {
    let s = std::str::from_utf8(&[195, 40]).expect("valid utf-8");
    // prints: thread 'main' panicked at 'valid utf-8: Utf8Error
    // { valid_up_to: 0, error_len: Some(1) }', src/libcore/result.rs:1165:5
}
```
或者使用模式匹配`match`

```rust
fn main() {
    match std::str::from_utf8(&[240, 159, 141, 137]) {
        Ok(s) => println!("{}", s),
        Err(e) => panic!(e),
    }
    // prints 🍉
}
```
或者使用:`if let`
```rust
fn main() {
    if let Ok(s) = std::str::from_utf8(&[240, 159, 141, 137]) {
        println!("{}", s);
    }
    // prints 🍉
}
```
或者将错误上报:
```rust
fn main() -> Result<(), std::str::Utf8Error> {
    match std::str::from_utf8(&[240, 159, 141, 137]) {
        Ok(s) => println!("{}", s),
        Err(e) => return Err(e),
    }
    Ok(())
}
```
或者使用 `?` 简单处理

```rust
fn main() -> Result<(), std::str::Utf8Error> {
    let s = std::str::from_utf8(&[240, 159, 141, 137])?;
    println!("{}", s);
    Ok(())
}
```
操作符 `*` 可以用于解引用，但是您不需要这样来访问字段或调用方法:

```rust
struct Point {
    x: f64,
    y: f64,
}

fn main() {
    let p = Point { x: 1.0, y: 3.0 };
    let p_ref = &p;
    println!("({}, {})", p_ref.x, p_ref.y);
}

// prints `(1, 3)`
```
如果类型为Copy时只需要这样:
```rust
struct Point {
    x: f64,
    y: f64,
}

fn negate(p: Point) -> Point {
    Point {
        x: -p.x,
        y: -p.y,
    }
}

fn main() {
    let p = Point { x: 1.0, y: 3.0 };
    let p_ref = &p;
    negate(*p_ref);
    // error: cannot move out of `*p_ref` which is behind a shared reference
}
```
```rust
// now `Point` is `Copy`
#[derive(Clone, Copy)]
struct Point {
    x: f64,
    y: f64,
}

fn negate(p: Point) -> Point {
    Point {
        x: -p.x,
        y: -p.y,
    }
}

fn main() {
    let p = Point { x: 1.0, y: 3.0 };
    let p_ref = &p;
    negate(*p_ref); // ...and now this works
}
```
`Closure`(闭包)只是 `Fn`、 `FnMut` 或 `FnOnce` 类型的函数，具有一定的`context`(语境/运行环境/上下文)

Their parameters are a comma-separated list of names within a pair of pipes (). They don't need curly braces, unless you want to have multiple statements.|
Closure的参数是一对管道(|)中以逗号分隔的名称列表。不需要花括号，除非您想要有多个语句。
```rust
fn for_each_planet<F>(f: F)
    where F: Fn(&'static str)
{
    f("Earth");
    f("Mars");
    f("Jupiter");
}
 
fn main() {
    for_each_planet(|planet| println!("Hello, {}", planet));
}

// prints:
// Hello, Earth
// Hello, Mars
// Hello, Jupiter
```
借用规则也适用与此
```rust
fn for_each_planet<F>(f: F)
    where F: Fn(&'static str)
{
    f("Earth");
    f("Mars");
    f("Jupiter");
}
 
fn main() {
    let greeting = String::from("Good to see you");
    for_each_planet(|planet| println!("{}, {}", greeting, planet));
    // our closure borrows `greeting`, so it cannot outlive it
}
```
这个例子不能运行:
```rust
fn for_each_planet<F>(f: F)
    where F: Fn(&'static str) + 'static // `F` must now have "'static" lifetime
{
    f("Earth");
    f("Mars");
    f("Jupiter");
}

fn main() {
    let greeting = String::from("Good to see you");
    for_each_planet(|planet| println!("{}, {}", greeting, planet));
    // error: closure may outlive the current function, but it borrows
    // `greeting`, which is owned by the current function
}
```
但这个可以:
```rust
fn main() {
    let greeting = String::from("You're doing great");
    for_each_planet(move |planet| println!("{}, {}", greeting, planet));
    // `greeting` is no longer borrowed, it is *moved* into
    // the closure.
}
```
An needs to be mutably borrowed to be called, so it can only be called once at a time.FnMut
`FnMut` 需要可变地借用才能被调用，因此一次只能调用一次

这是合法的:
```rust
fn foobar<F>(f: F)
    where F: Fn(i32) -> i32
{
    println!("{}", f(f(2))); 
}
 
fn main() {
    foobar(|x| x * 2);
}

// output: 8
```
而这不合法:
```rust
fn foobar<F>(mut f: F)
    where F: FnMut(i32) -> i32
{
    println!("{}", f(f(2))); 
    // error: cannot borrow `f` as mutable more than once at a time
}
 
fn main() {
    foobar(|x| x * 2);
}
```
这也是合法的:
```rust
fn foobar<F>(mut f: F)
    where F: FnMut(i32) -> i32
{
    let tmp = f(2);
    println!("{}", f(tmp)); 
}
 
fn main() {
    foobar(|x| x * 2);
}

// output: 8
```
`FnMut` 之所以存在，是因为一些闭包可变地借用本地变量
```rust
fn foobar<F>(mut f: F)
    where F: FnMut(i32) -> i32
{
    let tmp = f(2);
    println!("{}", f(tmp)); 
}
 
fn main() {
    let mut acc = 2;
    foobar(|x| {
        acc += 1;
        x * acc
    });
}

// output: 24
```
这些闭包不能被传递给期望的 Fn 函数:

```rust
fn foobar<F>(f: F)
    where F: Fn(i32) -> i32
{
    println!("{}", f(f(2))); 
}
 
fn main() {
    let mut acc = 2;
    foobar(|x| {
        acc += 1;
        // error: cannot assign to `acc`, as it is a
        // captured variable in a `Fn` closure.
        // the compiler suggests "changing foobar
        // to accept closures that implement `FnMut`"
        x * acc
    });
}
```
FnOnce 闭包只能调用一次。它们之所以存在，是因为某些闭包移出了在捕获时被移动的变量

```rust
fn foobar<F>(f: F)
    where F: FnOnce() -> String
{
    println!("{}", f()); 
}
 
fn main() {
    let s = String::from("alright");
    foobar(move || s);
    // `s` was moved into our closure, and our
    // closures moves it to the caller by returning
    // it. Remember that `String` is not `Copy`.
}
```
这是自然执行的，因为需要移动 FnOnce 闭包才能被调用。

所以这个例子是不合法的:
```rust
fn foobar<F>(f: F)
    where F: FnOnce() -> String
{
    println!("{}", f()); 
    println!("{}", f()); 
    // error: use of moved value: `f`
}
```
而且，如果你需要说服我们的关闭确实移动 s，这也是非法的

```rust
fn main() {
    let s = String::from("alright");
    foobar(move || s);
    foobar(move || s);
    // use of moved value: `s`
}
```
但这个是对的:
```rust
fn main() {
    let s = String::from("alright");
    foobar(|| s.clone());
    foobar(|| s.clone());
}
```
带有两个参数的闭包:
```rust
fn foobar<F>(x: i32, y: i32, is_greater: F)
    where F: Fn(i32, i32) -> bool
{
    let (greater, smaller) = if is_greater(x, y) {
        (x, y)
    } else {
        (y, x)
    };
    println!("{} is greater than {}", greater, smaller);
}
 
fn main() {
    foobar(32, 64, |x, y| x > y);
}
```
省略所有参数的闭包:
```rust
fn main() {
    foobar(32, 64, |_, _| panic!("Comparing is futile!"));
}
```
Here's a slightly worrying closure:
```rust
fn countdown<F>(count: usize, tick: F)
    where F: Fn(usize)
{
    for i in (1..=count).rev() {
        tick(i);
    }
}
 
fn main() {
    countdown(3, |i| println!("tick {}...", i));
}

// output:
// tick 3...
// tick 2...
// tick 1...
```
And here's a toilet closure:
```rust
fn main() {
    countdown(3, |_| ());
}
```
这样叫是因为`|_| ()`看起来像厕所


任何可迭代的内容都可以在 `for in` 循环中使用。

We've just seen a range being used, but it also works with a :Vec
我们之前看了range的使用，也可以用于Vec
```rust
fn main() {
    for i in vec![52, 49, 21] {
        println!("I like the number {}", i);
    }
}
```
或用于切片:
```rust
fn main() {
    for i in &[52, 49, 21] {
        println!("I like the number {}", i);
    }
}

// output:
// I like the number 52
// I like the number 49
// I like the number 21
```
或一个实际的迭代器(actual iterator)
```rust
fn main() {
    // note: `&str` also has a `.bytes()` iterator.
    // Rust's `char` type is a "Unicode scalar value"
    for c in "rust".chars() {
        println!("Give me a {}", c);
    }
}

// output:
// Give me a r
// Give me a u
// Give me a s
// Give me a t
```

即使迭代项被过滤、映射和扁平化
```rust
fn main() {
    for c in "SuRPRISE INbOUND"
        .chars()
        .filter(|c| c.is_lowercase())
        .flat_map(|c| c.to_uppercase())
    {
        print!("{}", c);
    }
    println!();
}

// output: UB
```
你可以从函数返回一个闭包
```rust
fn make_tester(answer: String) -> impl Fn(&str) -> bool {
    move |challenge| {
        challenge == answer
    }
}

fn main() {
    // you can use `.into()` to perform conversions
    // between various types, here `&'static str` and `String`
    let test = make_tester("hunter2".into());
    println!("{}", test("******"));
    println!("{}", test("hunter2"));
}
```

你甚至可以将一个对函数参数的引用移动到它返回的闭包中:
```rust
fn make_tester<'a>(answer: &'a str) -> impl Fn(&str) -> bool + 'a {
    move |challenge| {
        challenge == answer
    }
}

fn main() {
    let test = make_tester("hunter2");
    println!("{}", test("*******"));
    println!("{}", test("hunter2"));
}

// output:
// false
// true
```
或者略掉生命周期
```rust
fn make_tester(answer: &str) -> impl Fn(&str) -> bool + '_ {
    move |challenge| {
        challenge == answer
    }
}
