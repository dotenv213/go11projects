package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myweather",
	Short: "A fast weather CLI tool",
	Long:  `myweather is a CLI tool built with Go and Cobra to fetch weather data.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}