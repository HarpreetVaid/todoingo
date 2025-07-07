/*
Copyright © 2025 NAME HERE harry798677@gmail.com
*/
package cmd

import (
	"fmt"
	"todolist/taskfunc"

	"github.com/spf13/cobra"
)

var deleteID int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  `Delete an Task`,
	Run: func(cmd *cobra.Command, args []string) {
		if taskfunc.Delete_task(cfgFile, deleteID) {
			fmt.Println("Succesfully Delted the Task")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	deleteCmd.Flags().IntVarP(&deleteID, "id", "i", -1, "Name of the task to add")

}
