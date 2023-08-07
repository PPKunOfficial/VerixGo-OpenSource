package Core

import "github.com/go-resty/resty/v2"

func DownloadFile(url string) ([]byte, error) {
	resp, err := resty.New().RemoveProxy().R().Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
