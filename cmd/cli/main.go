package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A CLI application for managing users and running commands",
	Long: `A CLI application that provides functionality to:
- Manage users (add, list, remove)
- Run custom commands
- Configure settings`,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
