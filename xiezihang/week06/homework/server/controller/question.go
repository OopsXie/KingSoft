package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"server/database"
	"server/model"

	"github.com/gin-gonic/gin"
)

type AIRequest struct {
	Type       string `json:"type"`
	Count      int    `json:"count"`
	Difficulty string `json:"difficulty"`
	Language   string `json:"language"`
}

type AIQuestion struct {
	Title   string   `json:"title"`
	Options []string `json:"options,omitempty"`
	Answer  string   `json:"answer"`
}

func GetQuestions(c *gin.Context) {
	// 获取参数
	keyword := c.Query("keyword")
	questionType := c.Query("type")

	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	// 查询数据
	questions, total, err := database.QueryQuestions(keyword, questionType, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "查询失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "查询成功",
		"data":  questions,
		"total": total,
	})
}

func AddQuestion(c *gin.Context) {
	var q model.Question
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "参数错误"})
		return
	}
	if err := database.InsertQuestion(q); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "添加失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "添加成功"})
}

func UpdateQuestion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "无效的 ID"})
		return
	}
	var q model.Question
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "参数错误"})
		return
	}
	q.ID = uint(id)
	if err := database.UpdateQuestion(q); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "更新失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "修改成功"})
}

func DeleteQuestions(c *gin.Context) {
	var ids struct {
		IDList []int `json:"ids"`
	}
	if err := c.ShouldBindJSON(&ids); err != nil || len(ids.IDList) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "参数错误"})
		return
	}
	if err := database.DeleteQuestions(ids.IDList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "删除成功"})
}

func GenerateByAI(c *gin.Context) {
	var req AIRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "参数错误"})
		return
	}

	var prompt string
	if req.Type == "编程题" {
		prompt = fmt.Sprintf("请为 %s 语言生成 %d 道%s的编程题，每题包括：title（题干，含代码），answer（参考答案）。返回 JSON 数组。",
			req.Language, req.Count, req.Difficulty)
	} else if req.Type == "单选题" {
		prompt = fmt.Sprintf("请为 %s 语言生成 %d 道%s的单选题，每题包括：title、options（A~D 四个选项，字符串数组）、answer（正确答案）。返回 JSON 数组。",
			req.Language, req.Count, req.Difficulty)
	} else if req.Type == "多选题" {
		prompt = fmt.Sprintf("请为 %s 语言生成 %d 道%s的多选题，每题包括：title、options（A~D 四个选项，字符串数组）、answer（正确答案）。返回 JSON 数组。",
			req.Language, req.Count, req.Difficulty)
	}

	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	url := "https://api.deepseek.com/v1/chat/completions"

	payload := map[string]interface{}{
		"model": "deepseek-chat",
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
		"temperature": 0.7,
	}

	jsonData, _ := json.Marshal(payload)
	reqBody := bytes.NewBuffer(jsonData)

	httpReq, _ := http.NewRequest("POST", url, reqBody)
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "请求 AI 出错"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "AI响应解析失败"})
		return
	}

	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "AI返回内容为空或格式不符"})
		return
	}

	msg, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "AI消息内容缺失"})
		return
	}

	content := msg["content"].(string)
	content = strings.TrimSpace(content)
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimPrefix(content, "```")
	content = strings.TrimSuffix(content, "```")
	content = strings.TrimSpace(content)

	var generated []AIQuestion
	if err := json.Unmarshal([]byte(content), &generated); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "AI 生成成功，但解析失败，请手动检查",
			"raw":  content,
		})
		return
	}

	var inserted []model.Question
	for _, q := range generated {
		question := model.Question{
			Title:      q.Title,
			Type:       req.Type,
			Options:    strings.Join(q.Options, "|"),
			Answer:     q.Answer,
			Difficulty: req.Difficulty,
		}
		if req.Type == "编程题" {
			question.Options = ""
		} else if len(q.Options) != 4 {
			continue
		}

		// 处理多选题答案格式
		if req.Type == "多选题" {
			// 确保答案是一个字符串，包含多个选项字母
			if len(q.Answer) == 1 {
				// 如果只有一个答案，随机添加另一个答案
				options := []string{"A", "B", "C", "D"}
				otherOptions := make([]string, 0)
				for _, opt := range options {
					if opt != q.Answer {
						otherOptions = append(otherOptions, opt)
					}
				}
				randomIndex := time.Now().UnixNano() % int64(len(otherOptions))
				question.Answer = q.Answer + otherOptions[randomIndex]
			}
		}

		if id, err := database.InsertQuestionReturnID(question); err == nil {
			now := time.Now().Format(time.RFC3339)
			question.ID = uint(id)
			question.CreatedAt = now
			question.UpdatedAt = now
			inserted = append(inserted, question)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "AI 生成成功并已保存题目",
		"count": len(inserted),
		"data":  inserted,
	})
}

func DeleteOneQuestion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "无效的 ID"})
		return
	}
	if err := database.DeleteQuestions([]int{id}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "删除成功"})
}
