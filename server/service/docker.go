package service

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerClient struct{}

var dockerClient *client.Client
var once sync.Once

func GetDockerClient() *client.Client {
	if dockerClient == nil {
		once.Do(func() {
			cli, err := client.NewClientWithOpts(client.FromEnv)
			dockerClient = cli
			if err != nil {
				log.Print("create docker client failed: ", err)
			}
		})
		return dockerClient
	}
	return dockerClient
}

func GetContainers() ([]types.Container, error) {

	cli := GetDockerClient()
	return cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
}

func ContainerStart(containerID string, options types.ContainerStartOptions) error {

	cli := GetDockerClient()
	return cli.ContainerStart(context.Background(), containerID, options)
}

func ContainerStop(containerID string, timeout *time.Duration) error {

	cli := GetDockerClient()
	return cli.ContainerStop(context.Background(), containerID, timeout)
}
