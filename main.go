package main

import (
	"flag"
	"knocker/image"
	"log"
)

func main() {

	imageServerPtr := flag.Bool("image_server", false, "The image server")
	imageDownloadClient := flag.Bool("image_client", false, "The image client")

	flag.Parse()

	if *imageServerPtr {
		log.Println("Running Image Server")
		tcpServer := image.NewTCPServer("0.0.0.0", 5000)
		tcpServer.Start()
		return

	}

	if *imageDownloadClient {
		imageClient := image.NewClient()
		imageClient.Run("0.0.0.0:5000", "main.go")
		return

	}

	flag.PrintDefaults()

}
