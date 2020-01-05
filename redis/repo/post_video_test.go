package repo

import (
	"testing"

	"github.com/alallema/picture_dictionnary.git/core/service"
)

func TestPostVideo(t *testing.T) {
	type args struct {
		video service.Picture
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestPostCreateVideoLabel(t *testing.T) {
	type args struct {
		data service.VideoLabelData
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestPostCreateVideoCategory(t *testing.T) {
	type args struct {
		categories []service.Entity
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestPostLabelbyCategory(t *testing.T) {
	type args struct {
		label      service.VideoLabelData
		categories []service.Entity
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestPostSegbyVideobyLabel(t *testing.T) {
	type args struct {
		segments []service.LabelSegment
		data     service.VideoLabelData
		video    service.Picture
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestPostVideobyLabel(t *testing.T) {
	type args struct {
		video service.Picture
		data  service.VideoLabelData
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
