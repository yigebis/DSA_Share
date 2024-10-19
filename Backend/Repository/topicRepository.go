package Repository

import (
	"DSAShare/UseCase"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TopicRepository struct {
	Collection *mongo.Collection
	DbCtx      context.Context
}


func NewTopicRepository(collection *mongo.Collection, dbCtx context.Context) UseCase.ITopicRepository {
	return &TopicRepository{
		Collection: collection,
		DbCtx:      dbCtx,
	}
}

func (topr *TopicRepository) UpsertTopics(topics []string) error {
	var models []mongo.WriteModel

	for _, topic := range topics {
		update := bson.M{
			"$inc":         bson.M{"tag_count": 1},
			"$setOnInsert": bson.M{"name": topic},
		}

		filter := bson.M{"name": topic}

		updateModel := mongo.NewUpdateOneModel()
		updateModel.SetUpdate(update)
		updateModel.SetFilter(filter)
		updateModel.SetUpsert(true)

		models = append(models, updateModel)
	}

	bulkOptions := options.BulkWrite().SetOrdered(false) // Set unordered to run in parallel
	_, err := topr.Collection.BulkWrite(topr.DbCtx, models, bulkOptions)
	return err
}

func (topr *TopicRepository) IncrementTopicCount(topic string, by int) error {
	filter := bson.M{"name": topic}
	update := bson.M{
		"$inc": bson.M{"tag_count": by},
	}

	_, err := topr.Collection.UpdateOne(topr.DbCtx, filter, update)
	return err
}
func (topr *TopicRepository) DecrementTopicCount(topic string, by int) error {
	filter := bson.M{"name": topic}
	update := bson.M{
		"$inc":  bson.M{"tag_count": -by},
		"$pull": bson.M{"tag_count": bson.M{"$lte": 0}},
	}

	_, err := topr.Collection.UpdateOne(topr.DbCtx, filter, update)
	return err
}
