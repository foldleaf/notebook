//消息 结构体
type Message struct {
	//发送消息，
	//接收消息，

	//目标地址
	Address string `json:"address"`
	//消息内容，无则仅建立连接
	Body string `json:"body"`
}

//var wg sync.WaitGroup
//参数-连接消息；返回-连接成功的消息
func (a *App) NewConnection(Msg string) {
	//wg.Add(1)
	fmt.Println(Msg)
	//创建一个传递字符串的通道
	ch := make(chan string, 2)
	//接收通道的值
	//fmt.Println("通道失败1")
	//recv:=<-ch
	fmt.Println("通道失败")

	go ConnectAndRead(Msg, ch)

	var recv string

	println("主程序运行")
	recv = <-ch

	fmt.Println("接收成功", recv)
	//wg.Wait()
	//return recv
	fmt.Println("程序未停止")
	/*每循环一次 返回一个值。

	  而不是把循环都结束后再返回哦。

	  我想，要么你把循环哪到最外面，把所有你打算在循环外处理的逻辑都放在循环内处理；
	  要么开两线程，一个线程做循环，每循环一次，把值放到一个变量，然后等另一个线程取这个变量的值，等另一个线程处理完了，它接着下次循环...这就复杂了
	*/

}

func ConnectAndRead(Msg string, ch chan string) string {
	//一、建立连接
	conn, err := net.Dial("tcp", Msg)

	if err != nil {
		fmt.Printf("连接失败，Error:%v", err)
		return ""
	}
	fmt.Println("连接")
	defer conn.Close()

	//二、通过连接发送数据
	//1.准备数据
	str := Msg
	//reply := ""

	//var reply []string
	//return reply

	//2.写数据到conn中,需要转换成byte数组,即发送消息
	_, err = conn.Write([]byte(str))
	if err != nil {
		fmt.Printf("发送失败，Error:%v", err)
		return ""
	}
	fmt.Println("写入")

	for {

		//三、接收响应数据,数据为byte数组
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("获取失败，Error:%v", err)
		}

		fmt.Println("原始响应" + string(buf[:n]))
		//转换成string,拼接
		//reply = reply + string(buf[:n])
		//数据读取完毕，跳出循环
		//fmt.Println("字符串" + reply)

		//创建传递字符串的通道
		//ch := make(chan string)
		//创建goruntine接收值
		//go RecvChannel(send)
		ch <- string(buf[:n])
		fmt.Println("发送成功")
		//fmt.Println(&ch)
		//wg.Done()

	}

}

//接收通道数据-字符串
// func RecvChannel(str chan string) string{
// 	recv:=<-str
// 	fmt.Println("接收成功",recv)
// 	return recv

// }

//新建与前端的tcp长连接-服务端
func TcpWithVue() {
	//指定监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("监听端口无法使用")
	}
	//监听连接
	for {
		//监听到连接，建立连接socket1 connVue
		connVue, err := listen.Accept()
		if err != nil {
			fmt.Println("监听到前端连接，但连接失败")
			continue
		}
		go progress(connVue) //启动一个goroutine处理连接
	}
}

func progress(connVue net.Conn) {
	defer connVue.Close()
	for {

		reader := bufio.NewReader(connVue)
		var buf [1024]byte
		//读取数据
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("读取前端数据失败")
			break
		}
		//将接收的数据转换为字符串，设备的连接地址
		deviceAddr := string(buf[:n])

		fmt.Println("接收到前端的数据", deviceAddr)

		//根据设备的连接地址与设备进行连接

		connDevice, err := net.Dial("tcp", deviceAddr)
		if err != nil {
			fmt.Printf("连接失败，Error:%v", err)

		}
		fmt.Println("连接")
		defer connDevice.Close()
		//向该socket2写数据，即向设备发送数据

		//向socket1写数据，即向前端发送数据
		connVue.Write()
	}

}
