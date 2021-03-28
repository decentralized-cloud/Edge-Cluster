package main

import (
	"fmt"

	"github.com/decentralized-cloud/edge-cluster/internal/cmd"
	"github.com/micro-business/go-core/pkg/util"
)

func main() {
	fmt.Println("Test")
	rootCmd := cmd.NewRootCommand()
	util.PrintIfError(rootCmd.Execute())
}
