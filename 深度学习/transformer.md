https://zhuanlan.zhihu.com/p/338817680
https://blog.csdn.net/qq_38890412/article/details/120601834
# Seq2Seq(从序列到序列)
将一个作为输入的序列映射为一个作为输出的序列，这一过程由编码（Encoder）输入与解码（Decoder）输出两个环节组成, 前者负责把序列编码成一个固定长度的向量，这个向量作为输入传给后者，输出可变长度的向量。
RNN-CNN-Transformer都是如此
# RNN(循环神经网络)
简单来说，我们想翻译`I like deep learning`为中文，将这七个汉字的输入序列映射为输出序列`我喜欢深度学习`。在输出`我`时，会根据输入`I`来翻译，在输出`喜欢`时，会根据`I like`来翻译，输出`深度`时会·根据`I like deep`来翻译，输出`学习`是，会根据`I like deep learning`来翻译

RNN不能并行，指的是当前的状态会依赖前一个状态，前一个状态的计算结果会作为当前状态的输入。
# CNN(卷积神经网络)
简单来说是窗口式遍历，要翻译`I like deep learning`，假设每两个词一组，,`I like`,`like deep`,`deep learning`，分别提取每组的特征进行计算，所以 CNN 能够并行。但是仅如此这个神经网络并不能覆盖整个句子，所以需要再叠加 CNN，将多组的特征再提取一次进行计算，`I like`,`like deep`为一组以及`like deep`,`deep learning`为一组，以此类推地叠加，最后得到结果。

CNN能并行，指的是每一组的计算都能够同时进行，并不会依赖其他组的计算结果。
# Transformer
## self-attention(自注意力机制)
注意力机制: 权重分配机制，给需要关注的地方分配更高的权重，不相关的分配更低的权重。
https://blog.csdn.net/qq_38890412/article/details/120601834
### $XX^T$ 的含义
对于二维矩阵，用行向量$A_1,A_2$表示
$$
X=
 \begin{bmatrix}
   a_{11} & a_{12} \\
   a_{21} & a_{22} \\
  \end{bmatrix} =
  \begin{bmatrix}
   A_1 \\
   A_2 \\
  \end{bmatrix}
$$

$$
A_1=
\begin{bmatrix}
   a_{11} & a_{12} \\
\end{bmatrix},
A_2=
\begin{bmatrix}
   a_{21} & a_{22} \\
\end{bmatrix}
$$

$$
XX^T=
\begin{bmatrix}
   a_{11} & a_{12} \\
   a_{21} & a_{22} \\
  \end{bmatrix} 
  ·
  \begin{bmatrix}
   a_{11} & a_{21} \\
   a_{12} & a_{22} \\
  \end{bmatrix}=
  \begin{bmatrix}
   A_1A_1 & A_1A_2 \\
   A_2A_1 & A_2A_2 \\
  \end{bmatrix} 
$$
N维矩阵同理，对于$XX^T$得到的矩阵，其元素为原矩阵$X$的n个向量的两两内积，向量的内积的几何意义是
1. 表征或计算两个向量之间的夹角
2. 一个向量在另一个向量方向上的投影
$$A_1·A_2=|A_1|*|A_2|*cos\theta$$
投影值大小和夹角则表示向量之间的相关性，投影值大则意味两个向量相关度高，夹角为90度则两者线性无关

故矩阵$ XX^T$是一个方阵，我们以行向量的角度理解，里面保存了每个向量和自己与其他向量进行内积运算的结果。
### $Softmax$函数的含义
对 C 个元素中的第 i 个元素的softmax值为:
$$
Softmax(z_i)=\tfrac{e^{z_i}}{\sum^C_{c=1}e^{z_c}}
$$
如 {ln1,ln2,ln3} ,其softmax值分别为 1/6,2/6,3/6 
其意义在于将每个输出的结果都赋予一个概率值，比如在分类时可以用来表示属于每个类别的可能性
### $Softmax(XX^T)X$
$XX^T$得到可以表示两两向量之间的相关性的值，Softmax函数将这个值转换成概率值，其实就是为了得到权重，$Softmax(XX^T)$再与原来的$X$的积，其实就意味着加权求和，得到向量之间的相关度。

将$X$作为一个句子，那么其中的向量则为一个个词向量，通过这个算法得到这个句子本身中字词之间的相关度，所以叫`自注意`，例如对于`早上好`这句话中`早` `上` `好`三个字对于其中的`早`来说，`早上`的关联度就比`早早`、`早好`的关联度高。
### Q K V 矩阵
深度学习中不会直接使用$X^T$,为了增强模型的拟合能力，会使用可训练的参数矩阵:
$$W^Q,W^K,W^V$$
输入矩阵$X$分别与它们相乘得到Q(query)、K(key)、V(value)

$Attention(Q,K,V)=softmax(\tfrac{QK^T}{\sqrt{d_k}})V$
上述表示了Scaled Dot-Product Attention(缩放点积注意力)

${d_k}$为方差，总之就是使得Transformer在训练过程中的梯度值保持稳定

## multi-headed attention(多头注意力机制)
自注意力机制的缺陷就是：模型在对当前位置的信息进行编码时，会过度的将注意力集中于自身的位置， 使用多头注意力机制能增强self-attention关注多个位置的能力

例如`我养了一只猫，它很可爱`中,`它`可以指代前文的`猫`，但从另一个维度(形容词的角度)看`可爱`也是可以用来修饰`它`的，`它`与两者都有一定关联，使用多头注意力会使原本输出的一个向量变为多个向量

① 输入句子 
② 将句子进行标记,转换成词向量矩阵 $X$
③ 将 X 切分成 8 份，并与权重矩阵$W_i$ 相乘，构成输入向量$W_iX$
④ 计算 Attention 权重矩阵：$$Attention(Q,K,V)=softmax(\tfrac{QK^T}{\sqrt{d_k}})V$$
⑤ 最后将 8 个头的结果合并


## 位置编码
上面的计算只是知道词之间的关联关系，但并不知道可以将词放在句子的什么位置，即如何分析词序
以前的位置编码:
1. 1-N分配:
   太阳晒屁股啦 [1 2 3 4 5]
问题:越后面的数字越大，体现不了权重
2. [0,1]分配:
   太阳晒屁股啦 [0 0.2 0.4 0.6 0.8 1]
问题:句子长度不同位置编码不一样

Transformer模型的位置编码使用三角函数式的相对位置编码：有兴趣可以了解
https://blog.csdn.net/u013853733/article/details/107853989

https://zhuanlan.zhihu.com/p/106644634

https://kazemnejad.com/blog/transformer_architecture_positional_encoding/
