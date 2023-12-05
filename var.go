package gomongo

import (
	"errors"
	"fmt"

	env "github.com/yottab-io/go_env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client         *mongo.Client
	dbUser         = env.Get("Yb_MONGODB_USER", "root")
	dbPass         = env.Get("Yb_MONGODB_PASS", "")
	dbPort         = env.GetInt64("Yb_MONGODB_PORT", 27017)
	dbDomain       = env.Get("Yb_MONGODB_DOMAIN", "0.0.0.0")
	dbAddress      = ""
	ErrBadReq      = errors.New("MongoDB Bad request")
	ClearFilter    = bson.D{}
	ErrNoDocuments = mongo.ErrNoDocuments
)

func init() {
	if len(dbPass) == 0 {
		dbAddress = fmt.Sprintf("mongodb://%s:%d", dbDomain, dbPort)
	} else {
		dbAddress = fmt.Sprintf("mongodb://%s:%s@%s:%d", dbUser, dbPass, dbDomain, dbPort)
	}
}
