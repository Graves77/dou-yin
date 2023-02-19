package config

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

// MqConf RabbitMQ配置
type MqConf struct {
	Username       string
	Password       string
	Host           string
	Port           int
	ListenerQueues []QueueConf  `json:",optional"` // 需要监听点Queue列表
	Exchange       ExchangeConf `json:",optional"`
}

type ExchangeConf struct {
	Name       string
	Type       string
	Durable    bool `json:",default=false'"`
	AutoDelete bool `json:",default=false'"`
	Internal   bool `json:",default=false'"`
	NoWait     bool `json:",default=false'"`
}

type ConsumerConf struct {
	AutoAck   bool `json:",default=true"` // 消费时是否自动应答
	Exclusive bool `json:",default=false"`
	NoLocal   bool `json:",default=false"`
	NoWait    bool `json:",default=false"`
}

type BindConf struct {
	RouterKey string `json:",optional"`
	Exchange  string
	NotWait   bool `json:",default=false"`
}

type QueueConf struct {
	Name         string // 队列名称
	Durable      bool   `json:",default=false"`
	AutoDelete   bool   `json:",default=false"`
	Exclusive    bool   `json:",default=false"`
	NoLocal      bool   `json:",default=false"`
	NoWait       bool   `json:",default=false"`
	ConsumerConf ConsumerConf
	BindConf     BindConf `json:",optional"`
}

type Config struct {
	service.ServiceConf        // 基础配置
	Mq                  MqConf // RabbitMQ配置
	LikeRpc             zrpc.RpcClientConf
}

// GetUrl 获取amqp连接url
func (c *MqConf) GetUrl() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", c.Username, c.Password,
		c.Host, c.Port)
}
