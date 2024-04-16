package gomongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Get(db, col string, filter interface{}, result interface{}, opts ...*options.FindOneOptions) error {
	collection := client.Database(db).Collection(col)
	return collection.FindOne(
		context.Background(), filter, opts...).Decode(result)
}

func Add(db, col string, obj interface{}, opts ...*options.InsertOneOptions) (id primitive.ObjectID, err error) {
	bObj, err := StructToDoc(obj)
	if err != nil {
		return
	}
	collection := client.Database(db).Collection(col)
	ctx := context.Background()
	InsertOneResult, err := collection.InsertOne(ctx, bObj, opts...)
	return InsertOneResult.InsertedID.(primitive.ObjectID), err
}

func Delete(db, col string, filter interface{}, opts ...*options.DeleteOptions) error {
	collection := client.Database(db).Collection(col)
	ctx := context.Background()
	_, err := collection.DeleteOne(ctx, filter, opts...)
	return err
}

func DeleteMany(db, col string, filter interface{}, opts ...*options.DeleteOptions) error {
	collection := client.Database(db).Collection(col)
	ctx := context.Background()
	_, err := collection.DeleteMany(ctx, filter, opts...)
	return err
}

// type of results must be array of struct
func List(db, col string, filter interface{}, pointerToSliceResults interface{}, pageNum, limitNum int, opts ...*options.FindOptions) error {
	page := int64(pageNum)
	limit := int64(limitNum)

	if page < 1 {
		page = 1
	}

	if 100 < limit || limit < 1 {
		limit = 100
	}

	skip := page*limit - limit
	opts = append(opts,
		options.Find().SetLimit(limit).SetSkip(skip))
	collection := client.Database(db).Collection(col)
	cur, err := collection.Find(context.Background(), filter, opts...)
	if err != nil {
		return err
	}
	defer cur.Close(context.Background())
	return cur.All(context.Background(), pointerToSliceResults)
}

func Aggregate(db, col string, filter interface{}, pointerToSliceResults interface{}, opts ...*options.AggregateOptions) error {
	collection := client.Database(db).Collection(col)
	cursor, err := collection.Aggregate(context.Background(), filter, opts...)
	if err != nil {
		return err
	}

	return cursor.All(context.Background(), pointerToSliceResults)
}

func UpdateOne(db, col string, filter, update interface{}, opts ...*options.UpdateOptions) error {
	collection := client.Database(db).Collection(col)
	_, err := collection.UpdateOne(context.Background(), filter, update, opts...)
	return err
}

func UpdateMany(db, col string, filter, update interface{}, opts ...*options.UpdateOptions) error {
	collection := client.Database(db).Collection(col)
	_, err := collection.UpdateMany(context.Background(), filter, update, opts...)
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

func StructToUnorderedDoc(v interface{}) (doc *bson.M, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
