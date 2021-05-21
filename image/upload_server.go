package image

import (
	"context"
	"fmt"
	"knocker/docker_uploadpb"
	"knocker/utils"
	"log"
	"net"
	"path/filepath"

	"google.golang.org/grpc"
)

type UploadServer struct {
	Address string
	Port    int
}

func NewUploadServer(Address string, Port int) *UploadServer {
	return &UploadServer{Address: Address, Port: Port}

}

type uploadHelperServer struct {
	docker_uploadpb.UnimplementedDockerUploadServiceServer
}

func (u *uploadHelperServer) Upload(ctx context.Context, req *docker_uploadpb.DockerUploadRequest) (*docker_uploadpb.DockerUploadResponse, error) {
	imageName := req.GetDockerUpload().GetImageName()
	buffer := req.GetDockerUpload().GetFile()
	imagePath := filepath.Join(IMAGE_SERVER_DIR, imageName)

	err := utils.WriteFile(buffer, imagePath)

	var success bool

	if err != nil {
		success = false
	} else {
		success = true

	}

	res := docker_uploadpb.DockerUploadResponse{
		Success: success,
	}

	return &res, err

}

func (u *UploadServer) Start() {
	serverURL := fmt.Sprintf("%s:%d", u.Address, u.Port)
	listen, err := net.Listen("tcp", serverURL)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	docker_uploadpb.RegisterDockerUploadServiceServer(grpcServer, &uploadHelperServer{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to server : %v", err)
	}

}
