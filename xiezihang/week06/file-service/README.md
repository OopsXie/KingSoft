# 文件管理服务实现
## 整体思路
### 首先构建基础Gin框架
```
    my-gin-app/
    ├── main.go
    ├── routes/
    │   └── router.go
    ├── controllers/
    │   └── hello.go
    └── go.mod
```

- 参考了week05中classwork中的课堂练习
- 发觉这个考试项目和week05/classwork/img-upload很相似
- 对比课程中 week05/classwork/img-upload 项目，发现本项目是其功能拓展版本，主要增加了：
    1. 分页查询
    2. 类型过滤
    3. 文件预览
    4. 下载原始文件
    5. 类型统计
    6. SQLite 存储元信息

### 增加模块与功能结构
- 目录结构
```
    file-service/
    ├── db/db.go                   // 数据库初始化
    ├── models/file.go               // 数据模型
    ├── utils/response.go                // 通用响应封装
    ├── storage/              // 文件存储目录（按类型分类）
```
- 在文件存储中，我又实现了按类型存储
```
    storage/
    ├── jpg/
    ├── png/
    ├── html/
    ├── css/
    ├── js/

```
- 数据库初始化：db/db.go

- 在routes/router.go中实现路由，接口
```
	r.POST("/api/uploads", controllers.Upload)
	r.GET("/api/list", controllers.ListFiles)
	r.GET("/api/preview/:filename", controllers.PreviewFile)
	r.GET("/api/download/:filename", controllers.DownloadFile)
	r.GET("/api/stats", controllers.Stats)
	r.DELETE("/api/deleteimg", controllers.DeleteFile)
```

- utlis/response.go中实现错误与成功的反馈
- 所有接口使用该格式进行 JSON 响应，统一返回字段：

- controllers/ 功能实现
```
    upload.go：文件上传，自动按类型保存至 storage/{type} 子目录
    list.go：分页列出文件
    download.go：在线预览文件，以及下载并返回原始文件名
    stats.go：统计每类文件数量和大小
    deleteimg.go：删除文件并清除数据库记录
```


### 代码运行
```
    cd week06/file-service
    go mod tidy
    go run main.go
```
- 在POSTMAN中进行测试

### 测试案例
#### POSTMAN
```
    Method: POST
    URL: http://localhost:8080/api/uploads
    Body: form-data
    Key: files（类型：File）
```

- 成功返回示例：
```
{
    "code": 0,
    "data": {
        "uploaded": [
            "73b45322-3629-422f-9e0d-46ffc2208536.jpg"
        ]
    },
    "msg": "success"
}
```
