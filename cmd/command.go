/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// commandCmd represents the command command
var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "A brief description of your command",
	Long:  `This is the sample "command"`,
	Run: func(cmd *cobra.Command, args []string) {
		inputValue, _ := cmd.Flags().GetString("sampleFlag")

		fmt.Println(inputValue)
		fmt.Println("command called")
	},
}

func init() {
	rootCmd.AddCommand(commandCmd)
	commandCmd.PersistentFlags().String("sampleFlag", "defaultValue", "Description for Flag")
}
