package containers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetContainerStats(containerId string) (*json.Decoder, error) {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.43"))
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctx, cancel := context.WithCancel(context.Background())

	stats, err := cli.ContainerStats(ctx, containerId, false)
	if err != nil {
		return nil, err
	}
	defer stats.Body.Close()

	var containerStats types.StatsJSON
	decoder := json.NewDecoder(stats.Body)

	if err := decoder.Decode(&containerStats); err != nil {
		fmt.Printf("failed to decode json")
	}

	rss := containerStats.MemoryStats.Stats["rss"]

	fmt.Println("rss:", rss)

	cancel()
	return decoder, err
}
