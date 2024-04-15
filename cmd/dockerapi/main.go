package main

import (
	"fmt"
	"log"

	"github.com/paulwizviz/learn-docker/internal/dockerapi"
)

func main() {
	client := dockerapi.NewUnixSocketClient()
	images, err := dockerapi.ListImages(client)
	if err != nil {
		log.Fatal(err)
	}

	if len(images) == 0 {
		fmt.Println("There are no images")
		return
	}

	for _, image := range images {
		fmt.Println(image)
	}
}
