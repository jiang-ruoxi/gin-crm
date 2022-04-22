package lib

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

// DialUrl 格式 amqp://账号：密码@rabbitmq服务器地址：端口号/vhost
const DialUrl = "amqp://myuser:mypass@127.0.0.1:5672/"

//RabbitMQ 定义rabbitMQ的结构体
type RabbitMQ struct {
	//连接
	Conn *amqp.Connection
	//channel表示AMQP通道
	Channel *amqp.Channel
	//queue队列名
	Queue string
	//exchange交换机
	Exchange string
	//routing key路由key
	RoutingKey string
	//Dial url连接的地址信息
	DialUrl string
}

//CreateRabbitMQ 创建RabbiMQ对象连接
func CreateRabbitMQ(queue, exchange, routingKey string) *RabbitMQ {
	rabbitmqObj := &RabbitMQ{
		Queue:      queue,
		Exchange:   exchange,
		RoutingKey: routingKey,
		DialUrl:    DialUrl,
	}
	var err error

	//Dial接受AMQP URI格式的字符串，并使用PlainAuth通过TCP返回新连接
	rabbitmqObj.Conn, err = amqp.Dial(rabbitmqObj.DialUrl)
	//因为该错误函数是RabbitMQ结构体的方法，所以只对该结构体调用有效
	rabbitmqObj.ErrFailInfoLog(err, "AMQP创建连接错误！")

	//打开通道,一个唯一的并发服务器通道来处理大量AMQP消息
	rabbitmqObj.Channel, err = rabbitmqObj.Conn.Channel()
	rabbitmqObj.ErrFailInfoLog(err, "AMQP打开channel失败！")

	return rabbitmqObj
}

//ErrFailInfoLog 错误处理函数
func (r *RabbitMQ) ErrFailInfoLog(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

//RabbitMQClose 关闭连接和channel
func (r *RabbitMQ) RabbitMQClose() {
	_ = r.Conn.Close()
	_ = r.Channel.Close()
}

//PublishData 数据内容写入生成者
func (r *RabbitMQ) PublishData(message string) error {
	//发送消息到队列中
	//Publish(交换器名称, 队列名称, 是否为无法路由的消息进行返回处理, 是否对路由到无消费者队列的消息进行返回处理, 消息体)
	if err := r.Channel.Publish(r.Exchange, r.Queue, false, false, amqp.Publishing{
		Headers:     amqp.Table{},
		Timestamp:   time.Now(),
		ContentType: "text/plain",
		Body:        []byte(message),
	}); err != nil {
		log.Printf("[amqp] publish message error: %s\n", err)
		return err
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " [result]:: message send successful")
	return nil
}

//DeclareExchange 创建或者连接一个交换器
func (r *RabbitMQ) DeclareExchange(kind string) error {
	//ExchangeDeclare(交换机名称, 交换机类型, 是否持久化, 是否自动删除, 是否是内置交换机, 是否等待服务器确认,args Table)
	if err := r.Channel.ExchangeDeclare(r.Exchange, kind, true, false, false, true, nil); err != nil {
		log.Printf("[amqp] declare exchange error: %s\n", err)
		return err
	}
	return nil
}

//DeclareQueue 创建或者连接一个队列
func (r *RabbitMQ) DeclareQueue() error {
	//QueueDeclare(队列名称, 是否持久化, 是否自动删除, 是否排他, 是否等待服务器确认, args Table)
	_, err := r.Channel.QueueDeclare(r.Queue, false, false, false, false, nil)
	if err != nil {
		log.Printf("[amqp] declare queue error: %s\n", err)
		return err
	}
	return nil
}

//BindQueue 绑定交换器和队列
func (r *RabbitMQ) BindQueue(key string, nowait bool) error {
	// QueueBind(队列名称, 根据交换机类型来设定, 交换机名称 string, 是否等待服务器确认, args Table)
	if err := r.Channel.QueueBind(r.Queue, key, r.Exchange, nowait, nil); err != nil {
		log.Printf("[amqp] bind queue error: %s\n", err)
		return err
	}
	return nil
}

func (r *RabbitMQ) ConsumeQueue() error {
	// 接收消息
	msgChan, err := r.Channel.Consume(r.Queue, "", true, false, false, false, nil)
	if err != nil {
		log.Printf("[amqp] consume queue error: %s\n", err)
		return err
	}

	forever := make(chan bool)
	//go 协程方式处理数据
	go func() {
		for d := range msgChan {
			//log.Printf是以控制台输出的方式进行打印展示
			log.Printf("Received a message: %s", d.Body)

			//这里需要根据业务的需求处理对应的业务逻辑
			DealConsumeData(d.Body)
		}
	}()
	log.Printf("[*] Waiting for message, To exit press CTRL+C")
	<-forever
	return nil
}

//NewRabbitMQSimple 创建简单模式实例对象
func NewRabbitMQSimple(queue string) *RabbitMQ {
	return CreateRabbitMQ(queue, "", "")
}

/**
 * 在 Golang 中创建 rabbitmq 生产者基本步骤是：
 * (1)连接 Connection、(2)创建 Channel、(3)创建或连接一个交换器、(4)创建或连接一个队列、(5)交换器绑定队列、(6)投递消息、(7)关闭 Channel、(8)关闭 Connection
 */

//SimplePublishData 简单模式下的生产者
func (r *RabbitMQ) SimplePublishData(message string) {
	defer r.RabbitMQClose()
	//前两步骤已在创建对象连接的时候已经实现，现在需要从第三步骤开始即可
	//创建或者连接一个交换器
	//_ = r.DeclareExchange("direct")

	//创建或者连接一个队列
	_ = r.DeclareQueue()

	//绑定交换器和队列
	//_=r.BindQueue("direct", true)

	//发送消息到队列中
	_ = r.PublishData(message)
}

//SimpleConsumeData 简单模式消费者
func (r *RabbitMQ) SimpleConsumeData() {
	defer r.RabbitMQClose()
	// 接收消息
	// Consume(队列名称, 消费者, 是否自动确认, 是否具有排他性, noLocal, 队列消费是否阻塞, 其他配置)
	// 保证队列存在
	_ = r.DeclareQueue()

	// 接收消息
	_ = r.ConsumeQueue()
}

//People e.g. 案例示例
type People struct {
	Name string
	Addr string
	Age  string
	Sex  string
}

//DealConsumeData 处理消费数据的方法
func DealConsumeData(data []byte) {
	//定义一个student实例
	var st People

	err := json.Unmarshal(data, &st)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("这里是处理后的数据::%v \n", st)
}
