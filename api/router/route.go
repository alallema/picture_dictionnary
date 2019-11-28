package router

import (
	"net/http"

	"github.com/alallema/picture_dictionnary.git/api/service"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		service.Home,
	},
	Route{
		"Heartbeat",
		"GET",
		"/pulse",
		service.Heartbeat,
	},
	// Route{
	//     "Pictures",
	//     "GET",
	//     "/pictures",
	//     Pictures.GetAllPictures,
	// },
	// Route{
	//     "Pictures",
	//     "GET",
	//     "/picture/{tag}",
	//     Pictures.GetPicturesByTag,
	// },
	// Route{
	// 	   "Pictures by operationId",
	// 	   "GET",
	// 	   "/tags",
	// 	   Pictures.GetAllTags,
	// },
}
