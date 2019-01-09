package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"log"
	"time"
)

type logTime struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}
type logRecord struct {
	TaskName string  `bson:"taskName"`
	Command  string  `bson:"command"`
	Err      string  `bson:"err"`
	Content  string  `bson:"content"`
	LogTime  logTime `bson:logTime`
}

type findCond struct {
	TaskName string `bson:"taskName"`
}
type timeBeforeCond struct {
	Before int64 `bson:"$lt"`
}
type delCondition struct {
	beforTime timeBeforeCond `bson:"logTime.StartTime"`
}

func main() {
	var (
		client       *mongo.Client
		err          error
		database     *mongo.Database
		collection   *mongo.Collection
		logRec       *logRecord
		insertResult *mongo.InsertOneResult
		insertManyResult *mongo.InsertManyResult
		docId        primitive.ObjectID
		cond		 *findCond
		cursor		mongo.Cursor
		findOpt	     *options.FindOptions
		delCond 	*delCondition
		delResult   *mongo.DeleteResult
	)
	if client, err = mongo.Connect(context.TODO(), "mongodb://localhost:27017"); err != nil {
		log.Println(err)
	}

	database = client.Database("cronjob")
	collection = database.Collection("log")

	logRec = &logRecord{
		TaskName: "Job10",
		Command:  "echo Hello",
		Err:      "",
		Content:  "Hello",
		LogTime: logTime{
			StartTime: time.Now().Unix(),
			EndTime:   time.Now().Unix() + 10,
		},
	}

	if insertResult, err = collection.InsertOne(context.TODO(), logRec); err != nil {
		log.Println(err)
	}
	docId = insertResult.InsertedID.(primitive.ObjectID)
	fmt.Println(docId.Hex())

	//Twitter OSS: snowflake: current time with minisecond to generate + machine ID
	if insertManyResult, err = collection.InsertMany(context.TODO(),[]interface{}{logRec,logRec,logRec});err!=nil{
		log.Println(err)
		return
	}
	for _, value := range insertManyResult.InsertedIDs{
		docId = value.(primitive.ObjectID)
		fmt.Println(docId.Hex())
	}

	cond = &findCond{TaskName:"Job10"}
	findOpt = &options.FindOptions{}
	if cursor, err = collection.Find(context.TODO(),cond,findOpt.SetSkip(0),findOpt.SetLimit(10)); err!=nil{
		log.Println(err)
		return
	}
	defer  cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		rec := &logRecord{}
		if err = cursor.Decode(rec); err!=nil{
			log.Println(err)
		}
		fmt.Println("rec", rec)
	}

	delCond = &delCondition{beforTime:timeBeforeCond{Before:time.Now().Unix()}}
	if delResult, err = collection.DeleteMany(context.TODO(),delCond); err!=nil{
		log.Println(err)
		return
	}
	fmt.Println("Delete count:", delResult.DeletedCount)


}
