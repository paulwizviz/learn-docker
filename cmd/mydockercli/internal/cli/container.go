package cli

import (
	"context"
	"log"
	"strings"

	"github.com/docker/docker/client"
	"github.com/paulwizviz/learn-docker/internal/dockersdk"
	"github.com/spf13/cobra"
)

var (
	containerRunCmd = &cobra.Command{
		Use:   "run",
		Short: "run container in background",
		Run: func(cmd *cobra.Command, args []string) {
			cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
			if err != nil {
				log.Fatal(err)
			}

			ports := []dockersdk.PortMap{}
			for _, portMap := range portMaps {
				p := strings.Split(portMap, ":")
				pm := dockersdk.PortMap{}
				pm.Exposed = p[1]
				pm.Host = p[0]
				ports = append(ports, pm)
			}

			mounts := []dockersdk.VolMount{}
			for _, volume := range volumes {
				v := strings.Split(volume, ":")
				mount := dockersdk.VolMount{}
				mount.Source = v[0]
				mount.Target = v[1]
				mounts = append(mounts, mount)
			}

			if err := dockersdk.RunContainer(context.Background(), cli, containerName, imageName, mounts, ports, network); err != nil {
				log.Fatal(err)
			}

		},
	}

	containerInspectCmd = &cobra.Command{
		Use:   "health",
		Short: "health of container",
		Run: func(cmd *cobra.Command, args []string) {
			cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
			if err != nil {
				log.Fatal(err)
			}

			result, err := dockersdk.ContainerHealth(context.Background(), cli, containerName)
			if err != nil {
				log.Fatal(err)
			}

			log.Println(result.Status)
			log.Println(result.FailingStreak)
			for _, l := range result.Log {
				log.Println("Start", l.Start)
				log.Println("End", l.End)
				log.Println("Output", l.Output)
				log.Println("Exit code", l.ExitCode)
			}
		},
	}

	containerCmd = &cobra.Command{
		Use:   "container",
		Short: "container related operations",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {

	containerRunCmd.Flags().StringVarP(&imageName, "image", "i", "", "image name")
	containerRunCmd.MarkPersistentFlagRequired("image")
	containerRunCmd.Flags().StringVarP(&containerName, "container", "c", "", "container name")
	containerRunCmd.MarkPersistentFlagRequired("container")
	containerRunCmd.Flags().StringArrayVarP(&portMaps, "publish", "p", nil, "Publish port")
	containerRunCmd.MarkPersistentFlagRequired("ports")
	containerRunCmd.Flags().StringArrayVarP(&volumes, "volume", "v", nil, "Volumes")
	containerRunCmd.MarkPersistentFlagRequired("source")
	containerRunCmd.Flags().StringVarP(&network, "network", "n", "", "Network")
	containerRunCmd.MarkPersistentFlagRequired("network")

	containerInspectCmd.Flags().StringVarP(&containerName, "container", "c", "", "containe name")
	containerRunCmd.MarkPersistentFlagRequired("container")

	containerCmd.AddCommand(containerRunCmd)
	containerCmd.AddCommand(containerInspectCmd)
}
