package main

import (
	"fmt"
)

type video struct {
	Id          string
	Title       string
	Description string
	ImageURL    string
	Url         string
}

func getVideo() (videos []video) {

	video1 := video{
		Id:          "weqeqweqeqwe",
		Title:       "Azure of me",
		ImageURL:    "https://imageurl",
		Url:         "https://realusrl.url",
		Description: "Azure tutorial",
	}

	video2 := video{
		Id:          "weqeqweqeqwe",
		Title:       "Azure of me",
		ImageURL:    "https://imageurl",
		Url:         "https://realusrl.url",
		Description: "Azure tutorial",
	}

	return []video{video1, video2}
}
