package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	jsonStr := `{"name":"Jane Smith","age":25,"email":"janesmith@example.com"}`
	var person_1 Person

	err := json.Unmarshal([]byte(jsonStr), &person_1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//fmt.Printf("姓名: %s, 年龄: %d, 邮箱: %s\n", person_1.Name, person_1.Age, person_1.Email)

	fmt.Println("Name:", person_1.Name)
	fmt.Println("Age:", person_1.Age)
	fmt.Println("Email:", person_1.Email)
}
