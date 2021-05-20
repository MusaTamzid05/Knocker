package main

import (
	"flag"
	"log"
)

func main() {

	imageServerPtr := flag.Bool("image_server", false, "The image server")
	imageDownloadClient := flag.Bool("image_client", false, "The image client")

	flag.Parse()

	if *imageServerPtr {
		log.Println("Running Image Server")
		return

	}

	if *imageDownloadClient {
		log.Println("Running Image Client")
		return

	}

	flag.PrintDefaults()

}
