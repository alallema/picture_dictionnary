package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(server *Server) *mux.Router {

	var routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			server.Home,
		},
		Route{
			"Heartbeat",
			"GET",
			"/pulse",
			Heartbeat,
		},
		Route{
			"All Tags",
			"GET",
			"/tags",
			server.GetAllTags,
		},
		Route{
			"All Labels",
			"GET",
			"/labels",
			server.GetAllLabels,
		},
		Route{
			"All Objects",
			"GET",
			"/objects",
			server.GetAllObjects,
		},
		Route{
			"Pictures Result",
			"GET",
			"/picture/{tag}",
			server.GetPicturesByTag,
		},
		Route{
			"Video Result",
			"GET",
			"/video/{tag}",
			server.GetVideosByTag,
		},
		Route{
			"Categories Result",
			"GET",
			"/categories",
			server.GetAllCategories,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	router.
		Methods("GET").
		Path("/filteredtags/").
		Queries("key", "{[a-zA-Z0-9]*?}").
		Name("Pictures Filtered by Multiple Tags").
		Handler(http.HandlerFunc(server.GetPicturesFilteredByMultipleTags))
	return router
}
