package service

import (
	"github.com/golang/protobuf/ptypes/duration"
	"time"
)

//** Vision for Picture Data Structure **//

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

//** Vision for Video Data Structure **//

type VideoLabelData struct {
	Entity         Entity
	CategoryEntity []Entity
	Segments       []LabelSegment
	Frames         []LabelFrame
}

type Entity struct {
	Id          string
	Description string
	Language    string
}

type LabelSegment struct {
	Segment    Segment
	Confidence float32
}

type Segment struct {
	StartTimeOffset *duration.Duration
	EndTimeOffset   *duration.Duration
}

type LabelFrame struct {
	TimeOffset time.Time
	Confidence float32
}

//** ERROR **//

type ErrorResponse struct {
	Errors []string `json:"errors"`
}
