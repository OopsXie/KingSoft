# 调用ai出题（第三次考核）

## 整体思路

先通 Gin 架子 → 定义结构体 → 接请求 → 接入模型 API → 写入 JSON →  解析结果 → 统一返回 → 处理异常 → 多轮测试

### 第一步：搭建基本项目结构

准备好目录框架，能够运行 Gin 服务

创建并配置 `go.mod`

使用 Gin 启动 Web 服务（`main.go`）



### 第二步：设计请求结构体 & 响应结构体

明确输入输出数据格式，便于后续参数校验和响应生成

在 `model/question.go` 中定义：

```
type QuestionRequest struct {
  Model    string `json:"model"`
  Language string `json:"language"`
  Type     int    `json:"type"`
  Keyword  string `json:"keyword"`
}

type QuestionResponse struct {
  Title   string   `json:"title"`
  Answers []string `json:"answers"`
  Right   []int    `json:"right"`
}

```



### 第三步：实现基本的 Handler 层逻辑（接收请求）

接收请求体 JSON 并完成基础参数校验



### 第四步：接入真实模型 API

添加 `.env` 配置（API Key）

接收模型返回文本，提取其中的 JSON



### 第五步：写入 JSON 文件

每次调用结果保存到 `data/YYYY_MM_DD.json`

- 每天一个文件，记录为数组格式
- 如果文件存在则读取 → 追加 → 写回
- 使用 `os.MkdirAll` 确保目录存在



### 项目模块划分

| 模块       | 说明                                           |
| ---------- | ---------------------------------------------- |
| `main.go`  | 程序入口，启动 Gin 服务，加载配置              |
| `handler/` | 路由注册与 API 接口实现，接收请求并调用业务层  |
| `model/`   | 请求体与响应体结构体定义                       |
| `service/` | 调用 Deepseek/Tongyi API，构造提示词并解析响应 |
| `config/`  | 读取 `.env` 文件中的 API Key 信息              |
| `utils/`   | 文件追加写入工具，保存生成的题库数据           |
| `data/`    | 每天一个 JSON 文件，记录当天所有题目请求       |



## 代码运行

```go
cd week05/homework
go mod tidy
go run main.go
```

### 用 Postman 或 curl 测试接口

**接口地址：**`POST http://localhost:8081/api/questions/create`

**测试样例：**

```json
{
  "model": "deepseek",
  "language": "go",
  "type": 1,
  "keyword": "Gin 框架中间件"
}
```

**HTTP响应样式：**

```
{
    "aiRes": {
        "title": "在Gin框架中，中间件的主要作用是什么？",
        "answers": [
            "处理HTTP请求和响应，可以在请求到达处理器之前或之后执行特定逻辑",
            "仅用于数据库连接和操作",
            "用于前端页面的渲染和展示",
            "用于定义路由规则和URL匹配"
        ],
        "right": [
            0
        ]
    },
    "code": 0,
    "msg": ""
}
```

**json文件存储样式**

```json
[
  {
    "aiStartTime": "2025-04-23T19:55:04+08:00",
    "aiEndTime": "2025-04-23T19:55:14+08:00",
    "aiCostTime": "10.3954623s",
    "aiReq": {
      "model": "deepseek",
      "language": "go",
      "type": 1,
      "keyword": "Gin 框架中间件"
    },
    "aiRes": {
      "title": "在Gin框架中，中间件的主要作用是什么？",
      "answers": [
        "处理HTTP请求和响应，可以在请求到达处理器之前或之后执行特定逻辑",
        "仅用于数据库连接和操作",
        "用于前端页面的渲染和展示",
        "用于定义路由规则和URL匹配"
      ],
      "right": [
        0
      ]
    },
    "httpCode": 0,
    "httpMsg": ""
  }
]
```



### **接口说明**

| 字段       | 含义       | 说明                                     |
| ---------- | ---------- | ---------------------------------------- |
| `model`    | 使用模型   | deepseek / tongyi（默认 tongyi）         |
| `language` | 编程语言   | go / js / python / java / c++（默认 go） |
| `type`     | 题型       | 1 单选题 / 2 多选题（默认 1）            |
| `keyword`  | 关键词描述 | 必填，用于生成题目内容                   |

代码实现：

默认tongyi，默认go，默认单选，keyword必填

```go
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
```

