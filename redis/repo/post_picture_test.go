package repo

import (
	"github.com/google/go-cmp/cmp"
	// "time"
	core "github.com/alallema/picture_dictionnary.git/core/service"
	"testing"
)

func Test_PostAndGetPicture(t *testing.T) {
	tests := []struct {
		name    string
		want    core.Picture
		wantErr bool
	}{
		{
			name:    "simple picture",
			want:    core.Picture{
				Id          : "19461154",
				Title       : "img_657.jpg",
				Format      : "jpg",
				Source      : "instagram:Lala",
				PicturePath : "/path/img_657.jpg",
				// CreatedDate : time.Now(),
				},

			wantErr: false,
		},
		{
			name:    "picture already in redis",
			want:    core.Picture{
				Id          : "19461154",
				Title       : "img_657.jpg",
				Format      : "jpg",
				Source      : "instagram:Lala",
				PicturePath : "/path/img_657.jpg",
				// CreatedDate : time.Now(),
				},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostPicture(tt.want)
			got := GetPicture(tt.want.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PostAndGetLabelByPicture(t *testing.T) {
	tests := []struct {
		name    string
		haveP   core.Picture
		haveL	core.LabelData
		want	[]string
		wantErr	bool
	}{
		{
			name:    "simple label",
			haveP:    core.Picture{
				Id          : "19461154",
				},
			haveL:	  core.LabelData{
				Id			: "/m/0bt9lr",
				Score		: 0.98,
			},
			want:	 []string{"/m/0bt9lr"},
			wantErr: false,
		},
		{
			name:    "second label same picture",
			haveP:    core.Picture{
				Id          : "19461154",
				},
			haveL:	  core.LabelData{
				Id			: "/m/0b9ags",
				Score		: 0.34,
			},
			want:	 []string{"/m/0b9ags", "/m/0bt9lr"},
			wantErr: false,
		},
		{
			name:    "another label same picture",
			haveP:    core.Picture{
				Id          : "19461154",
				},
			haveL:	  core.LabelData{
				Id			: "/m/0b8ajh",
				Score		: 0.76,
			},
			want:	 []string{"/m/0b9ags", "/m/0b8ajh", "/m/0bt9lr"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostLabelByPicture(tt.haveP, tt.haveL)
			got := GetLabelByPicture(tt.haveP.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PostAndGetLocalizedObjectByPicture(t *testing.T) {
	tests := []struct {
		name    string
		haveP   core.Picture
		haveL	core.LocalizedObjectData
		want	[]string
		wantErr	bool
	}{
		{
			name:    "simple object",
			haveP:    core.Picture{
				Id          : "19461154",
			},
			haveL:	  core.LocalizedObjectData{
				Id			: "/m/0bt9lr",
				Score		: 0.98,
			},
			want:	 []string{"/m/0bt9lr"},
			wantErr: false,
		},
		{
			name:    "second object same picture",
			haveP:    core.Picture{
				Id          : "19461154",
			},
			haveL:	  core.LocalizedObjectData{
				Id			: "/m/0b9ags",
				Score		: 0.34,
			},
			want:	 []string{"/m/0b9ags", "/m/0bt9lr"},
			wantErr: false,
		},
		{
			name:    "another object same picture",
			haveP:    core.Picture{
				Id          : "19461154",
			},
			haveL:	  core.LocalizedObjectData{
				Id			: "/m/0b8ajh",
				Score		: 0.76,
			},
			want:	 []string{"/m/0b9ags", "/m/0b8ajh", "/m/0bt9lr"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostLocalizedObjectByPicture(tt.haveP, tt.haveL)
			got := GettLocalizedObjectByPicture(tt.haveP.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
		})
	}
	Client.ZRemRangeByScore("pictureIdLabel:19461154", "-inf", "+inf")
	Client.HDel("picture:19461154", "id", "title", "format", "source", "picturePath", "pictureURL")
	Client.ZRemRangeByScore("pictureIdObject:19461154", "-inf", "+inf")
}

func Test_PostCreateLabel(t *testing.T) {
	tests := []struct {
		name    string
		have   	core.LabelData
		want	string
		wantErr	bool
	}{
		{
			name:    "simple object",
			have:    core.LabelData{
				Id          : "/m/0bt9lr",
				Language	: "en",
				Description : "test",
			},
			want:	 "test",
			wantErr: false,
		},
		{
			name:    "second object same LabelData",
			have:    core.LabelData{
				Id          : "/m/0b9ags",
				Language	: "en",
				Description : "test2",
			},
			want:	 "test2",
			wantErr: false,
		},
		{
			name:    "another object same LabelData",
			have:    core.LabelData{
				Id          : "/m/32kja0",
				Language	: "en",
				Description : "test3",
			},
			want:	 "test3",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostCreateLabel(tt.have)
			got := GetLabelDescription(tt.have.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
			Client.HDel("labelDescr:"+tt.have.Id, "en")
		})
	}
}

func Test_PostCreateObject(t *testing.T) {
	tests := []struct {
		name    string
		have   	core.LocalizedObjectData
		want	string
		wantErr	bool
	}{
		{
			name:    "simple object",
			have:    core.LocalizedObjectData{
				Id          : "/m/0bt9lr",
				Language	: "en",
				Name 		: "test",
			},
			want:	 "test",
			wantErr: false,
		},
		{
			name:    "second object same ObjectData",
			have:    core.LocalizedObjectData{
				Id          : "/m/0b9ags",
				Language	: "en",
				Name 		: "test2",
			},
			want:	 "test2",
			wantErr: false,
		},
		{
			name:    "another object same ObjectData",
			have:    core.LocalizedObjectData{
				Id          : "/m/32kja0",
				Language	: "en",
				Name 		: "test3",
			},
			want:	 "test3",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PostCreateObject(tt.have)
			got := GetObjectDescription(tt.have.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("result = %v, want %v", got, tt.want)
			}
			Client.HDel("objectDescr:"+tt.have.Id, "en")
		})
	}
}