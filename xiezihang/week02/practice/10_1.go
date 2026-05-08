package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func OutPrint() {
	fmt.Println("===================================")
	fmt.Println("欢迎使用任务管理CLI应用")
	fmt.Println("1. 添加任务")
	fmt.Println("2. 显示所有未完成任务")
	fmt.Println("3. 标记任务为完成")
	fmt.Println("4. 删除任务")
	fmt.Println("5. 显示所有任务")
	fmt.Println("6. 退出")
	fmt.Println("输入你要进行的操作序号：")
}

type task struct {
	ID        int    `json:"id"`
	Details   string `json:"details"`
	Completed bool   `json:"completed"`
}

const taskFile = "tasks.json"

func getTasks() ([]task, error) {
	var task_1 []task
	file, err := ioutil.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return task_1, nil
		}
		return nil, err
	}
	err = json.Unmarshal(file, &task_1)
	if err != nil {
		return nil, err
	}
	return task_1, nil
}

func saveTasks(task_1 []task) error {
	data, err := json.Marshal(task_1)
	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	err = ioutil.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func addTask(task_1 []task, task_details string) []task {
	//fmt.Println("这是添加任务的功能")
	new_task := task{
		ID:        len(task_1) + 1,
		Details:   task_details,
		Completed: false,
	}
	task_1 = append(task_1, new_task)
	saveTasks(task_1)
	fmt.Println("添加任务成功")
	return task_1
}

func showUnfinishedTask(task_1 []task) {
	//fmt.Println("这是显示所有未完成任务的功能")
	isFinished := false
	for _, task := range task_1 {
		if !task.Completed {
			fmt.Printf("任务ID: %d, 内容: %s\n, 状态：%s\n", task.ID, task.Details, "未完成")
			isFinished = true
		}
	}
	if !isFinished {
		fmt.Println("没有未完成的任务")
	}
}

func markTaskAsCompleted(task_1 []task, id int) []task {
	//fmt.Println("这是标记任务为完成的功能")
	for i, task := range task_1 {
		if task.ID == id {
			task_1[i].Completed = true
			saveTasks(task_1)
			fmt.Println("任务已标记为完成:", task.Details)
			return task_1
		}
	}
	fmt.Println("找不到该任务ID")
	return task_1
}

func showAllTask(task_1 []task) {
	//fmt.Println("这是显示所有任务的功能")
	if len(task_1) == 0 {
		fmt.Println("task中没有内容")
		return
	}
	for _, task := range task_1 {
		status := "未完成"
		if task.Completed {
			status = "已完成"
		}
		fmt.Printf("ID: %d, 内容: %s, 状态: %s\n", task.ID, task.Details, status)
	}
}

func deleteTask(task_1 []task, id int) []task {
	//fmt.Println("这是删除任务的功能")
	for i, task := range task_1 {
		if task.ID == id {
			task_1 = append(task_1[:i], task_1[i+1:]...)
			saveTasks(task_1)
			fmt.Println("任务已删除")
			return task_1
		}
	}
	fmt.Println("找不到该任务ID")
	return task_1
}

func main() {
	var task_1 []task
	task_1, _ = getTasks()

	for {
		OutPrint()
		var n int
		fmt.Scan(&n)

		if n == 1 {
			//添加任务
			var task_details string
			fmt.Println("请输入任务详情：")
			fmt.Scan(&task_details)
			task_1 = addTask(task_1, task_details)
		} else if n == 2 {
			//显示所有未完成任务
			showUnfinishedTask(task_1)
		} else if n == 3 {
			//标记任务为完成
			var id int
			fmt.Println("请输入要标记完成的任务ID：")
			fmt.Scan(&id)
			task_1 = markTaskAsCompleted(task_1, id)
		} else if n == 4 {
			//删除任务
			fmt.Println("---------------任务列表---------------")
			showAllTask(task_1)
			fmt.Println("--------------------------------------")
			var id int
			fmt.Println("请输入要删除的任务ID：")
			fmt.Scan(&id)
			task_1 = deleteTask(task_1, id)
		} else if n == 5 {
			//显示所有任务
			showAllTask(task_1)
		} else if n == 6 {
			//退出
			break
		} else {
			fmt.Println("请输入正确的操作序号")
		}
	}

	fmt.Println("退出成功")
	fmt.Println("感谢使用任务管理CLI应用")
}
