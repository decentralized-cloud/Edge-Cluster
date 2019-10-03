// Package cmd implements different commands that can be executed against EdgeCluster service
package cmd

import (
	"github.com/micro-business/go-core/pkg/util"
	"github.com/spf13/cobra"
)

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Get Edge Cluster CLI version",
		Run: func(cmd *cobra.Command, args []string) {
			util.PrintInfo("Edge Cluster CLI\n")
			util.PrintInfo("Copyright (C) 2019, Micro Business Ltd.\n")
			util.PrintYAML(util.GetVersion())
		},
	}
}
