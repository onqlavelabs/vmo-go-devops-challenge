package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
    Use:   "migration",
    Short: "programming challenge migration command",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("programming challenge migration command called")
    },
}

func init() {
    rootCmd.AddCommand(migrationCmd)
}
