package utils

import (
	"io/ioutil"
	"os"
)

func DoesDirExists(path string) bool {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true

}

func WriteFile(buffer []byte, path string) error {
	return ioutil.WriteFile(path, buffer, 0644)
}

func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
