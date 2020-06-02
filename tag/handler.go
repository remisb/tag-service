package tag

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/remisb/tag-service/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strconv"
	"time"
)

func getLinkedTagsHandler(w http.ResponseWriter, r *http.Request) {
	tagID, err := getTagID(r)
	if err != nil {
		helper.SendJsonError(w, http.StatusBadRequest, err)
		return
	}

	linkedTags, err := getLinkedTags(tagID)
	if err != nil {
		helper.SendJsonError(w, http.StatusInternalServerError, err)
		return
	}
	helper.SendJson(w, http.StatusOK, linkedTags)
}

type SummaryTag struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Tags []*TagID           `json:"tags" bson:"tags"`
	set  TagSet
	//tags []interface{}           `json:"tags" bson:"tags"`
	//Tags []tags.Tag `json:"tags,omitempty" bson:"tags,omitempty"`
}

type Tag struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	TagID     int64              `json:"tagId" bson:"tagId"`
	Name      string             `json:"name" bson:"name"`
	GroupID   int64              `json:"groupId" bson:"groupId"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

type TagID struct {
	TagID int64 `json:"tagId" bson:"tagId"`
}

func getTagID(r *http.Request) (int64, error) {
	tagID := chi.URLParam(r, "tagId")
	if tagID == "" {
		err := errors.New("invalid request data tagId")
		return 0, err
	}
	id, err := strconv.ParseInt(tagID, 10, 64)
	if err != nil {
		return 0, errors.New("invalid request data format")
	}
	return id, nil
}

func getTagsListHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := getTagsList()
	if err != nil {
		log.Printf("error: ", err)
		helper.SendJsonError(w, http.StatusInternalServerError, err)
		return
	}
	helper.SendJsonOk(w, tags)
}

func getTagByTagIDHandler(w http.ResponseWriter, r *http.Request) {
	tagID := chi.URLParam(r, "tagId")
	if tagID == "" {
		helper.SendJsonError(w, http.StatusBadRequest, "invalid request data offerId")
		return
	}

	tag, err := getTagByTagIDString(tagID)
	if err != nil {
		helper.SendJsonError(w, http.StatusInternalServerError, err)
		return
	}
	helper.SendJsonOk(w, tag)
}
