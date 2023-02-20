深度学习，决定从pytorch开始
pytorch中文手册
https://github.com/zergtant/pytorch-handbook
## 下载安装Anaconda
在archlinux下
```bash
# 完全版
yay -S annaconda
# 本人安装的轻量版
yay -S miniconda
```
按照提示
```bash
If your shell is Bash or a Bourne variant, enable conda for the current user with

    $ source /opt/miniconda/etc/profile.d/conda.sh
You could also add the above line to `~/.bashrc`

or, for all users, enable conda with

    $ sudo ln -s /opt/miniconda/etc/profile.d/conda.sh /etc/profile.d/conda.sh
```
在终端输入以下命令以启用conda
```bash
source /opt/miniconda/etc/profile.d/conda.sh
```
## 创建conda环境
```
# 查看 python 版本
python
# 使用 exit() 或 Ctrl+D 退出 python 环境

# 创建 conda 环境
# ( our_envName 是你定义的环境名，这里使用 pytorch )
# ( x.x 是指定的 python 版本，我这里是 3.10 )
conda create -n your_envName python=x.x
```
激活环境
```
# 激活环境，pytorch 是你之前定义的环境名
conda activate pytorch
# 取消激活状态
conda deactivate
```
## 通过anaconda安装pytorch
如果需要换源：换清华源
https://blog.csdn.net/weixin_42570192/article/details/124760665
```bash
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free/
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/main/
conda config --set show_channel_urls yes
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/cloud/pytorch/
```
如果需要安装cuda
```bash
yay -S cuda
```
这里不安装cuda，选择cpu,在官网选择对应的安装命令复制粘贴
```bash
conda install pytorch torchvision torchaudio cpuonly -c pytorch 
```
安装好pytorch后
```bash
import torch
import torchvision
```
## 更多操作命令
https://blog.csdn.net/qq_38870718/article/details/122796306

