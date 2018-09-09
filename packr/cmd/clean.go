package cmd

import (
	"os"

	"github.com/gobuffalo/packr/costello/store"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "removes any *-packr.go files",
	RunE: func(cmd *cobra.Command, args []string) error {
		pwd, err := os.Getwd()
		if err != nil {
			return errors.WithStack(err)
		}
		args = append(args, pwd)
		for _, root := range args {
			if err := store.Clean(root); err != nil {
				return errors.WithStack(err)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}