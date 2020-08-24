package db

import (
	"fmt"
	"context"
	"reflect"
	"errors"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/seheee/PDK/logic-core/domain/model"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/seheee/PDK/logic-core/setting"
)

type logicRepository struct {
	client *mongo.Client
	collection *mongo.Collection
}

func NewLogicRepository() *logicRepository {
	uri := "mongodb://"+setting.MongoDbSetting.Address + ":" +setting.MongoDbSetting.Port
	clientOptions := options.Client().ApplyURI(uri)

	cli, err := mongo.Connect(context.TODO(), clientOptions)
	fmt.Println("\nresult type:", reflect.TypeOf(cli))
	if err != nil {
		fmt.Println("connect error: ", err.Error())
	}

	err = cli.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println("connect error: ", err.Error())
	}

	fmt.Println("Connected to MongoDB!")
	col := cli.Database("test").Collection("logics")

	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"logicname": 1}, Options: opt}
	if _, err := col.Indexes().CreateOne(context.TODO(), index); err != nil {
		fmt.Println("Could not create index:", err)
	}

	return &logicRepository{
		client: cli,
		collection: col,
	}
}

func (lr *logicRepository) GetAll() (r []model.Ring, err error) {
	r = make([]model.Ring,0)
	cur, err := lr.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println("find error:", err.Error())
		return nil, err
	}
	
	for cur.Next(context.TODO()) {
		var elem model.Ring
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("decode error: ", err.Error())
		}
		r = append(r, elem)
	}

	return r, err
}

func (lr *logicRepository) Create(r *model.RingRequest) (string, error) {
	result, err := lr.collection.InsertOne(context.TODO(),r)
	if err != nil {
		return "", errors.New("duplicate logic name")
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
} 

func (lr *logicRepository)Delete(id string) error {
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	} 
	res, err := lr.collection.DeleteOne(context.TODO(), bson.M{"_id": idPrimitive})
	if err != nil {
		fmt.Println("delete error:", err.Error())
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no logic")
	}
	return nil
}