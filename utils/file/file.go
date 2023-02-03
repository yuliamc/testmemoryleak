package file

import (
	"errors"
	"net/http"
	"os"
)

// Exists check file exist
func Exists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !(info.IsDir())
}

// Get remote file via HTTP request and then turn it into bytes
func GetRemoteFileInBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	// check response status code and response ContentLength
	if resp.StatusCode != http.StatusOK || resp.ContentLength < 0 {
		return nil, errors.New("REMOTE_FILE_INACESSIBLE")
	}

	content := make([]byte, resp.ContentLength)
	_, err = resp.Body.Read(content)
	defer resp.Body.Close()

	return content, nil
}

func GetRemoteFileSize(url *string) (*int64, error) {
	resp, err := http.Get(*url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK || resp.ContentLength < 0 {
		return nil, errors.New("REMOTE_FILE_INACESSIBLE")
	}

	return &resp.ContentLength, nil
}

func IsRemoteFileValidSize(url *string, validSize *int64) (*bool, error) {
	if remoteFileSize, err := GetRemoteFileSize(url); err != nil {
		return nil, err
	} else {
		isValid := *remoteFileSize < *validSize
		return &isValid, nil
	}
}
