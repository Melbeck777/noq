package cli

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	root := &cobra.Command{Use: "noq"}

	root.AddCommand(NewInitCmd())
	return root
}
