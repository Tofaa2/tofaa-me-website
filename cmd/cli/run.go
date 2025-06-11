package main

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [command]",
	Short: "Run a custom command",
	Long:  `Execute a custom command with optional arguments`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]
		commandArgs := args[1:]

		// Execute the command
		execCmd := exec.Command(command, commandArgs...)
		output, err := execCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}

		fmt.Printf("Command output:\n%s\n", output)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
