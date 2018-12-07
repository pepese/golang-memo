package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "A root CLI written in Go.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pepese")
	},
}

var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "A root CLI written in Go.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sub")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	rootCmd.AddCommand(subCmd)
	rootCmd.Execute()
}
