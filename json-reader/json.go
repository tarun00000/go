package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type video struct {
	Id          string
	Title       string
	Description string
	ImageURL    string
	Url         string
}

func getVideo() (_ []video) {

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

func getVideofromFile() (videos []video) {

	fileByte, err := ioutil.ReadFile("./video.json")

	if err != nil {
		panic(err)
	}

	fileContent := string(fileByte)
	fmt.Println(fileContent)
	err = json.Unmarshal(fileByte, &videos)

	if err != nil {
		panic(err)
	}

	return videos
}

func saveVideo(videos []video) {
	videoByte, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./video.json", videoByte, 0644)

	if err != nil {
		panic(err)
	}
}
