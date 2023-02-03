package file

import (
	"errors"
	"net/http"
)

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
