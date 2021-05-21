package image

import (
	"context"
	"fmt"
	"knocker/docker_uploadpb"
	"knocker/utils"
	"log"
	"strings"

	"google.golang.org/grpc"
)

type UploadClient struct {
}

func NewUploadClient() *UploadClient {
	return &UploadClient{}
}

func (c *UploadClient) Run(serverAddress, imagePath string) {
	cc, err := grpc.Dial(serverAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}

	defer cc.Close()

	grpUploadClient := docker_uploadpb.NewDockerUploadServiceClient(cc)
	imageData := strings.Split(imagePath, "/")
	imageName := imageData[len(imageData)-1]

	buffer, err := utils.ReadFile(imagePath)

	if err != nil {
		log.Fatalln(err)
	}

	req := &docker_uploadpb.DockerUploadRequest{
		DockerUpload: &docker_uploadpb.DockerUpload{
			ImageName: imageName,
			File:      buffer,
		},
	}

	res, err := grpUploadClient.Upload(context.Background(), req)

	if err != nil {
		log.Fatalf("Image not found %v\n", imageName)
	}

	if res.Success {
		fmt.Println("Image uploaded")
	} else {
		fmt.Println("Image not uploaded")

	}

}
