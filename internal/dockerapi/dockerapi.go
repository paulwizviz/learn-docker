package dockerapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

func NewUnixSocketClient() http.Client {
	return http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/var/run/docker.sock")
			},
		},
	}
}

func NewTCOSocketClient() http.Client {
	return http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("tcp", "0.0.0.0:3001")
			},
		},
	}
}

type Image struct {
	ID          string      `json:"Id"`
	ParentID    string      `json:"ParentId"`
	RepoTags    []string    `json:"RepoTags"`
	RepoDigests []string    `json:"RepoDigests"`
	Created     int         `json:"Created"`
	Size        int         `json:"Size"`
	Labels      interface{} `json:"Labels"`
	Containers  int         `json:"Containers"`
}

func ListImages(client http.Client) ([]Image, error) {

	res, err := client.Get("http://v1.40/images/json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var images []Image
	err = json.Unmarshal(body, &images)
	if err != nil {
		return nil, nil
	}

	return images, nil
}

func PullImages(client http.Client, imageName string, tag string) (io.ReadCloser, error) {

	queryParams := fmt.Sprintf("?fromImage=%s&fromSrc=-&tag=%s", imageName, tag)
	res, err := client.Post(fmt.Sprintf("http://v1.40/images/create%s", queryParams), "application/json", nil)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
