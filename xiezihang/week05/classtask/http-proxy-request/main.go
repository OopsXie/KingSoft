package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/resource/list", func(c *gin.Context) {
		url := "https://api.juejin.cn/content_api/v1/resource_placements/shows"
		// 构造 POST 请求，注意参数结构
		reqBody := []byte(`{
			"placement_ids": ["home_banner"]
		}`)

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "代理请求失败"})
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "读取响应失败"})
			return
		}

		fmt.Println("返回内容是：", string(body)) // 调试打印

		var raw map[string]interface{}
		if err := json.Unmarshal(body, &raw); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "JSON解析失败"})
			return
		}

		result := gin.H{
			"code":    raw["err_no"],
			"message": raw["err_msg"],
			"data":    raw["data"],
		}

		c.JSON(http.StatusOK, result)
	})

	r.Run(":8080")
}
