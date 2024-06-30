package main

// 有三个角色： broker,consumer, producer

/*
	broker的功能：1.初始化MongoDB  2.初始化Redis  3. 接收请求
	consumer的功能：1.连接到MongoDB  2.登录验证  3.注册消费者组  4.请求新消息
	producer的功能：1.连接到MongoDB  2.验证登录  4.向broker发消息

	根据上面的功能，大致分为：1.数据访问层(封装MongoDB和Redis的交互逻辑)
							2.业务逻辑层(登录验证，注册消费者组，消息处理逻辑)
							3.API层(接收外部请求，如HTTP。调用业务逻辑层的服务处理这些请求)
							4.可以单独实现一个安全层，专门处理安全性和身份认证

*/
