/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Fetch or add dependencies.",
	Args:    cobra.ArbitraryArgs, //TODO: check for valid formatting
	Aliases: []string{"add"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		logger.Debug("Test")
	},
	/*
	 Note: `dev get` will support semver constraints with @v
	 Note: and commit constraints with @git.
	 Note: Packages starting with @ will be compatible with
	 Note: npm-like formatting.
	 The format that `Args` will check for will support pretty
	 much everything, including `dev get`'s supported versioning
	 and individual package naming rules.
	 Each platform has its own rules on package names and version constraints.
	 The default behavior is to add the most recent known version.
	*/
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().BoolP("dry-run", "N", false, "Do not download any files")
	getCmd.Flags().BoolP("dev", "d", false, "Mark the listed dependencies as dev dependencies where supported")
	getCmd.Flags().Bool("optional", false, "Mark the listed dependencies as optional where supported")
}
