package tag

import (
	"context"
	"errors"
	"fmt"
	"github.com/remisb/tag-service/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
)

var ErrNoDocuments = errors.New("Tag: no documents in result")

var tagsMap Tags

type TagSet []int64

func (ts TagSet) contains(tagID int64) bool {
	for _, id := range ts {
		if id == tagID {
			return true
		}
	}
	return false
}

type SummaryTags struct {
	summaryOfferID primitive.ObjectID
	tags           TagSet
}

func getCollection() *mongo.Collection {
	return db.GetMongoCollection("summary")
}

func getTagsList() (*[]Tag, error) {
	var tagsList = make([]Tag, 0)
	for _, value := range tagsMap.ByID {
		tagsList = append(tagsList, value)
	}
	return &tagsList, nil
}

func getTagByTagIDString(tagId string) (*Tag, error) {
	id, err := strconv.ParseInt(tagId, 10, 64)
	if err == nil {
		fmt.Printf("%d of type %T", id, id)
	}
	return getTagByTagID(id)
}

func getTagByTagID(tagId int64) (*Tag, error) {
	if tag, ok := tagsMap.ByID[tagId]; !ok {
		return &tag, nil
	}
	return nil, ErrNoDocuments
}

func getAllTagSets() ([]*SummaryTag, error) {
	ctx, _ := db.GetTimeoutContext()
	// db.summary.find({}, {'tags.tagId': 1} ).limit(1).pretty()
	options := options.Find()
	options.SetProjection(bson.M{"tags.tagId": 1})
	emptyFilter := bson.M{}
	cursor, err := getCollection().Find(ctx, emptyFilter, options)
	if err != nil {
		return nil, err
	}
	tagSets, err := readSummaryTags(ctx, cursor)
	if err != nil {
		return nil, err
	}
	return tagSets, nil
}

// Cursor reader functions


func readSummaryTags(ctx context.Context, cursor *mongo.Cursor) ([]*SummaryTag, error) {
	defer cursor.Close(ctx)

	result := make([]*SummaryTag, 0)
	for cursor.Next(ctx) {
		summaryTag := SummaryTag{}
		err := cursor.Decode(&summaryTag)
		if err != nil {
			return result, err
		}
		result = append(result, &summaryTag)
	}
	return result, nil
}
