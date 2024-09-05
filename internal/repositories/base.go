package repositories

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	db *mongo.Database
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		db: db,
	}
}

func (r *Repositories) getCollection(name string) *mongo.Collection {
	return r.db.Collection(name)
}

func filterNotDeleted() bson.M {
	filter := bson.M{
		"$or": []bson.M{
			{"deleted_at": bson.M{"$exists": false}},
			{"deleted_at": nil},
		},
	}

	return filter
}
