package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func OutPrint() {
	fmt.Println("===================================")
	fmt.Println("欢迎使用学生成绩管理系统")
	fmt.Println("1.录入学生成绩")
	fmt.Println("2.查询学生成绩")
	fmt.Println("3.修改学生成绩")
	fmt.Println("4.删除学生成绩")
	fmt.Println("5.退出系统")
	fmt.Println("请选择你要进行的操作：")
}

type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

const studentFile = "students.json"

func getStudents() ([]Student, error) {
	var student_1 []Student
	file, err := ioutil.ReadFile("students.json")
	if err != nil {
		if os.IsNotExist(err) {
			return student_1, nil
		}
		return nil, err
	}
	err = json.Unmarshal(file, &student_1)
	if err != nil {
		return nil, err
	}
	return student_1, nil
}

func saveStudents(student_1 []Student) error {
	data, err := json.Marshal(student_1)
	if err != nil {
		fmt.Println("Error ", err)
		return err
	}

	err = ioutil.WriteFile("students.json", data, 0644)
	if err != nil {
		return err
	}
	return nil

}

func addStudent(student_1 []Student) []Student {
	var s_id string
	var s_name string
	var s_score float64
	fmt.Println("请输入学生的ID：")
	fmt.Scanln(&s_id)
	fmt.Println("请输入学生的姓名：")
	fmt.Scanln(&s_name)
	fmt.Println("请输入学生的成绩：")
	fmt.Scanln(&s_score)

	student_2 := Student{
		ID:    s_id,
		Name:  s_name,
		Score: s_score,
	}
	// student_2.id = id
	// student_2.name = name
	// student_2.score = score

	student_1 = append(student_1, student_2)
	saveStudents(student_1)
	fmt.Println("添加学生成绩成功")
	return student_1
}

func findStudent(student_1 []Student) {
	var id string
	fmt.Println("请输入要查询的学生ID：")
	fmt.Scanln(&id)

	for _, student := range student_1 {
		if student.ID == id {
			fmt.Printf("ID: %s, 姓名: %s, 成绩: %.2f\n", student.ID, student.Name, student.Score)
			return
		}
	}
	fmt.Println("该id不存在")
}

func changeStudent(student_1 []Student) []Student {
	var id string
	fmt.Println("请输入要修改的学生ID：")
	fmt.Scanln(&id)

	for i, student := range student_1 {
		if student.ID == id {
			var name string
			var score float64
			fmt.Println("请输入新的姓名：")
			fmt.Scanln(&name)
			fmt.Println("请输入新的成绩：")
			fmt.Scanln(&score)

			student_1[i].Name = name
			student_1[i].Score = score

			saveStudents(student_1)
			fmt.Println("修改学生成绩成功")
			return student_1
		}
	}
	fmt.Println("该id不存在")
	return student_1
}

func showStudent(student_1 []Student) {
	fmt.Println("----------学生成绩列表----------")
	if len(student_1) == 0 {
		fmt.Println("没有学生成绩")
		return
	}
	for _, student := range student_1 {
		fmt.Printf("ID: %s, 姓名: %s, 成绩: %.2f\n", student.ID, student.Name, student.Score)
	}
	fmt.Println("-------------------------------")
}

func deleteStudent(student_1 []Student) []Student {
	showStudent(student_1)
	var id string
	fmt.Println("请输入要删除的学生ID：")
	fmt.Scanln(&id)

	for i, student := range student_1 {
		if student.ID == id {
			student_1 = append(student_1[:i], student_1[i+1:]...)
			saveStudents(student_1)
			fmt.Println("删除学生成绩成功")
			return student_1
		}
	}
	fmt.Println("该id不存在")
	return student_1
}

func main() {

	var student_1 []Student
	student_1, _ = getStudents()

	for {
		OutPrint()
		var n int
		fmt.Scanln(&n)

		if n == 1 {
			// 录入学生成绩
			student_1 = addStudent(student_1)
		} else if n == 2 {
			// 查询学生成绩
			findStudent(student_1)
		} else if n == 3 {
			// 修改学生成绩
			student_1 = changeStudent(student_1)
		} else if n == 4 {
			// 删除学生成绩
			student_1 = deleteStudent(student_1)
		} else if n == 5 {
			// 退出系统
			break
		} else {
			fmt.Println("请输入正确的操作序号")
		}

	}

	fmt.Println("退出成功")
	fmt.Println("感谢使用学生成绩管理系统")
}
