package cli

import "github.com/spf13/cobra"

func NewInitCmd() *cobra.Command {
	var memoRootFlag string
	var articleRootFlag string

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize noq project",
		RunE: func(cmd *cobra.Command, args []string) error {
			input, err := CollectInitInput(memoRootFlag, articleRootFlag)
			if err != nil {
				return err
			}

			_ = input
			return nil
		},
	}

	cmd.Flags().StringVar(&memoRootFlag, "memo-root", "", "Path to memo root")
	cmd.Flags().StringVar(&articleRootFlag, "article-root", "", "Path to article root")
	return cmd
}
