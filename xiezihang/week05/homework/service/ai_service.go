package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"homework/config"
	"homework/model"
)

func CallAI(req model.QuestionRequest) (model.QuestionResponse, error) {
	prompt := fmt.Sprintf(`你是一名专业的编程出题官，请你针对关键词“%s”，为“%s”编程语言生成一道%s选择题，答案有四个选项，标出正确答案，结果以 JSON 格式输出，格式如下：\n\n{
  "title": "题目内容",
  "answers": ["选项A", "选项B", "选项C", "选项D"],
  "right": [0] // 正确选项索引，从0开始
}`,
		req.Keyword,
		req.Language,
		map[int]string{1: "单选", 2: "多选"}[req.Type],
	)

	switch strings.ToLower(req.Model) {
	case "deepseek":
		return callDeepseek(prompt)
	case "tongyi":
		return callTongyi(prompt)
	default:
		return model.QuestionResponse{}, errors.New("不支持的模型")
	}
}

func callDeepseek(prompt string) (model.QuestionResponse, error) {
	payload := map[string]interface{}{
		"model": "deepseek-chat",
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}
	data, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "https://api.deepseek.com/v1/chat/completions", bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+config.DeepseekKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.QuestionResponse{}, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var parsed map[string]interface{}
	_ = json.Unmarshal(body, &parsed)

	choices := parsed["choices"].([]interface{})
	msg := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	content := msg["content"].(string)

	return parseAIContent(content)
}

func callTongyi(prompt string) (model.QuestionResponse, error) {
	payload := map[string]interface{}{
		"model": "qwen-plus",
		"input": map[string]string{"prompt": prompt},
	}
	data, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation", bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+config.TongyiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.QuestionResponse{}, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var parsed map[string]interface{}
	_ = json.Unmarshal(body, &parsed)
	output := parsed["output"].(map[string]interface{})
	content := output["text"].(string)

	return parseAIContent(content)
}

func parseAIContent(content string) (model.QuestionResponse, error) {
	content = strings.TrimSpace(content)

	if strings.HasPrefix(content, "```") {
		content = strings.TrimPrefix(content, "```json")
		content = strings.TrimPrefix(content, "```")
	}
	content = strings.TrimSuffix(content, "```")
	content = strings.TrimSpace(content)

	var result model.QuestionResponse
	err := json.Unmarshal([]byte(content), &result)
	if err != nil {
		return model.QuestionResponse{}, errors.New("AI返回格式非JSON: " + err.Error())
	}
	if len(result.Answers) != 4 {
		return model.QuestionResponse{}, errors.New("AI返回答案选项数量必须为4个")
	}
	return result, nil
}
