package main

import "github.com/gowriswarupk/cobra-cli/pkg/pkg/commands"

func main() {
	rootCmd := commands.NewRootCMD()
	// Check those vaues
	_, _ = rootCmd.ExecuteC()
}

abcdefg
