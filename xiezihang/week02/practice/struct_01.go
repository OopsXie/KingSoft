package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func NewPerson(name string, age int, email string) Person {
	return Person{
		Name:  name,
		Age:   age,
		Email: email,
	}
}

func PrintPerson(p Person) {
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("JSON:", string(data))

	//fmt.Printf("Person    姓名: %s, 年龄: %d, 邮箱: %s\n", p.Name, p.Age, p.Email)
}

func main() {
	var person_1 Person

	var Name string
	var Age int
	var Email string
	fmt.Println("请输入姓名：")
	fmt.Scanln(&Name)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&Age)
	fmt.Println("请输入邮箱：")
	fmt.Scanln(&Email)
	person_1 = NewPerson(Name, Age, Email)

	PrintPerson(person_1)
	//fmt.Println("Person: ", person_1.Name, person_1.Age, person_1.Email)
}
