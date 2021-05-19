package image

type Image struct {
	imageDirPath string
}

func NewImage(imageDirPath string) *Image {
	image := Image{imageDirPath: imageDirPath}
	return &image
}

func (i *Image) List() []string {
	imageNames := []string{}
	return imageNames
}
