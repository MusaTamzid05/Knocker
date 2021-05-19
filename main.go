package main

import (
	"knocker/image"
	"log"
)

func main() {
	image, err := image.NewImage("./images")

	if err != nil {
		log.Fatalln(err)
	}
	image.List()
}
