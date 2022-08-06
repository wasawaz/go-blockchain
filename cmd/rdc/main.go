package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	var rdcCmd = &cobra.Command{
		Use:   "rdc",
		Short: "The Roodic Blockchain CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	rdcCmd.AddCommand(versionCmd)
	rdcCmd.AddCommand(balanceCmd())
	rdcCmd.AddCommand(txCmd())

	err := rdcCmd.Execute()
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
