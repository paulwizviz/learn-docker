package dockersdk

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

// Unix base volume
func NixVolOpt(name string, memSize int, uid int) volume.VolumeCreateBody {

	driverOps := map[string]string{
		"type":   "vfs",
		"device": "vfs",
		"o":      fmt.Sprintf("size=%v,uid=%v", memSize, uid),
	}

	return volume.VolumeCreateBody{
		Driver:     "local",
		DriverOpts: driverOps,
		Labels:     map[string]string{},
		Name:       name,
	}
}

func NewVol(ctx context.Context, cli *client.Client, name string) error {

	volBody := NixVolOpt(name, 256, 1000)

	vol, err := cli.VolumeCreate(ctx, volBody)
	if err != nil {
		return err
	}
	log.Println(vol)
	return nil
}

func ListVolumes(ctx context.Context, cli *client.Client) {

	filter := filters.Args{}
	vols, err := cli.VolumeList(ctx, filter)
	if err != nil {
		log.Println(err)
	}
	for _, vol := range vols.Volumes {
		log.Println(vol)
	}
}
