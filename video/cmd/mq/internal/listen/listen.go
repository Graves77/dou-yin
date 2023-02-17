package listen

import (
	"awesomeProject/dou-yin/video/cmd/mq/internal/svc"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/queue"
	"log"
)

type (
	ConsumeHandler func(svcCtx *svc.ServiceContext, message string) error

	RabbitListener struct {
		svcCtx  *svc.ServiceContext
		conn    *amqp.Connection
		channel *amqp.Channel
		forever chan bool
		handler ConsumeHandler
	}
)

func (q RabbitListener) Start() {
	for _, que := range q.svcCtx.Config.Mq.ListenerQueues {
		msg, err := q.channel.Consume(
			que.Name,
			"",
			que.AutoAck,
			que.Exclusive,
			que.NoLocal,
			que.NoWait,
			nil,
		)
		if err != nil {
			log.Fatalf("failed to listener, error: %v", err)
		}

		go func() {
			for d := range msg {
				if err := q.handler(q.svcCtx, string(d.Body)); err != nil {
					logx.Errorf("Error on consuming: %s, error: %v", string(d.Body), err)
				}
			}
		}()
	}

	<-q.forever
}

func (q RabbitListener) Stop() {
	q.channel.Close()
	q.conn.Close()
	close(q.forever)
}

// MustNewListener 创建RabbitMQ监听者
func MustNewListener(svcCtx *svc.ServiceContext, handler ConsumeHandler) queue.MessageQueue {
	listener := RabbitListener{
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
