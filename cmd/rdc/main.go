package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rdcCmd = &cobra.Command{
	Use:   "rdc",
	Short: "The Roodic Blockchain CLI",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func main() {

	err := rdcCmd.Execute()
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
