## 整体思路
- 实现对`go run main.go 1 1000000`命令行的读取
- 使用for循环实现对给定范围内素数的查找
- 将这些素数写入txt文件中，并记录素数个数
- 实现运行时间的记录
- 构思并发操作
### 构思并发操作
- 将当前的顺序执行改为 使用 4 个 goroutine 并发处理素数计算任务
- 把 [start, end] 区间平均分成 4 段，每个 goroutine 负责一段：
```
段长 = (end - start + 1) / 4
段1: [start1, end1]
段2: [start2, end2]
段3: [start3, end3]
段4: [start4, end]
```
- 每段用一个 goroutine 查找素数
- 每个 goroutine 调用 findPrimesInRange(start, end)
- 主线程合并所有 goroutine 的结果
- 写入文件 + 打印信息