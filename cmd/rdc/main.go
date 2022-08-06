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

	err := rdcCmd.Execute()
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
