package config

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
)

// MqConf RabbitMQ配置
type MqConf struct {
	Username       string
	Password       string
	Host           string
	Port           int
	ListenerQueues []ConsumerConf // 需要监听点Queue列表
}

type ConsumerConf struct {
	Name      string // 队列名称
	AutoAck   bool   `json:",default=true"` // 是否自动应答
	Exclusive bool   `json:",default=false"`
	NoLocal   bool   `json:",default=false"`
	NoWait    bool   `json:",default=false"`
}

type Config struct {
	service.ServiceConf        // 基础配置
	Mq                  MqConf // RabbitMQ配置
}

// GetUrl 获取amqp连接url
func (c *MqConf) GetUrl() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", c.Username, c.Password,
		c.Host, c.Port)
}
