// Package cmd implements different commands that can be executed against EdgeCluster service
package cmd

import (
	"github.com/decentralized-cloud/edge-cluster/pkg/util"
	"github.com/spf13/cobra"
)

func newStartCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the Edge Cluster service",
		Run: func(cmd *cobra.Command, args []string) {
			util.StartService()
		},
	}
}
