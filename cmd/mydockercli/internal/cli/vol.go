package cli

import (
	"context"
	"log"

	"github.com/docker/docker/client"
	"github.com/paulwizviz/learn-container/internal/dockersdk"
	"github.com/spf13/cobra"
)

var volCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create volume",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			log.Fatal(err)
		}
		dockersdk.NewVol(context.Background(), cli, "test")
	},
}

var volListCmd = &cobra.Command{
	Use:   "list",
	Short: "list volumes",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			log.Fatal(err)
		}
		dockersdk.ListVolumes(context.Background(), cli)
	},
}

var volCmd = &cobra.Command{
	Use:   "volume",
	Short: "volume related operations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	volCmd.AddCommand(volCreateCmd)
	volCmd.AddCommand(volListCmd)
}
