/*
Copyright © 2025 Harpreet Singh harry798677@gmail.com
*/
package taskfunc

import (
	"encoding/json"
	"fmt"
	"todolist/utils"
)

func Delete_task(cfgFile string, id int) bool {
	var temp []TaskStruct
	encodedString := utils.Getfilecontent(cfgFile)
	if encodedString != "[]" {
		err := json.Unmarshal([]byte(encodedString), &taskList)
		if err != nil {
			fmt.Println("Something went wrong in Reading File")
			return false
		}
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
	}
	temp = append(taskList[:id-1], taskList[id:]...)
	data, err := json.Marshal(temp)
	if err != nil {
		fmt.Println("Something went wrong in Converting JSON Object")
		return false
	}
	utils.Savefilecontent(string(data), cfgFile)
	return true
}
