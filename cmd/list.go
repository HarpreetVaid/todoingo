/*
Copyright © 2025 Harpreet Singh harry798677@gmail.com
*/
package cmd

import (
	"fmt"
	"todolist/taskfunc"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `List an Task`,
	Run: func(cmd *cobra.Command, args []string) {
		taskfunc.List_task(cfgFile)
		fmt.Println("List task is called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
