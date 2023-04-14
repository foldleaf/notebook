# 环境配置
官网安装:
https://www.rust-lang.org/zh-CN/
注意:安装rust可能需要gcc等依赖，可能还需要visual studio安装一些必需模块，按照提示进行安装即可

vscode插件: rust-analyzer
# Hello World
创建项目
```bash
cargo new hello
```
编写程序主文件: hello/src/main.rs
```rust
fn main() {
    println!("Hello, world!");
}
```
执行代码:
```bash
# 进入目录
cd hello/src
#　编译文件
rustc main.rs
# 执行
main
```
推荐vscode插件`Code Runner`，一键编译并执行
