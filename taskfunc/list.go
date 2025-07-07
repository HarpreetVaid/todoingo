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

func List_task(cfgFile string) {
	encodedString := utils.Getfilecontent(cfgFile)
	if encodedString != "[]" {
		err := json.Unmarshal([]byte(encodedString), &taskList)
		if err != nil {
			fmt.Println("Something went wrong in Reading File")
		}
		for index, tasks := range taskList {
			if int64(tasks.CompleteDate) == -1 {
				fmt.Printf(" %v) The %v Task id %v is and Started at %v and in Progress\n", index+1, tasks.Task, tasks.Task_id+1, time.Unix(int64(tasks.CreationTime), 0).Format(time.RFC822))
			} else {
				fmt.Printf(" %v) The %v Task id %v is and Started at %v and completed at %v\n", index+1, tasks.Task, tasks.Task_id+1, time.Unix(int64(tasks.CreationTime), 0).Format(time.RFC822), time.Unix(int64(tasks.CompleteDate), 0).Format(time.RFC822))
			}
		}
	} else {
		fmt.Println("Sorry Your Tasklist is empty")
	}

}
