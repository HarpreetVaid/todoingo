/*
Copyright © 2025 Harpreet Singh harry798677@gmail.com
*/
package cmd

import (
	"fmt"
	"todolist/taskfunc"

	"github.com/spf13/cobra"
)

var updateID int

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long:  `Update an Task`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
		taskfunc.Update_task(cfgFile, updateID)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	updateCmd.Flags().IntVarP(&updateID, "id", "i", -1, "Id of the task to update")

}
