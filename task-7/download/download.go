package download

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var (
	ErrNotFound         = errors.New("resourse not found, check the url or availability of file")
	ErrInvalidURL       = errors.New("incorrect url, enter valid url")
	ErrInvalidProtocol  = errors.New("protocol is not http/https, try again: ")
	ErrConnectionFailed = errors.New("connection failed")
	ErrDownloadFailed   = errors.New("download failed, check if the file is available for downloading")
)

func DownloadFile(fileURL string, fileName string) error {

	u, err := url.ParseRequestURI(fileURL)
	if err != nil {
		return ErrInvalidURL
	}

	if u.Scheme != "https" && u.Scheme != "http" {
		return fmt.Errorf("%w: %s", ErrInvalidProtocol, u.Scheme)
	}

	resp, err := http.Get(fileURL)
	if err != nil {
		return ErrConnectionFailed
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusForbidden:
		return ErrDownloadFailed
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusInternalServerError:
		return ErrConnectionFailed
	}

	file, err := os.Create(fileName)
	if err != nil {
		return errors.New("create file error")
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return ErrDownloadFailed
	}

	return nil

}
