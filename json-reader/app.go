package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	videos := getVideo()
	fmt.Println(videos)

	videos = getVideofromFile()
	fmt.Println(videos)

	for _, eachvideo := range videos {
		fmt.Println(eachvideo.Id)
		eachvideo.Description = eachvideo.Description + "\n Remember to like & Subscribe!" // this will not work because its a copy not actual
	}

	for i, _ := range videos {
		fmt.Println(videos[i].Id)
		videos[i].Description = videos[i].Description + "\n Remember to like & Subscribe!" // this will not work because its a copy not actual
	}
	fmt.Println(videos)

	saveVideo(videos)
}
