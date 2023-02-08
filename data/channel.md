源码位置 src/runtime/chan.go

// channel 类型定义
type hchan struct {
    // channel 中的元素数量, len
    qcount   uint           // total data in the queue
    
    // channel 的大小, cap
    dataqsiz uint           // size of the circular queue
    
    // channel 的缓冲区，环形数组实现
    buf      unsafe.Pointer // points to an array of dataqsiz elements
    
    // 单个元素的大小
    elemsize uint16
    
    // closed 标志位
    closed   uint32
    
    // 元素的类型
    elemtype *_type // element type
    
    // send 和 recieve 的索引，用于实现环形数组队列
    sendx    uint   // send index
    recvx    uint   // receive index
    
    // recv goroutine 等待队列
    recvq    waitq  // list of recv waiters
    
    // send goroutine 等待队列
    sendq    waitq  // list of send waiters

    // lock protects all fields in hchan, as well as several
    // fields in sudogs blocked on this channel.
    //
    // Do not change another G's status while holding this lock
    // (in particular, do not ready a G), as this can deadlock
    // with stack shrinking.
    lock mutex
}

// 等待队列的链表实现
type waitq struct {    
    first *sudog       
    last  *sudog       
}

// in src/runtime/runtime2.go
// 对 G 的封装
type sudog struct {
    // The following fields are protected by the hchan.lock of the
    // channel this sudog is blocking on. shrinkstack depends on
    // this for sudogs involved in channel ops.

    g          *g
    selectdone *uint32 // CAS to 1 to win select race (may point to stack)
    next       *sudog
    prev       *sudog
    elem       unsafe.Pointer // data element (may point to stack)

    // The following fields are never accessed concurrently.
    // For channels, waitlink is only accessed by g.
    // For semaphores, all fields (including the ones above)
    // are only accessed when holding a semaRoot lock.

    acquiretime int64
    releasetime int64
    ticket      uint32
    parent      *sudog // semaRoot binary tree
    waitlink    *sudog // g.waiting list or semaRoot
    waittail    *sudog // semaRoot
    c           *hchan // channel
}
