package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "Success", Code: 200})
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
	tag := "/" + id[:1] + "/" + id[1:]
	pictureList = server.PictureByTag(tag)
	if len(*pictureList) != 0 {
		json.NewEncoder(w).Encode(GetPictureResponse{Status: "Success", Total: len(*pictureList), Pictures: *pictureList})
	} else {
		errs = append(errs, "No picture found")
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}

func (server *Server) GetAllTags(w http.ResponseWriter, r *http.Request) {
	var labelList *[]Tag
	var objectList *[]Tag
	var tagList []Tag
	var errs []string
	var alreadyExist bool

	labelList = server.AllLabels()
	objectList = server.AllObjects()
	// tagList := append(*labelList, *objectList...)
	tagList = append(tagList, *labelList...)
	for _, object := range *objectList {
		alreadyExist = false
		for _, label := range *labelList {
			if label.Id == object.Id {
				alreadyExist = true
			}
		}
		if alreadyExist == false {
			tagList = append(tagList, object)
		}
	}
	if len(tagList) != 0 {
		json.NewEncoder(w).Encode(AllTagsResponse{Status: "Success", Total: len(tagList), Tags: tagList})
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
		json.NewEncoder(w).Encode(AllTagsResponse{Status: "Success", Total: len(*labelList), Tags: *labelList})
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
		json.NewEncoder(w).Encode(AllTagsResponse{Status: "Success", Total: len(*objectList), Tags: *objectList})
	} else {
		errs = append(errs, "No objects found")
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}

func (server *Server) GetPicturesFilteredByMultipleTags(w http.ResponseWriter, r *http.Request) {
	var resultList *[]Picture
	var errs []string

	key := r.FormValue("key")
	tags := strings.Split(key, ",")
	for _, tag := range tags {

		pictureList := server.PictureByTag("/" + tag[:1] + "/" + tag[1:])
		resultList = filterArray(pictureList, resultList)
	}
	if resultList != nil && len(*resultList) != 0 {
		json.NewEncoder(w).Encode(GetPictureResponse{Status: "Success", Total: len(*resultList), Pictures: *resultList})
	} else {
		errs = append(errs, "No picture found")
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}

func filterArray(new *[]Picture, old *[]Picture) *[]Picture {
	var resultList []Picture

	if old == nil || len(*old) == 0 {
		return new
	}
	for _, itemNew := range *new {
		for _, itemOld := range *old {
			if itemNew.Id == itemOld.Id && itemNew.Id != "" {
				resultList = append(resultList, itemNew)
			}
		}
	}
	return &resultList
}
