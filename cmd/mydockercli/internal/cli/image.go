package cli

import (
	"context"
	"log"

	"github.com/docker/docker/client"
	"github.com/paulwizviz/learn-docker/internal/dockersdk"
	"github.com/spf13/cobra"
)

var (
	imagePullCmd = &cobra.Command{
		Use:   "pull",
		Short: "pull images",
		Run: func(cmd *cobra.Command, args []string) {
			cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
			if err != nil {
				log.Fatal(err)
			}
			dockersdk.PullImages(context.Background(), cli, imageName)
		},
	}

	imageCmd = &cobra.Command{
		Use:   "image",
		Short: "image related operations",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	imageCmd.PersistentFlags().StringVarP(&imageName, "name", "n", "", "image name")
	imageCmd.MarkPersistentFlagRequired("name")
	imageCmd.AddCommand(imagePullCmd)
}
