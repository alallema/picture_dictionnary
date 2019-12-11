package server

import (
	"time"
)

//** Picture data **//

type Picture struct {
	Id          int       `json:"id"`
	Title       string    `json:"Title"`
	Format      string    `json:"Format"`
	Source      string    `json:"Source"`
	PictureURL  string    `json:"mainPictureURL"`
	CreatedDate time.Time `json:"createdDate"`
	EndDate     time.Time `json:"endDate"`
}

type GetPictureResponse struct {
	Pictures []Picture `json:"Picture"`
}

type AllPicturesResponse struct {
	Pictures []Picture `json:"Pictures"`
}

//** Tag data **//

type Tag struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type AllTagsResponse struct {
	Tags []Tag `json:"Pictures"`
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
