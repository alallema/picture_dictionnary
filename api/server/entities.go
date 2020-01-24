package server

import (
	"time"
)

type Response struct {
	Status   string		`json:"statusText"`
	Code	 int		`json:"status"`
}

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
	// Status   string    `json:"status"`
	Response Response  `json:"response"`
	Total    int       `json:"total"`
	Pictures []Picture `json:"result"`
}

type AllPicturesResponse struct {
	Pictures []Picture `json:"pictures"`
}

//** Tag data **//

type Tag struct {
	Id    string	`json:"id"`
	Title string	`json:"title"`
	Type  string	`json:"type"`
}

type AllTagsResponse struct {
	// Status string `json:"status"`
	Response	Response	`json:"response"`
	Total		int  		`json:"total"`
	Tags		[]Tag		`json:"result"`
}
