package handler

import (
	"net/http"
	"strings"
	"time"

	"homework/model"
	"homework/service"
	"homework/utils"

	"github.com/gin-gonic/gin"
)

func CreateQuestion(c *gin.Context) {
	var req model.QuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		record := model.APIResult{
			HTTPCode: 1,
			HTTPMsg:  "参数错误",
			AIReq:    req,
		}
		_ = utils.AppendToJSONFile(record)
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "参数错误", "aiRes": nil})
		return
	}

	if req.Model == "" {
		req.Model = "tongyi"
	}
	if req.Language == "" {
		req.Language = "go"
	}
	if req.Type == 0 {
		req.Type = 1
	}

	langs := map[string]bool{"go": true, "javascript": true, "java": true, "python": true, "c++": true}
	if _, ok := langs[strings.ToLower(req.Language)]; !ok {
		record := model.APIResult{AIReq: req, HTTPCode: 1, HTTPMsg: "语言必须为 go/javascript/java/python/c++"}
		_ = utils.AppendToJSONFile(record)
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "语言必须为 go/javascript/java/python/c++", "aiRes": nil})
		return
	}
	if req.Type != 1 && req.Type != 2 {
		record := model.APIResult{AIReq: req, HTTPCode: 1, HTTPMsg: "type 只能为 1 或 2"}
		_ = utils.AppendToJSONFile(record)
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "type 只能为 1 或 2", "aiRes": nil})
		return
	}
	if req.Keyword == "" {
		record := model.APIResult{AIReq: req, HTTPCode: 1, HTTPMsg: "关键词 keyword 是必填项"}
		_ = utils.AppendToJSONFile(record)
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "关键词 keyword 是必填项", "aiRes": nil})
		return
	}

	start := time.Now()
	res, err := service.CallAI(req)
	end := time.Now()

	record := model.APIResult{
		AIStartTime: start.Format(time.RFC3339),
		AIEndTime:   end.Format(time.RFC3339),
		AICostTime:  end.Sub(start).String(),
		AIReq:       req,
	}

	if err != nil {
		record.HTTPCode = 1
		record.HTTPMsg = err.Error()
		_ = utils.AppendToJSONFile(record)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "aiRes": nil})
		return
	}

	record.HTTPCode = 0
	record.HTTPMsg = ""
	record.AIRes = res
	_ = utils.AppendToJSONFile(record)
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "aiRes": res})
}
