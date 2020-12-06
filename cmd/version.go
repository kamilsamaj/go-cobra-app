package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func versionHandler(cmd *cobra.Command, args []string) {
	fmt.Println("This is printed from a function")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of this app",
	Long:  `All software has versions. This is the app's version'`,
	Run:   versionHandler,
}
