package tag

import (
	"fmt"
	"sort"
)

type TagUse struct {
	TagID   int64  `json:"tagId"`
	TagName string `json:"tagName,omitempty"`
	Count   int64  `json:"count"`
}

func getLinkedTags(tagID int64) ([]*TagUse, error) {
	allTagSets, err := getAllTagSets()
	if err != nil {
		return nil, err
	}

	// naive implementation without counting tag occurrences

	// more advance implementation where count of tag usage is counted
	tMap := map[int64]TagUse{}

	//containingTagSets := make([]TagUse, 0)
	for _, v := range allTagSets {
		if contains(v.Tags, tagID) {
			for _, id := range v.Tags {
				v , ok := tMap[id.TagID]
				if !ok {
					tMap[id.TagID] = TagUse{TagID: id.TagID , TagName: "", Count: 0}
					continue
				}
				v.Count++
				tMap[id.TagID] =  v
			}
		}
	}

	v := make([]*TagUse, 0)
	fmt.Println("Length:", len(tMap))
	for  _, value := range tMap {
		v = append(v, &value)
	}

	compares := 0
	// sort based on use count
	sort.Slice(v, func(i, j int) bool {
		compares++
		return v[i].Count > v[j].Count
	})
	fmt.Println("Compares: ", compares)
	return v, nil
}

func contains(set []*TagID, tagID int64) bool {
	for _, t := range set {
		if t.TagID == tagID {
			return true
		}
	}
	return false
}
