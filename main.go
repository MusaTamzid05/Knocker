package main

import (
	"flag"
	"knocker/image"
	"log"
)

func main() {

	imageServerPtr := flag.Bool("image_server", false, "The image server")
	imageDownloadClient := flag.Bool("image_client", false, "The image client")
	imageUploadServerPtr := flag.Bool("image_upload_server", false, "The image upload server")
	imageUploadClientPtr := flag.Bool("image_upload_client", false, "The image upload client")

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

	if *imageUploadServerPtr {
		log.Println("Running Image upload Server")
		imageUploadServer := image.NewUploadServer("0.0.0.0", 5001)
		imageUploadServer.Start()
		return
	}

	if *imageUploadClientPtr {
		log.Println("Running Image upload client")
		imageUploadClient := image.NewUploadClient()
		imageUploadClient.Run("0.0.0.0:5001", "main.go")
		return
	}

	flag.PrintDefaults()

}
