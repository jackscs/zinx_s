package ziface

type IMsgHandle interface {
	DoMsghandle(request IRequest)           //马上以非阻塞方式处理消息
	AddRouter(msgID uint32, router IRouter) //为消息添加具体的处理逻辑
	StartWorkerPool()                       // 启动工作线程池
	SendMsgToTaskQueue(request IRequest)    // 将消息交给TaskQueue,由worker进行处理
}
