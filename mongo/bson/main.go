package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gofer/pkg/log"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func main() {
	var (
		client *mongo.Client
		err    error
		db     *mongo.Database
		//collection *mongo.Collection
	)
	var ctx = context.Background()
	//1.建立连接
	if client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://time_machine:mxchip%402019@106.14.218.147:27017/time_machine").SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Print(err)
		return
	}
	//2.选择数据库 my_db
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	//3.选择表 my_collection
	db = client.Database("time_machine")
	//collection = collection
	//ops := options.Find().SetSort(bson.D{{"_id", -1}})

	ret, err := db.Collection("device_history").Find(context.TODO(), bson.D{{"iotId", "2bd5B0m1xDXLcRQjj17x000000"}})
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			log.Errorf("s.dao.Mongo.Collection(), err:%#v", err)
			return
		}
		log.Errorf("s.dao.Mongo.Collection(), err:%#v", err)
		return
	}
	result := make([]*DeviceInfo, 0)
	for ret.Next(context.TODO()) {
		dec := &DeviceInfo{}
		if err = ret.Decode(&dec); err != nil {
			log.Errorf("onlineLog cur.Decode err %+v", err)
			continue
		}
		result = append(result, dec)
	}
	log.Info("dasda", result)

}

type DeviceInfo struct {
	Iotid string `json:"iotId" bson:"iotId"`
}
