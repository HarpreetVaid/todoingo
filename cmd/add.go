/*
Copyright © 2025 Harpreet Singh harry798677@gmail.com
*/
package cmd

import (
	"fmt"
	"todolist/taskfunc"

	"github.com/spf13/cobra"
)

var task string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add File",
	Long:  `Add File.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding a new Task ", task)
		success := taskfunc.Add_task(cfgFile, task)

		if success {
			fmt.Println("Sucessfully added Task")
		} else {
			fmt.Println("Something Went Wrong \n Please Try again")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringVarP(&task, "task", "t", "", "Name of the task to add")
}
