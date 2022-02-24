package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

//
func NewRootCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "root",
		Short:   "I do vacumm",
		Long:    "I do vacum",
		Example: "spyder vaccum",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("I do vacumm")
		},
	}

	// cmd.AddCommand(
	//    //NewClusterCommandVaccum(),
	// )

	return cmd
}
