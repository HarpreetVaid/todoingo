/*
Copyright © 2025 Harpreet Singh harry798677@gmail.com
*/
package taskfunc

import (
	"encoding/json"
	"fmt"
	"time"
	"todolist/utils"
)

type TaskStruct struct {
	Task_id      int    `json:"id"`
	Task         string `json:"task"`
	Status       string `json:"status"`
	CreationTime int64  `json:"created_at"`
	CompleteDate int64  `json:"completed_at"`
}

var taskList []TaskStruct

func Add_task(cfgFile string, task string) bool {
	encodedString := utils.Getfilecontent(cfgFile)
	var id int = 0
	if encodedString != "[]" {
		err := json.Unmarshal([]byte(encodedString), &taskList)
		if err != nil {
			fmt.Println("Something went wrong in Reading File")
			return false
		}
	}
	if len(taskList) > 0 {
		id = taskList[len(taskList)-1].Task_id + 1
	}
	newTask := TaskStruct{
		Task_id:      id,
		Task:         task,
		Status:       "InProgress",
		CreationTime: time.Now().Unix(),
		CompleteDate: -1,
	}
	taskList = append(taskList, newTask)
	data, err := json.Marshal(taskList)
	if err != nil {
		fmt.Println("Something went wrong in Converting JSON Object")
		return false
	}
	utils.Savefilecontent(string(data), cfgFile)
	return true
}
