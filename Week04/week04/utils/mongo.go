package utils

import (
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"week04/public"
)

//var MongoEngine *mgo.Database
var MongoEngine *mongo.Database
var MongoClient *mongo.Client
func InitMongo(ConfName string) {
	//连接数据库
	var err error
	Conf := InitConfig(ConfName)

	//需要验证

	Password := public.EncodeURIComponent(Conf.Databases.Password)
	dbConf := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authMechanism=SCRAM-SHA-1",Conf.Databases.Username,Password,Conf.Databases.Host,Conf.Databases.Port,Conf.Databases.Dbname)
	//Session,err:= mgo.Dial("mongodb://root:123456@188.131.192.242:27017")
	//dbConf := fmt.Sprintf("mongodb://%s:%s",Conf.Databases.Host,Conf.Databases.Po

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	Client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbConf))
	MongoClient = Client
	if err != nil{
		fmt.Println("mgo.Dial Error:",err)
		return
	}
	MongoEngine = MongoClient.Database("week04",nil)


}


