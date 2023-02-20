package viperconfigread

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/core/logx"
)

type MySQLConfig struct {
	UserName        string //账号
	PassWord        string //密码
	Host            string //数据库地址，可以是Ip或者域名
	Port            int32  //数据库端口
	DBname          string //数据库名
	TimeOut         string //连接超时，10秒
	MaxIdleConns    int    //最大空闲连接数
	MaxOpenConns    int    //最大连接数
	ConnMaxLifetime int
}

type MinioConfig struct {
	Endpoint        string // 用的API端口，不是Console
	AccessKeyID     string // 账户名
	SecretAccessKey string // 密码
	UseSSL          bool   // 是否开启SSL加密
	BucketName      string // 桶的名字
}

type MongoConfig struct {
	MongoUserName string
	MongoPwd      string
	MongoUrl      string
	MongoPort     int
	MongoDB       string
}

// 获取当前可执行文件位置
func getExeFile() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		logx.Errorf("%v %v", exePath, err)
		return "", err
	}
	// 返回上级目录
	yamlFile := filepath.Dir(filepath.Dir(filepath.Dir(exePath)))
	// 这个路径是手动指定的，经常需要修改
	configFile := fmt.Sprintf("%s/bin/config/databaseconfig.yaml", yamlFile)
	logx.Infof("yamlFile:%v outputDir:%v", yamlFile, configFile)
	return configFile, nil
}

// 通过Viper获取配置文件
func getViperConfig(configFileName string) (*viper.Viper, error) {
	configFile, _ := getExeFile()
	viper.SetConfigFile(configFile)
	content, err := os.ReadFile(configFile)
	if err != nil {
		logx.Errorf("ioutil err:%v", err)
		return nil, err
	}
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		logx.Errorf("viper.ReadConfig: %v", err)
		return nil, err
	}
	config := viper.Sub(configFileName)
	return config, nil
}

// 获取MySQL的配置文件
func ConfigReadToMySQL() (*MySQLConfig, error) {
	config, err := getViperConfig("MySQL")
	if err != nil {
		logx.Errorf("ConfigReadToMySQL: %v", err)
		return nil, err
	}
	mysqlConfig := &MySQLConfig{
		UserName:        config.GetString("UserName"),
		PassWord:        config.GetString("PassWord"),
		Host:            config.GetString("Host"),
		Port:            config.GetInt32("Port"),
		DBname:          config.GetString("DBname"),
		TimeOut:         config.GetString("TimeOut"),
		MaxIdleConns:    config.GetInt("MaxIdleConns"),
		MaxOpenConns:    config.GetInt("MaxOpenConns"),
		ConnMaxLifetime: config.GetInt("ConnMaxLifetime"),
	}
	return mysqlConfig, nil
}

// 获取Minio的配置文件
func ConfigReadToMinio() (*MinioConfig, error) {
	config, err := getViperConfig("Minio")
	if err != nil {
		logx.Errorf("ConfigReadToMinio: %v", err)
		return nil, err
	}
	minioConfig := &MinioConfig{
		Endpoint:        config.GetString("Endpoint"),
		AccessKeyID:     config.GetString("AccessKeyID"),
		SecretAccessKey: config.GetString("SecretAccessKey"),
		UseSSL:          config.GetBool("UseSSL"),
		BucketName:      config.GetString("BucketName"),
	}
	return minioConfig, nil
}

func ConfigReadToMongoDB() (*MongoConfig, error) {
	config, err := getViperConfig("MongoDB")
	if err != nil {
		logx.Errorf("ConfigReadToMongoDB: %v", err)
	}
	mongoClient := &MongoConfig{
		MongoUserName: config.GetString("mongoUserName"),
		MongoPwd:      config.GetString("mongoPwd"),
		MongoUrl:      config.GetString("mongoUrl"),
		MongoPort:     config.GetInt("mongoPort"),
		MongoDB:       config.GetString("mongoDatabase"),
	}
	return mongoClient, nil
}
