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

func Update_task(cfgFile string, id int) bool {
	encodedString := utils.Getfilecontent(cfgFile)
	if encodedString != "[]" {
		err := json.Unmarshal([]byte(encodedString), &taskList)
		if err != nil {
			fmt.Println("Something went wrong in Reading File")
			return false
		}
	}
	if id <= 0 {
		fmt.Println("Invalid index")
		return false
	}
	taskIndex := -1
	for index, tasks := range taskList {
		if id == tasks.Task_id {
			taskIndex = index
			break
		}
	}
	if taskIndex == -1 {
		fmt.Println("ID Not Found")
		return false

	} else {
		taskList[taskIndex-1].Status = "Completed"
		taskList[taskIndex-1].CompleteDate = time.Now().Unix()
		fmt.Println("Updated the ", taskList[taskIndex-1], " Task")
	}
	data, err := json.Marshal(taskList)
	if err != nil {
		fmt.Println("Something went wrong in Converting JSON Object")
		return false
	}
	utils.Savefilecontent(string(data), cfgFile)
	fmt.Println(taskIndex, id)
	return true
}
