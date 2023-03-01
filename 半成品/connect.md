# 简介
1. 与目标设备进行tcp连接，接受目标设备发送的数据并存储。
2. 能随意建立复数连接并分别存储对应接收的数据数据。
3. 发送命令可以取消指定的连接。

# 准备
## mongodb
本人使用 docker 进行安装使用 mongodb，你也可以在[mongodb官网](https://www.mongodb.com/atlas/database)下载安装。
***
首先安装 docker，推荐可以使用docker管理工具 [Docker Desktop](https://www.docker.com/products/docker-desktop/)，windows下的安装可参考[Docker 简介和安装](https://docker.easydoc.net/doc/81170005/cCewZWoN/lTKfePfP)。
linux下请选择对应发行版的方法，ArchLinux的安装可以参考:
```bash
yay -S docker
# 可选：linux 下的 docker 管理工具
yay -S lazydocker
# 启动 docker 服务
sudo systemctl start docker
```
***
安装完docker后，使用docker 拉取 mongo 镜像并在名为 mongodb 的容器中运行，指定运行端口为20717
```bash
docker pull mongo
docker run -d -p 20717:20717 --name mongodb -d mongo
```
访问localhost:20717，出现以下信息则表明服务启动成功
```txt
It looks like you are trying to access MongoDB over HTTP on the native driver port.
```
## Gin
https://gin-gonic.com/zh-cn/docs/quickstart/
```bash
# 安装 gin ，这是一个 web 框架
go get -u github.com/gin-gonic/gin
# 安装跨域中间件 cors，如果有需要的话
go get github.com/gin-contrib/cors
```
## mongo-driver
https://github.com/mongodb/mongo-go-driver
```bash
# 安装 mongo-driver，用于操作 mongodb
go get go.mongodb.org/mongo-driver/mongo
```
# 路由设置
/router
```go
// 路由初始化
func InitRouter() {
	r := gin.Default()
	// 允许跨域
	r.Use(cors.Default())
	// 路由
	r.POST("connect/:ip",api.NewConnect)
	r.POST("disconnect/:ip",api.Disconnect)

	//在本地 20000 端口运行
	r.Run(":20000")
}
```
# 数据模型
/model
数据模型：
```go
// 消息结构体
type Msg struct {
	// 获取数据的时间，不是设备产生数据的时间
	Time string 
	// 连接的设备 ip
	Ip   string
	// 设备发送的数据，原始数据未解析
	Data string 	
}
```
获取数据库连接对象：
```go
// 初始化数据库，返回 *mongo.Client 用于操作数据库
func InitDB(ctx *gin.Context) *mongo.Client{
	ip:=ctx.Param("ip")
	// 设置客户端连接配置
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接 mongodb，获取连接对象
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		fmt.Println(ip,":连接mongo失败",err)
	}
	fmt.Println(ip,":连接mongo成功")

	return client
}
```
# tcp连接
## 思路
1. 用到可缓冲channel的阻塞机制，当channel为空时，读取channel会阻塞，当传入数据才能读取，解除阻塞。
2. 还有 context 控制 goroutine ，不过感觉有些多余。
3. 用全局 map 保存 tcp 的连接状态，ip 对应一个控制该 ip 连接的 channel 。创建ip连接时会创建一个对应的 k/v 对
4. 当传入一个请求时，程序调用连接函数运行一个 goroutine，该 goroutine 会根据请求的`ip`参数创建一个 tcp 连接，读取数据并写入数据库。
5. 如果继续向下执行就会取消 goroutine， 关闭连接，但我们使用读空channel的操作`<-channel`来阻塞，直到执行取消函数时将向channel写入数据以解除阻塞，退出 goroutine
6. 然后将该 ip/channel 对删除
7. 能够准确控制 goroutine，也不会产生内存泄露，可喜可贺，可喜可贺
## 代码 
定义一个连接状态map
```go
// key为ip地址，value为控制该 ip 连接阻塞的 channel
var MapIpChan=make(map[string]chan int)
```
连接并读取数据;可以考虑把这里面的 go 协程单独抽出来做一个函数。
```go
func NewConnect(ctx *gin.Context) {
	// 根据请求的参数获取 ip
	ip := ctx.Param("ip")
	// 创建一个数据库连接对象
	dbclient:=model.InitDB(ctx)
	// 需要操作的数据库的集合对象
	collection:=dbclient.Database("msg").Collection(ip)
	// 有缓冲的 channel，阻塞该ip的cancel()，当 disconnect执行时，取消阻塞，使cancel()执行
	cancelChan := make(chan int, 1)
	model.MapIpChan[ip] = cancelChan
	// 可取消的 context，使用cancel()取消，依赖于该context的 goroutine也会停止
	ctxConn, cancel := context.WithCancel(ctx)
	// 启动一个goroutine，建立tcp连接，获取传来的数据
	go func(ip string) {
		fmt.Println("准备连接ip:",ip)
		// GetConnect(ip1)
		conn, err := net.Dial("tcp", ip)
		if err != nil {
			fmt.Println("连接失败:", err)
			// 退出
			return
		}
		buf := make([]byte, 1024)
		for {
			select {
			case <-ctxConn.Done():
				// cancel()执行时关闭设备连接，关闭数据库连接，退出该 goroutine
				conn.Close()
				fmt.Printf("-----关闭设备%v的tcp连接-----\n",ip)
				dbclient.Disconnect(ctx)
				fmt.Printf("+++++关闭操作数据库%v的连接+++++\n",ip)
				fmt.Printf("*****退出连接%v的协程*****\n",ip)
				return
			default:
				fmt.Println("连接后准备读取数据")
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("已连接，读取错误")
				}
				// 编码转换,在utils包下创建转码函数，然后调用
				data:=utils.Transcode(buf[:n])
				fmt.Printf("%s: %s\n", ip, data)
				// 数据对象				
				msg:=model.Msg{
					Time: time.Now().String(),
					Ip: ip,
					Data: data,
				}
				// 插入数据
				insertResult,err:=collection.InsertOne(ctx,msg)
				if err != nil {
					fmt.Println("插入数据失败")
				}
				fmt.Println(msg.Data,"插入数据成功:",insertResult.InsertedID)
			}
		}
	}(ip)
	// 从该ip对应的 channel接收数据，由于初始为空则阻塞，
	// 当 disconnect函数执行时向 channel写入数据，解除阻塞，向下执行
	<-cancelChan
	// 删除该 ip/channel 对
	delete(model.MapIpChan, ip)
	fmt.Printf("准备取消任务:%v\n", ip)
	cancel()
	ctx.JSON(200, gin.H{
		"message": "已关闭连接",
	})
}
```
取消连接
```go
func Disconnect(ctx *gin.Context) {
	// 根据请求的参数获取 ip
	ip := ctx.Param("ip")
	// 在执行 NewConnect() 时会根据 ip 创建键值对，启动 goroutine 建立连接后，NewConnect()被阻塞
	// NewConnect() 等待 Disconnect() 取消阻塞后关闭其启动的 goroutine，退出
	// 判断 ip 是否在 map 中，即是否处于连接状态
	if _, isOk := model.MapIpChan[ip]; isOk {
		// channel为空时接收会导致阻塞，向channel发送数据使其不为空则取消阻塞
		model.MapIpChan[ip] <- 1
	} else {
		fmt.Println("ip未连接")
	}
	fmt.Println("解除阻塞")
}
```
这个for-select条件判断是关键，当channel为空时，无法读取，case条件不满足，所以运行default的functionB();
向channel发送数据后，channel可以读取，case条件满足，运行functionA()
```go
for{
    select {
        case:<-channel
            functionA()
        default:
            functionB()
    }
}
```
<-ctxConn.Done()就是在cancel()执行后解除阻塞。
# 测试
api测试工具访问，如果你的服务端会自动发数据的话，使用connect会创建tcp连接，将接受的数据会在控制台打印，并写入mongodb；
使用disconnect接口，会断开对应已有的tcp连接
POST:http://localhost:20000/connect/192.168.6.66:6666
POST:http://localhost:20000/disconnect/192.168.6.66：6666
