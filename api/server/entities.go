package server

import (
	"time"
)

//** Picture data **//

type Picture struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Format      string    `json:"format"`
	Source      string    `json:"source"`
	PictureURL  string    `json:"mainPictureURL"`
	PicturePath string    `json:"mainPicturePath"`
	CreatedDate time.Time `json:"createdDate"`
	EndDate     time.Time `json:"endDate"`
}

type GetPictureResponse struct {
	Status   string    `json:"status"`
	Total    int       `json:"total"`
	Pictures []Picture `json:"result"`
}

type AllPicturesResponse struct {
	Pictures []Picture `json:"pictures"`
}

//** Tag data **//

type Tag struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type AllTagsResponse struct {
	Status string `json:"status"`
	Total  int    `json:"total"`
	Tags   []Tag  `json:"result"`
}

// Pulse Response

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

//** ERROR **//

type ErrorResponse struct {
	Errors []string `json:"errors"`
}
