package service

import "time"

type LabelData struct {
	Id          string
	Language    string
	Description string
	Score       float32
	Confidence  float32
	Topicality  float32
}

type AllLabel struct {
	Tags []LabelData `json:"Tags"`
}

type LocalizedObjectData struct {
	Id           string
	Language     string
	Name         string
	Score        float32
	BoundingPoly BoundingPoly
}

type BoundingPoly struct {
	Vertices           []*Vertex
	NormalizedVertices []*NormalizedVertex
}

type NormalizedVertex struct {
	// X coordinate.
	X float32
	// Y coordinate.
	Y float32
}

type Vertex struct {
	// X coordinate.
	X int32
	// Y coordinate.
	Y int32
}

//** Picture data **//

type Picture struct {
	Id          string    `json:"id"`
	Title       string    `json:"Title"`
	Format      string    `json:"Format"`
	Source      string    `json:"Source"`
	PicturePath string    `json:"PicturePath"`
	PictureURL  string    `json:"PictureURL"`
	CreatedDate time.Time `json:"createdDate"`
}

//** ERROR **//

type ErrorResponse struct {
	Errors []string `json:"errors"`
}
