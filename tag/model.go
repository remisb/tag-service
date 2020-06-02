package tag

import (
	"errors"
)

//type Tag struct {
//	ID        primitive.ObjectID `json:"id" bson:"_id"`
//	TagID     int64              `json:"tagId" bson:"tagId"`
//	Name      string             `json:"name" bson:"name"`
//	GroupID   int64              `json:"groupId" bson:"groupId"`
//	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
//}

type Tags struct {
	LastTagID int64
	ByID      map[int64]Tag
}

func (t *Tags) incrementLastID() {
	t.LastTagID += 1
}

func (t *Tags) push(tag Tag) error {
	tagExist := t.ByID[tag.TagID]

	if tag.GroupID == 0 {
		tag.GroupID = 1 //other
	}

	if tagExist.Name != tag.Name && tagExist.GroupID != tag.GroupID {
		t.ByID[tag.TagID] = tag
		t.incrementLastID()
		return nil
	}
	return errors.New("tags: tag already exist")
}

func NewTags() Tags {
	t := Tags{
		LastTagID: 1,
		ByID:      make(map[int64]Tag),
	}
	return t
}
