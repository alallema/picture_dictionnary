package service

import "time"

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

//** Tag data **//

type Tag struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type AllTags struct {
	Tags []Tag `json:"Tags"`
}

//** ERROR **//

type ErrorResponse struct {
	Errors []string `json:"errors"`
}
