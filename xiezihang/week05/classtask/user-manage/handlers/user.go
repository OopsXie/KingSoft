package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age"`
}

var filePath = "user.json"

// 获取所有用户
func GetUsers(c *gin.Context) {
	users := readUsers()
	c.JSON(http.StatusOK, users)
}

// 创建用户（确保 email 唯一）
func CreateUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users := readUsers()
	// 检查 email 是否已存在
	for _, u := range users {
		if u.Email == newUser.Email {
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
			return
		}
	}
	users = append(users, newUser)
	writeUsers(users)
	c.JSON(http.StatusCreated, newUser)
}

// 修改用户（根据 email 查找）
func UpdateUser(c *gin.Context) {
	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users := readUsers()
	updated := false
	for i, u := range users {
		if u.Email == updatedUser.Email {
			users[i] = updatedUser
			updated = true
			break
		}
	}
	if !updated {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	writeUsers(users)
	c.JSON(http.StatusOK, updatedUser)
}

// 删除用户（通过 email）
func DeleteUser(c *gin.Context) {
	email := c.Param("email")
	users := readUsers()
	found := false
	newUsers := make([]User, 0)
	for _, u := range users {
		if u.Email == email {
			found = true
			continue
		}
		newUsers = append(newUsers, u)
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	writeUsers(newUsers)
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

// 工具函数：读取用户列表
func readUsers() []User {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil || len(bytes) == 0 {
		return []User{}
	}
	var users []User
	_ = json.Unmarshal(bytes, &users)
	return users
}

// 工具函数：写入用户列表
func writeUsers(users []User) {
	bytes, _ := json.MarshalIndent(users, "", "  ")
	_ = ioutil.WriteFile(filePath, bytes, 0644)
}
