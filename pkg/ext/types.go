package ext

import "github.com/spf13/cobra"

type Plugin interface {
	Command() *cobra.Command
}
