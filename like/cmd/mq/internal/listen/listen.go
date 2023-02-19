package listen

import (
	"awesomeProject/dou-yin/like/cmd/mq/internal/svc"
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/queue"
	"log"
)

type (
	ConsumeHandler func(ctx context.Context, svcCtx *svc.ServiceContext, data []byte) error

	RabbitListener struct {
		ctx     context.Context
		svcCtx  *svc.ServiceContext
		conn    *amqp.Connection
		channel *amqp.Channel
		forever chan bool
		handler ConsumeHandler
	}
)

func (q RabbitListener) Start() {
	if err := q.ExchangeDeclare(nil); err != nil {
		log.Fatalln(err)
	}
	if err := q.DeclareQueueAndBind(nil); err != nil {
		log.Fatalln(err)
	}

	for _, que := range q.svcCtx.Config.Mq.ListenerQueues {
		msg, err := q.channel.Consume(
			que.Name,
			"",
			que.ConsumerConf.AutoAck,
			que.ConsumerConf.Exclusive,
			que.ConsumerConf.NoLocal,
			que.ConsumerConf.NoWait,
			nil,
		)
		if err != nil {
			log.Fatalf("failed to listener, error: %v", err)
		}

		go func() {
			for d := range msg {
				if err := q.handler(q.ctx, q.svcCtx, d.Body); err != nil {
					logx.Errorf("Error on consuming: %s, error: %v", string(d.Body), err)
				}
			}
		}()
	}
	fmt.Println(q.svcCtx.Config.Name + " listening...")
	<-q.forever
}

func (q RabbitListener) Stop() {
	q.channel.Close()
	q.conn.Close()
	close(q.forever)
}

// MustNewListener 创建RabbitMQ监听者
func MustNewListener(ctx context.Context, svcCtx *svc.ServiceContext, handler ConsumeHandler) queue.MessageQueue {
	listener := RabbitListener{
		ctx:     ctx,
		svcCtx:  svcCtx,
		handler: handler,
		forever: make(chan bool),
	}
	conn, err := amqp.Dial(svcCtx.Config.Mq.GetUrl())
	if err != nil {
		log.Fatalf("failed to connect rabbitmq, error: %v", err)
	}

	listener.conn = conn
	channel, err := listener.conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel: %v", err)
	}

	listener.channel = channel
	return listener
}

// ExchangeDeclare 声明交换机
func (q RabbitListener) ExchangeDeclare(table amqp.Table) error {
	exchange := q.svcCtx.Config.Mq.Exchange
	return q.channel.ExchangeDeclare(
		exchange.Name,
		exchange.Type,
		exchange.Durable,
		exchange.AutoDelete,
		exchange.Internal,
		exchange.NoWait,
		table,
	)
}

func (q RabbitListener) DeclareQueueAndBind(table amqp.Table) error {
	if len(q.svcCtx.Config.Mq.ListenerQueues) == 0 {
		return nil
	}
	var err error
	for _, conf := range q.svcCtx.Config.Mq.ListenerQueues {
		_, err = q.channel.QueueDeclare(
			conf.Name,
			conf.Durable,
			conf.AutoDelete,
			conf.Exclusive,
			conf.NoWait,
			table,
		)
		err = q.bind(conf.Name, conf.BindConf.RouterKey, conf.BindConf.Exchange, conf.BindConf.NotWait, nil)
	}
	return err
}

func (q RabbitListener) bind(queueName string, RouterKey string, exchange string, notWait bool, table amqp.Table) error {
	return q.channel.QueueBind(
		queueName,
		RouterKey,
		exchange,
		notWait,
		table,
	)
}
