package service

import (
	"strings"
)

// Format 1 Picture 2 Video
// 	"JPEG"  Format = 1
// 	"PNG8"  Format = 1
// 	"PNG24" Format = 1
// 	"GIF"	Format = 1
// 	"BMP"  	Format = 1
// 	"WEBP"  Format = 1
// 	"RAW"  	Format = 1
// 	"ICO"  	Format = 1
// 	"PDF"  	Format = 1
// 	"TIFF"  Format = 1
// 	"MOV"  	Format = 2
// 	"MPEG4" Format = 2
// 	"MP4"  	Format = 2
// 	"AVI"  	Format = 2

func CheckFormat(name string) int {

	nameSplit := strings.Split(name, ".")
	lenName := len(nameSplit)
	if lenName >= 1 {
		formatName := strings.ToUpper(nameSplit[lenName-1])
		if formatName == "JPEG" || formatName == "JPG" || formatName == "PNG" || formatName == "PNG8" || formatName == "PNG24" || formatName == "GIF" || formatName == "BMP" || formatName == "WEBP" || formatName == "RAW" || formatName == "ICO" || formatName == "PDF" || formatName == "TIFF" {
			return 1
		}
		if formatName == "MOV" || formatName == "MPEG" || formatName == "MPEG4" || formatName == "MP4" || formatName == "AVI" {
			return 2
		}
	}
	return 0
}
