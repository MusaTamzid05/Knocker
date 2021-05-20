package image

import (
	"context"
	"knocker/docker_downloadpb"
	"knocker/utils"
	"log"
	"path/filepath"

	"google.golang.org/grpc"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SaveImage(res *docker_downloadpb.DockerDownloadResponse) error {
	imagePath := filepath.Join(IMAGE_DOWNLOAD_DIR, res.ImageName)
	return utils.WriteFile(res.File, imagePath)
}

func (c *Client) Run(serverAddress, imageName string) {
	cc, err := grpc.Dial(serverAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}

	defer cc.Close()

	grpClient := docker_downloadpb.NewDockerDownloadServiceClient(cc)

	req := &docker_downloadpb.DockerDownloadRequest{
		DockerDownload: &docker_downloadpb.DockerDownload{
			ImageName: imageName,
		},
	}

	res, err := grpClient.Download(context.Background(), req)

	if err != nil {
		log.Fatalf("Image not found %v\n", imageName)
	}

	log.Printf("Response from server : %v\n", res.ImageFound)
	err = c.SaveImage(res)

	if err != nil {
		log.Println("Error saving the image => ", err)
	}

}
