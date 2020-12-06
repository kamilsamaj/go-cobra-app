package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

func createHandler(cmd *cobra.Command, args []string) {
	fmt.Println("This is printed from a function in create")
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create something new",
	Long:  `This is a long description of create command`,
	Run:   createHandler,
}
