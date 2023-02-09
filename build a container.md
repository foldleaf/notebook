https://www.infoq.com/articles/build-a-container-golang/

我想玩个游戏。现在，在你的脑海里思考，然后告诉我“容器(container)”是什么。好了吗？好吧。让我猜猜你会说什么:

你可能会想到下面一点或更多：

* 一种共享资源的方式
* 进程隔离
* 有点像轻量级的虚拟化
* 将根文件系统和元数据打包在一起
* 有点像 chroot jail
* 集装箱什么的
* docker所做的一切

一个词的意思太多了！“容器”这个词已经开始用于许多(有时是重叠的)概念。它被用于类似容器化，以及用于实现它的技术。如果我们把这些分开考虑，我们会得到一个更清晰的画面。那么，让我们来讨论一下为什么要使用容器，以及如何使用容器。(然后我们再说回为什么使用容器)。

## 起初
最初，有这么一个程序。让我们称之为run.sh,现在需要把它拷贝到一个远程服务器上,并且运行它。然而，在远程计算机上运行任意代码是不安全的，并且难以管理和扩展。所以我们发明了虚拟专用服务器和用户权限。一切都很顺利。

但是run.sh有一些依赖项，它需要主机有某些库，而且它在远程和本地工作完全不同。于是我们发明了AMIs (Amazon Machine Images) 、 VMDKs (VMware images) 、 Vagrantfiles 等等, 一切又很顺利。

嗯，这些都很不错。这些包很大，也并非标准化，所以很难有效地运输。因此，我们发明了缓存(caching)。一切又又很顺利
正是caching使得 Docker 映像比 vmdks 或vagrantfiles更加有效。它允许我们在一些常见的基础映像上运输增量，而不是移动整个映像。这意味着我们有能力把整个环境从一个地方运送到另一个地方。这就是为什么当你“ docker run whatever”的时候，即使它描述了整个操作系统映像，它也会立即开始运行。我们将在(N 部分)中更详细地讨论这是如何工作的。

这就是容器的意义所在。它们是基于捆绑依赖关系的，这样我们就可以用可重复、安全的方式发送代码。但这是高层次的目标，而不是定义。我们来谈谈实际吧。

## 创建一个容器
所以(这次是真的!)什么是容器？如果创建一个容器就像 create _ Container 系统调用那样简单就好了。可惜不是，不过很接近了。要在底层讨论容器，我们必须讨论三样东西。分别是命名空间(namespace)、 cgroup 和分层文件系统(layered filesystems)。虽然还有其他的东西，不过这三者是这个魔术的主要组成部分。

### 命名空间 Namespace
命名空间提供了在一台计算机上运行多个容器所需的隔离（环境），同时为每个容器提供看起来像是它自己的环境的东西。就目前来说命名空间有 6 个，每个都可以被独立地请求，相当于给一个进程(及其子进程)一个机器资源子集的视图

命名空间有：
* PID：Pid 命名空间为进程及其子进程提供了系统中这些进程子集的视图。可以把它想象成一个映射表。当 pid 命名空间中的进程向内核请求进程列表时，内核查看映射表。如果进程存在于表中，则使用映射的 ID 而不是真正的 ID。如果它不存在于映射表中，内核就假装它根本不存在。Pid 名命名空间创建 pid 1中创建的第一个进程(通过将其主机 ID 映射为1)，展现出容器中的隔离进程树。
* MNT：在某种程度上，这是最重要的。Mount 命名空间为进程提供了它们自己的 mount 表。这意味着它们可以在不影响其他名称空间的情况下挂载和卸载目录（包括主机命名空间），更重要的是，结合 pivot _ root 系统调用 ，它允许进程拥有自己的文件系统。这就是我们如何让一个进程认为它在 ubuntu、 busybox 或者 alpine 上运行的方法——通过交换容器的文件系统能够发现
* NET：网络命名空间为使用它的进程提供它们自己的网络栈。一般来说，只有主网络名称空间(当您开始使用计算机时启动的进程)实际上会附加任何实际的物理网卡。但是我们可以创建虚拟以太网对一一连接的以太网卡，其中一端可以放在一个网络名称空间中，另一端可以放在另一个网络名称空间中，从而在网络名称空间之间创建一个虚拟链接。这有点像在一台主机上有多个 IP 协议栈相互通信。通过一点路由魔术，允许每个容器与现实世界对话，同时将每个容器与其自己的网络栈隔离开来。
* UTS：UTS 命名空间为其进程提供了它们自己的系统主机名和域名视图。输入 UTS 命名空间后，设置主机名或域名不会影响其他进程。
* IPC：IPC 命名空间隔离各种行程间通讯机制，例如消息队列。
* USER：用户名称空间是最近添加的，从安全角度来看，它可能是最强大的。用户名称空间将进程看到的 uid 映射到主机上的一组不同的 uid (以及 gids)。这很有用。使用用户名称空间，我们可以将容器的根用户 ID (即0)映射到主机上的任意(以及无特权的) uid。这意味着我们可以让容器认为它具有 root 访问权限——我们甚至可以在特定于容器的资源上给它类似于 root 的权限——而不必在根名称空间中给它任何特权。容器可以自由地以 uid0的形式运行进程——它通常等同于拥有 root 权限——但是内核实际上是将这个 uid 映射到一个非特权的实际 uid。大多数容器系统不会将容器中的任何 uid 映射到调用名称空间中的 uid0：换句话说，容器中根本没有具有真正 root 权限的 uid。

大多数容器技术将用户的过程放入上述所有名称空间，并初始化名称空间以提供标准环境。例如，这相当于在容器的隔离网络名称空间中创建一个初始 Internet 卡，并与主机上的实际网络连接。

### CGroups
Cgroup 可以诚实地成为它们自己的整篇文章(我保留编写一篇文章的权利!)。我将在这里相当简短地介绍它们，因为一旦您理解了这些概念，您就可以直接在文档中找到很多内容。基本上，cgroup 将一组进程或任务 ID 集合在一起，并对它们应用限制。在名称空间隔离进程的地方，cgroup 强制进程之间进行公平(或不公平——这取决于您，疯狂吧)的资源共享。

内核将 Cgroup 公开为一个可以挂载的特殊文件系统。通过简单地将进程 ID 添加到任务文件，可以将进程或线程添加到 cgroup，然后通过编辑该目录中的文件来读取和配置各种值。

### 分层文件系统
名称空间和 CGroup是关于容器化的隔离和资源共享。他们就像码头的大金属板和保安。分层文件系统是我们能够有效地移动整个机器映像的方式: 所以船会浮起来而不是下沉。

在基本层次上，分层文件系统相当于优化调用，为每个容器创建根文件系统的副本。有很多方法可以做到这一点。Btrfs 使用拷贝在文件系统层进行写，Aufs 则使用“ union mount”。因为有很多方法可以达到这一步，这篇文章将使用一些非常简单的东西: 我们将真的做一个拷贝。虽然很慢，但能够工作。

### 创建容器
#### 第一步：建立骨架
我们先把粗糙的骨架搭好。假设您已经安装了 golang 编程语言 SDK 的最新版本，然后打开一个编辑器，并复制到以下清单中。

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		panic("wat should I do")
	}
}

func parent() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

func child() { 
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
```

那这是做什么的呢? 我们从 main 开始，我们读第一个部分。如果它是‘ run’，那么我们运行 father ()方法，如果它是 child () ，那么我们运行 child 方法。parent方法运行“/proc/self/exe”，这是一个特殊文件，包含当前可执行文件的内存映像。换句话说，我们重新运行自己，但是传递 child 作为第一个参数

这是怎么回事？其实没什么，它只是让我们执行另一个程序(`os.Args[2:]`) ，执行用户请求的程序。但是，通过这个简单的脚手架，我们可以创建一个容器。

#### 第二步：添加命名空间
要向我们的程序添加一些命名空间，我们只需要添加一行
```go
cmd.SysProcAttr = &syscall.SysProcAttr{
	Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
}
```
如果您现在运行您的程序，您的程序将在 UTS、 PID 和 MNT 名称空间内运行

#### 第三步：根文件系统
目前你的进程位于一组独立的命名空间中（此时可以随意尝试将其他名称空间添加到您的上面的Cloneflags 中）。但是文件系统看起来和主机一样，这是因为您处于挂载命名空间中，但是初始挂载是从创建的名称空间继承的。让我们改变一下。我们需要以下四行简单的代码来切换到根文件系统。将它们放在‘ child ()’函数的开始位置。
```go
must(syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""))
	must(os.MkdirAll("rootfs/oldrootfs", 0700))
	must(syscall.PivotRoot("rootfs", "rootfs/oldrootfs"))
	must(os.Chdir("/"))
```
最后两行很重要，它们告诉操作系统将“/”的工作目录移动到“ rootfs/oldrootfs”，并将新的 rootfs 目录交换到“/”。在 pivotroot 调用完成之后，容器中的/目录将引用 rootfs。（需要绑定挂载调用来满足“ pivotroot”命令的某些需求——操作系统要求使用“ pivetroot”来交换两个文件系统，这两个文件系统不是同一棵树的一部分，而是将rootf绑定到它自己实现的。没错，这很蠢）

#### 第四步： 初始化容器的世界
此时，您已经在一组独立的名称空间中运行了一个进程，并选择了一个根文件系统。我们已经跳过了设置 cgroups，尽管这非常简单，并且我们已经跳过了根文件系统管理，它允许您有效地下载和缓存我们“ pipitroot”到的根文件系统映像。

我们还跳过了容器设置。这里有一个独立名称空间中的新容器。我们已经通过以 rootfs 为轴心设置了 mount 名称空间，但是其他名称空间有它们的默认内容。在实际的容器中，在运行用户进程之前，我们需要为容器配置“世界”。因此，例如，我们要设置网络，在运行进程之前切换到正确的 uid，设置我们想要的任何其他限制(比如删除功能和设置 rlimit)等等。这可能会推动我们超过100行。
#### 第五步：归一
因此，在这里，它是一个超级超级简单的容器，在(方式)不到100行的去。显然这是有意为之的简单。如果你在生产中使用它，你是疯了，更重要的是，你自己。但我认为，看到一些简单而粗糙的东西，可以让我们对正在发生的事情有一个真正有用的了解。让我们来看看清单 A。
```go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		panic("wat should I do")
	}
}

func parent() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

func child() {
	must(syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""))
	must(os.MkdirAll("rootfs/oldrootfs", 0700))
	must(syscall.PivotRoot("rootfs", "rootfs/oldrootfs"))
	must(os.Chdir("/"))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
```


## 所以容器究竟是什么？
接下来可能会有点争议。对于我来说，容器是一种非常好的方式，可以在很大程度上隔离的情况下运行代码，并且成本低廉，但这并不是谈话的结束。容器是一种技术，而不是一种用户体验。作为一个用户，我不想把集装箱推向生产，就像一个使用 amazon.com 的购物者不想打电话给码头安排货物装运一样。容器是构建在其上的一项非常棒的技术，但是我们不应该被移动机器映像的能力分散注意力，因为我们需要构建真正伟大的开发人员体验.

平台即服务(PaaS)构建在容器之上，比如 Cloud Foundry，从基于代码而非容器的用户体验开始。对于大多数开发人员来说，他们想要做的就是推送他们的代码并让它运行。在幕后，Cloud Foundry ——以及大多数其他 PaaSes ——使用该代码并创建一个可缩放和管理的容器化映像.在 Cloud Foundry 的例子中，它使用了一个 buildpack，但是您可以跳过这一步，也可以推出一个从 Dockerfile 创建的 Docker 映像。

使用 PaaS，容器的所有优势仍然存在——一致的环境、高效的资源管理等等——但是通过控制用户体验，PaaS 既可以为开发人员提供更简单的用户体验，又可以执行一些额外的技巧，比如在存在安全漏洞时修补根文件系统。更重要的是，平台提供了诸如数据库和消息队列之类的服务，您可以将其绑定到应用程序，从而消除了将所有内容都视为容器的需要。




