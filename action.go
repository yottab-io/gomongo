package gomongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Get(db, col string, filter bson.D, result interface{}) error {
	collection := client.Database(db).Collection(col)
	return collection.FindOne(context.Background(), filter).Decode(result)
}

func Add(db, col string, obj interface{}) (id primitive.ObjectID, err error) {
	bObj, err := StructToDoc(obj)
	if err != nil {
		return
	}
	collection := client.Database(db).Collection(col)
	ctx := context.Background()
	InsertOneResult, err := collection.InsertOne(ctx, bObj)
	return InsertOneResult.InsertedID.(primitive.ObjectID), err
}

func Delete(db, col string, filter bson.D) error {
	collection := client.Database(db).Collection(col)
	ctx := context.Background()
	_, err := collection.DeleteOne(ctx, filter)
	return err
}
func DeleteMany(db, col string, filter bson.D) error {
	collection := client.Database(db).Collection(col)
	ctx := context.Background()
	_, err := collection.DeleteMany(ctx, filter)
	return err
}

// type of results must be array of struct
func List(db, col string, filter bson.D, results interface{}, page, limit int64) error {
	skip := page*limit - limit
	opt := options.FindOptions{Limit: &limit, Skip: &skip}
	collection := client.Database(db).Collection(col)
	cur, err := collection.Find(context.Background(), filter, &opt)
	if err != nil {
		return err
	}
	defer cur.Close(context.Background())
	return cur.All(context.Background(), results)
}

func UpdateOne(db, col string, filter bson.D, update interface{}) error {
	collection := client.Database(db).Collection(col)
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

func UpdateMany(db, col string, filter bson.D, update interface{}) error {
	collection := client.Database(db).Collection(col)
	_, err := collection.UpdateMany(context.Background(), filter, update)
	return err
}

func StructToDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
