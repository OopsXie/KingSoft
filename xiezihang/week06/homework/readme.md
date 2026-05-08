# 考试出题中后台系统

## 项目简介
本项目是一个基于 React + Ant Design 前端、Go + Gin 后端的考试出题中后台系统。系统支持手动出题、AI自动生成试题、题库管理、学习心得展示等功能。

## 主要功能
**AI智能出题**：
- 选择题型（单选题、多选题、编程题）
- 选择语言（Go、Java、Python、JavaScript）
- 选择难度（简单、中等、困难）
- 选择题目数量（最多 10 道）
- 通过调用 [DeepSeek](https://deepseek.com) API 自动生成并保存到数据库

**手动出题**：
- 可自定义题型、语言、难度
- 自动展示选项输入框及答案选择
- 复用该表单实现题目的编辑功能
  
**题库管理**：
- 查询题目：支持关键字和题型筛选
- 编辑题目：点击编辑弹出与手动出题相同的表单
- 删除题目：支持单个删除和批量删除
- 分页显示：每页支持切换条数，显示总数
- 
**学习心得**：
- 首页读取 `homework/README.md`，渲染为 Markdown 内容
- 使用 `react-markdown` 实现解析与展示

## 搭建过程及关键问题

### 前端开发
1. **初始化项目**
   - 使用 `npm create vite@latest` 创建 React + TypeScript 项目。
   - 配置 `vite.config.ts`，将 `/api` 请求代理到后端 `http://localhost:8080`，以解决跨域问题：
     ```ts
     export default defineConfig({
        server: {
            port: 3000,
            proxy: {
            '/api': {
                target: 'http://localhost:8080',
                changeOrigin: true,
                rewrite: path => path.replace(/^\/api/, '/api')
            }
            }
        },
        plugins: [react()]
        })
     ```

2. **集成 Ant Design**
   - 安装依赖：`npm install antd`
   - 使用 Ant Design 提供的组件构建页面结构，包括：
     - `Table` 实现题库展示
     - `Modal + Form` 实现新增和编辑表单
     - `Select`、`Radio`、`Input` 等用于填写题目信息

3. **实现题库管理模块**
   - 构建题型筛选按钮组、搜索输入框
   - 封装分页请求，通过 `axios.get('/api/questions', { params: {...} })` 获取后端数据
   - 支持单题删除、批量删除、编辑功能
   - 编辑功能通过 `setFieldsValue()` 自动填充已有值到表单

4. **实现 AI 出题界面**
   - 使用多个下拉框（Select）输入题型、语言、难度
   - 数字输入框（InputNumber）控制生成题目数量
   - 点击按钮发起 POST 请求 `/api/questions/ai_generate`，生成题目
   - 异常情况如请求失败、格式错误等使用 `message.error()` 反馈

5. **展示 Markdown 学习心得**
   - 在首页通过调用 `/api/readme` 接口获取 `README.md` 内容
   - 使用 `react-markdown` 将 Markdown 文本渲染为页面
   - 支持标题、列表、代码块等基本格式展示

### 后端开发
1. **Gin 项目初始化**
   - 创建 `main.go` 并初始化 `gin.Default()` 实例
   - 在 `router.go` 中配置 API 路由组 `/api`
   - 路由包括：`GET /questions`、`POST /questions`、`PUT`、`DELETE`、`POST /questions/delete` 等

2. **数据库集成**
   - 使用 `modernc.org/sqlite` 嵌入式 SQLite 实现数据库
   - 在 `database/db.go` 中封装数据库初始化逻辑（建表 + 连接）
   - 定义字段包括：题目、题型、难度、选项、答案、时间戳等

3. **AI 生成题目**
   - 在 `controller/question.go` 中通过 `os.Getenv("DEEPSEEK_API_KEY")` 读取 API Key
   - 构造 prompt 请求 DeepSeek 接口，解析返回的 JSON 并插入数据库
   - 针对部分 AI 返回非法 JSON 的情况，加入前缀清理与解析失败的兜底处理

4. **学习心得接口**
   - 新增 `/api/readme` 接口，使用 `os.ReadFile("../../README.md")` 读取 Markdown 内容并返回
   - 注意路径必须基于服务器 `main.go` 的运行位置来确定

5. **静态资源托管**
   - 生产环境打包前端后，将 `client/dist` 设置为静态目录

### 遇到的问题
1. **路径问题**
   - 后端在读取 Markdown 文件或加载前端页面时可能会报错。此问题通常由相对路径错误引起。解决方法包括使用 os.Getwd() 获取当前工作目录，或手动调整路径为 ../client/dist/index.html 等实际文件位置。
  
2. **分页逻辑丢失**
   - 当前端切换页码时，如果未将当前页码同步传递到后端，则会导致数据加载错误。解决方案是在前端分页组件的 pagination.current 和 onChange 中手动传参，调用 fetchData() 函数重新获取对应页的数据。

3. **AI JSON 错误**
   - 有时调用 AI 接口返回的内容不是标准的 JSON 格式，导致解析失败。建议在解析前先使用 TrimPrefix("json") 去除多余内容，并在 json.Unmarshal() 过程中加入异常捕获，以避免程序崩溃。

4. **编辑功能无法复用表单**
   - 在复用“新增”表单组件进行“编辑”时，无法正确填充已有数据。推荐方法是引入 editingQuestion 状态作为判断标志，并在组件的 useEffect 中使用 form.setFieldsValue() 将选中题目的内容自动填充到表单中。

5. **前后端端口不一致**
   - 在部署生产环境时，如果前后端端口不一致，会导致页面无法正常访问。解决方案是使用 vite build 命令打包前端为静态资源，由 Go 后端（通常监听 8080 端口）统一托管这些 HTML、CSS 和 JS 文件。

6. **学习心得不展示**
   - 当前端请求 /api/readme 接口时无响应，通常是由于后端读取路径错误或文件不存在。应确认路径设置是否为 ../../README.md，同时确保该文件真实存在于服务器对应目录中。

**目录结构**

```
        homework/
        ├── client/                       # 前端项目目录 
        │   ├── dist/                     # 前端打包后生成的静态资源目录
        │   ├── public/
        │   ├── src/                      # 前端源码目录
        │   │   ├── assets/
        │   │   ├── pages/
        │   │   ├── App.tsx
        │   │   └── ...
        │   ├── index.html
        │   ├── vite.config.ts           # Vite 配置文件，含 API 代理配置
        │   └── package.json
        │
        ├── server/                      # 后端项目目录
        │   ├── controller/              # 控制器层，处理业务逻辑
        │   │   ├── question.go          # 题库管理相关接口
        │   │   ├── readme.go            # 读取 README.md 接口
        │   │   └── ...
        │   ├── database/               
        │   │   └── db.go                # SQLite 数据库操作逻辑
        │   ├── model/                   
        │   │   └── question.go          
        │   ├── router/                  # 路由注册
        │   │   └── router.go            # 路由配置
        │   └── main.go                  # Go 项目主入口
        │
        ├── questions.db                 
        ├── README.md                    
        └── .gitignore                  
```

## 技术栈
- 前端：React、TypeScript、Ant Design、react-markdown
- 后端：Go、Gin
- 数据库：SQLite

## 快速开始

### 1. 启动

- cd week06/homework/server
- go mod tidy
- go run main.go
- 浏览器打开 http://localhost:8080/



### 2. 访问系统
- 浏览器访问 [http://localhost:8080](http://localhost:8080)
- 首页即为"学习心得"页面，题库管理和AI出题功能可通过侧边栏进入。

## 配置说明
- AI出题功能需配置 `DEEPSEEK_API_KEY` 环境变量，详见 `server/controller/question.go`。
- 数据库存储在 `questions.db` 文件中，自动创建。