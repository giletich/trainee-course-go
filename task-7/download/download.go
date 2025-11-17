package download

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"errors"
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
		err := errors.New(u.Scheme)
		return errors.Join(ErrInvalidProtocol, err)
	}

	resp, err := http.Get(fileURL)

	if err != nil {
		return ErrConnectionFailed
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		return ErrDownloadFailed
	}

	if resp.StatusCode == http.StatusNotFound {
		return ErrNotFound
	}

	if resp.StatusCode == http.StatusInternalServerError {
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