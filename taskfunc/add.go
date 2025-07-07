/*
Copyright © 2025 NAME HERE harry798677@gmail.com
*/
package taskfunc

import (
	"fmt"
	"time"
	"todolist/utils"
)

type TaskStruct struct {
	task         string
	status       string
	creationTime int64
	completeDate int64
}

var taskList []TaskStruct

func Add_task(cfgFile string, task string) bool {
	newTask := TaskStruct{
		task:         task,
		status:       "InProgress",
		creationTime: time.Now().Unix(),
		completeDate: 0,
	}
	taskList = append(taskList, newTask)
	fileContent := utils.Getfilecontent(cfgFile)
	fmt.Println("File Data is", fileContent)
	return true
}
