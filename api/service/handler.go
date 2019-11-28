package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome\n- GET /pulse — heartbeat check if our API is online\n- GET /pictures — fetch all pictures from the database\n- GET /picture/[tag] — fetch pictures by tag from the database\n- GET /tags — fetch all tag available from the database\n")
}

// func GetPicturesByTags(w http.ResponseWriter, r *http.Request) {
// 	var pictureRecommandation []int
// 	var errs []string
// }

// func GetAllPictures(w http.ResponseWriter, r *http.Request) {
// 	var pictureList *[]Picture
// 	var errs []string
// }

// func GetAllTags(w http.ResponseWriter, r *http.Request) {
// 	var tagList *[]Tag
// 	var errs []string
// }
