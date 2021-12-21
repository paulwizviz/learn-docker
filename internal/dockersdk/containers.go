package dockersdk

import (
	"context"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type PortMap struct {
	Exposed string
	Host    string
}

type VolMount struct {
	Source string
	Target string
}

func RunContainer(ctx context.Context, cli *client.Client, name string, image string, mnts []VolMount, portMaps []PortMap, networkName string) error {

	config := &container.Config{
		Image: image,
	}

	hostConfig := &container.HostConfig{}
	hostConfig.PortBindings = map[nat.Port][]nat.PortBinding{}

	for _, portMap := range portMaps {
		config.ExposedPorts = nat.PortSet{
			nat.Port(portMap.Exposed): {},
		}
		hostConfig.PortBindings[nat.Port(portMap.Exposed)] = append(hostConfig.PortBindings[nat.Port(portMap.Exposed)], nat.PortBinding{
			HostIP:   "0.0.0.0",
			HostPort: portMap.Host,
		})
	}

	for _, mnt := range mnts {
		mntPt := mount.Mount{}
		mntPt.Type = mount.TypeVolume
		mntPt.Source = mnt.Source
		mntPt.Target = mnt.Target
		hostConfig.Mounts = append(hostConfig.Mounts, mntPt)
	}
	log.Println(hostConfig.Mounts)

	netConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			"network": {
				NetworkID: networkName,
			},
		},
	}

	resp, err := cli.ContainerCreate(ctx, config, hostConfig, netConfig, nil, name)
	if err != nil {
		return err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	log.Println(resp.ID)

	return nil

}

type Health struct {
	Status        string
	FailingStreak int
	Log           []struct {
		Start    time.Time
		End      time.Time
		ExitCode int
		Output   string
	}
}

func ContainerHealth(ctx context.Context, cli *client.Client, name string) (Health, error) {
	result, err := cli.ContainerInspect(ctx, name)
	if err != nil {
		return Health{}, err
	}

	health := Health{}
	health.Status = result.State.Health.Status
	health.FailingStreak = result.State.Health.FailingStreak

	for _, l := range result.State.Health.Log {
		hl := struct {
			Start    time.Time
			End      time.Time
			ExitCode int
			Output   string
		}{
			Start:    l.Start,
			End:      l.End,
			ExitCode: l.ExitCode,
			Output:   l.Output,
		}
		health.Log = append(health.Log, hl)
	}

	return health, nil
}
