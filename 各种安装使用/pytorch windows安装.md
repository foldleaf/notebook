# Anaconda安装
## 安装配置
镜像站：
https://repo.anaconda.com/archive/

环境变量添加路径:

D:\anaconda3  
D:\anaconda3\Scripts  
D:\anaconda3\Library\bin  

其中`D:\anaconda3`是选择的Anaconda安装路径
## 修改镜像源配置
清华源
```bash
conda config --add channels http://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free/win-64/
conda config --add channels http://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/main/win-64/
conda config --set show_channel_urls yes
conda config --add channels http://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/pytorch/win-64/
```
恢复原来的默认配置
```bash
conda config --remove-key channels
```
## 创建虚拟环境
创建环境:
命令python可以查看python版本，命令exit()退出python环境
```bash
conda create -n pytorch_cpu python=3.11
```
命令python可以查看python版本，命令exit()退出python环境
激活环境:
```bash
activate pytorch_cpu
```
退出激活环境:
```bash
deactivate
```
# pytorch安装
官网: https://pytorch.org/
根据环境选择版本，选择stable-windows-conda-python-none，此配置下的安装命令:
```bash
conda install pytorch torchvision torchaudio cpuonly -c pytorch
```
查看已安装的库
```bash
conda list
```
# 编辑器配置-vscode
+ 安装python插件，其中包含python相关多个插件
+ 选择编译环境:
  - 随便新建一个python文件
  - 点击右下角python语言模式的右边选择解释器(python.exe)
+ 注意:选择的是带'base'的那个！不是带'pytorch'的那个！它的路径应该为你的Anaconda安装根目录下的python.exe，例如D:\anaconda3\python.exe

至此，环境配置完毕





