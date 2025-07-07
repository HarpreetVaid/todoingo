/*
Copyright © 2025 NAME HERE harry798677@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "todolist",
	Short: "A Simple ToDo List For Strating",
	Long:  `A Simple ToDO List to use this use  You Can Use add list update delete your Task`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("WelCome in the ToDo List Root")
		fmt.Println("Config File is", cfgFile)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Path to config file")
	cobra.OnInitialize(loadPreviousConfig)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getConfigPath() string {
	usr, _ := user.Current()
	return filepath.Join(usr.HomeDir, ".todolist_last_config")
}

func writeLastUsedConfig(path string) {
	_ = os.WriteFile(getConfigPath(), []byte(path), 0644)
}

func readLastUsedConfig() string {
	data, err := os.ReadFile(getConfigPath())
	if err != nil {
		return ""
	}
	return string(data)
}

func loadPreviousConfig() {
	if cfgFile != "" {
		// User provided --config, so save it
		writeLastUsedConfig(cfgFile)
		fmt.Println("New Config File is added and will be used.")
		fmt.Println("Path is ", cfgFile)
		return
	}

	// If not provided, try to load last used config
	last := readLastUsedConfig()
	if last != "" {
		fmt.Println("New Config File is added and will be used.")
		fmt.Println("Path is ", cfgFile)
		cfgFile = last
	} else {
		fmt.Println("No config provided and no previous config found.")
		os.Exit(1)
	}
}
