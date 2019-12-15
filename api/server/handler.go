package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Launching Home ...")
	fmt.Fprint(w, "Welcome\n- GET /pulse — heartbeat check if our API is online\n- GET /pictures — fetch all pictures from the database\n- GET /picture/[tag] — fetch pictures by tag from the database\n- GET /tags — fetch all tag available from the database\n")
	//Reset the Content-Type Header back to text
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// dat, err := ioutil.ReadFile("./web/api.html")
	// if err != nil {
	// 	log.Error().Err(err).Msg("Error opening/reading api.html file")
	// }
	// _, _ = fmt.Fprintf(w, string(dat))
}

func (server *Server) GetPicturesByTag(w http.ResponseWriter, r *http.Request) {
	var pictureList *[]Picture
	var errs []string

	vars := mux.Vars(r)
	id := vars["tag"]
	id = "/m/" + id
	pictureList = server.PictureByTag(id)
	if len(*pictureList) != 0 {
		json.NewEncoder(w).Encode(GetPictureResponse{Pictures: *pictureList})
	} else {
		errs = append(errs, "No picture found")
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}

func (server *Server) GetAllTags(w http.ResponseWriter, r *http.Request) {
	var labelList *[]Tag
	var objectList *[]Tag
	var errs []string

	labelList = server.AllLabels()
	objectList = server.AllObjects()
	tagList := append(*labelList, *objectList...)
	if len(*labelList) != 0 {
		json.NewEncoder(w).Encode(AllTagsResponse{Tags: tagList})
	} else {
		errs = append(errs, "No tags found")
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}

func (server *Server) GetAllLabels(w http.ResponseWriter, r *http.Request) {
	var labelList *[]Tag
	var errs []string

	labelList = server.AllLabels()
	if len(*labelList) != 0 {
		json.NewEncoder(w).Encode(AllTagsResponse{Tags: *labelList})
	} else {
		errs = append(errs, "No labels found")
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}

func (server *Server) GetAllObjects(w http.ResponseWriter, r *http.Request) {
	var objectList *[]Tag
	var errs []string

	objectList = server.AllObjects()
	if len(*objectList) != 0 {
		json.NewEncoder(w).Encode(AllTagsResponse{Tags: *objectList})
	} else {
		errs = append(errs, "No objects found")
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}

// func GetAllPictures(w http.ResponseWriter, r *http.Request) {
// 	var pictureList *[]Picture
// 	var errs []string
// }
