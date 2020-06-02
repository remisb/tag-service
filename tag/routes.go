package tag

import (
	"github.com/go-chi/chi"
	"net/http"
)

func InitRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", getTagsListHandler)
	r.Get("/{tagId}", getTagByTagIDHandler)
	r.Get("/linked/{tagId}", getLinkedTagsHandler)
	//r.Post("/", postTagHandler)
	//r.Post("/list", postTagListHandler)
	//r.Get("/groups", getTagGroupsHandler)
	//r.Post("/groups", postTagGroupsHandler)
	return r
}
