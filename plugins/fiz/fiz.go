package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var plugin2Cmd = &cobra.Command{
	Use:   "fiz",
	Short: "fiz plugin",
	Long:  "implement me",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("plugin2 called")
	},
}

type plugin struct{}

func (plugin) Command() *cobra.Command {
	return plugin2Cmd
}

var Plugin plugin
