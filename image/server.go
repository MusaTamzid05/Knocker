package image

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"knocker/docker_downloadpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type TCPServer struct {
	Address string
	Port    int
}

func NewTCPServer(address string, port int) *TCPServer {
	return &TCPServer{Address: address, Port: port}
}

type helperServer struct {
	docker_downloadpb.UnimplementedDockerDownloadServiceServer
}

func (e *helperServer) Download(ctx context.Context, req *docker_downloadpb.DockerDownloadRequest) (*docker_downloadpb.DockerDownloadResponse, error) {
	imageName := req.GetDockerDownload().GetImageName()

	imageFound := false
	buffer := []byte{}

	buffer, err := e.readFile(imageName)

	if err == nil {
		imageFound = true
	} else {
		err = errors.New("Image not found.")
	}

	res := &docker_downloadpb.DockerDownloadResponse{
		File:       buffer,
		ImageFound: imageFound,
		ImageName:  imageName,
	}

	return res, err

}

func (e *helperServer) LoadImage(imageName string) {
	e.readFile(imageName)
}

func (e *helperServer) readFile(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}

func (s *TCPServer) Start() {

	serverURL := fmt.Sprintf("%s:%d", s.Address, s.Port)
	listen, err := net.Listen("tcp", serverURL)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	docker_downloadpb.RegisterDockerDownloadServiceServer(grpcServer, &helperServer{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to server : %v", err)
	}

}
