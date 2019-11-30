package service

type LabelData struct {
	Mid         string
	Locale      string
	Description string
	Score       float32
	Confidence  float32
	Topicality  float32
}

type LocalizedObjectData struct {
	Mid          string
	LanguageCode string
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
