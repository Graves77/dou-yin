package mongodb

import (
	"awesomeProject/dou-yin/common/viperconfigread"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type autoIncrement struct {
	Name  string `bson:"name"`
	Value int64  `bson:"value"`
}

var MongoDatabase *mongo.Database

func init() {
	mongoConfig, err := viperconfigread.ConfigReadToMongoDB()
	if err != nil {
		logx.Errorf("[pkg]mongodb [func]init [msg]get MongoDB config, [err]%v", err)
	}
	url := fmt.Sprintf("mongodb://%v:%v@%v:%v", mongoConfig.MongoUserName, mongoConfig.MongoPwd, mongoConfig.MongoUrl, mongoConfig.MongoPort)
	mongoDB := mongoConfig.MongoDB
	MongoDatabase, err = MongoDBConnect(mongoDB, url)
	if err != nil {
		logx.Errorf("[pkg]mongodb [func]init [msg]get MongoDB Database, [err]%v", err)
	}
}

func MongoDBConnect(database string, url string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logx.Errorf("[pkg]mongodb [func]MongoDBConnect [msg]connect mongodb failed, [err]%v", err)
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		logx.Errorf("[pkg]mongodb [func]MongoDBConnect [msg]ping mongodb failed, [err]%v", err)
		return nil, err
	}

	mongoDatabase := client.Database(database)
	return mongoDatabase, nil
}
