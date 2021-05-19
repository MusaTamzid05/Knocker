package image

import (
	"errors"
	"fmt"
	"knocker/utils"
)

type Image struct {
	imageDirPath string
}

func NewImage(imageDirPath string) (*Image, error) {

	if utils.DoesDirExists(imageDirPath) == false {
		message := fmt.Sprintf("Image dir does not exists path %s\n", imageDirPath)
		return nil, errors.New(message)
	}

	image := Image{imageDirPath: imageDirPath}
	return &image, nil
}

func (i *Image) List() []string {
	imageNames := []string{}
	return imageNames
}
