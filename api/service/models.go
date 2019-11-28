package service

import (
	"time"
)

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

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
	Picture Picture `json:"Picture"`
}

type AllPicturesResponse struct {
	Pictures []Picture `json:"Pictures"`
}

//** Tag data **//

type Tag struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type AllTags struct {
	Tags []Tag `json:"Tags"`
}

type GetTagResponse struct {
	Picture Picture `json:"Picture"`
}

type AllTagsResponse struct {
	Pictures []Picture `json:"Pictures"`
}

//** ERROR **//

type ErrorResponse struct {
	Errors []string `json:"errors"`
}
