package containers

import (
	"context"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func GetImages() ([]image.Summary, error) {
	apiClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	images, err := apiClient.ImageList(context.Background(), image.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	return images, nil
}
