package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{}
	rootCmd.AddCommand(Run())
	rootCmd.Execute()
}
func Run() *cobra.Command {
	var help = cobra.Command{
		Use:   "run",
		Short: "Get help",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			file, _ := cmd.Flags().GetString("f")
			if file != "" {
				fmt.Println("Found file")
			} else {
				fmt.Printf("Hello %s\n", args[0])
			}

		},
	}
	return &help
}
