package internal

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/docker/docker/api/types/build"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/go-connections/nat"
	"github.com/moby/go-archive"
	"github.com/moby/term"
)

func Deploy(ctx context.Context, path string) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	buildContext, err := archive.TarWithOptions(path, &archive.TarOptions{})
	if err != nil {
		log.Fatalf("creating image build context: %v", err)
	}

	fmt.Println("Generating image...")

	imageBuildOps := build.ImageBuildOptions{
		Tags:   []string{"example-app"},
		Remove: true,
	}

	image, err := cli.ImageBuild(ctx, buildContext, imageBuildOps)
	if err != nil {
		panic(err)
	}
	defer image.Body.Close()

	fd, isTerminal := term.GetFdInfo(os.Stdout)
	if err := jsonmessage.DisplayJSONMessagesStream(image.Body, os.Stdout, fd, isTerminal, nil); err != nil {
		panic(err)
	}

	fmt.Println("Image generated!")

	fmt.Println("Running image...")

	contPort, err := nat.NewPort("tcp", "8081")
	if err != nil {
		panic(err)
	}

	containerConfig := &container.Config{
		Image:        "example-app",
		ExposedPorts: nat.PortSet{contPort: struct{}{}},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			contPort: []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "8081",
				},
			},
		},
	}

	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println("You app is running!")
}
