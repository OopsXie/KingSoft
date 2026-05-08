package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

type Log struct {
	Timestamp time.Time
	UserID    string
	Action    string
	Details   string
}

func main() {

	file, err := os.Open("user_behavior.log")
	if err != nil {
		panic(fmt.Sprintf("无法打开日志文件: %v", err))
	}
	defer file.Close()

	layout := "2006-01-02 15:04:05" // 修正时间格式
	scanner := bufio.NewScanner(file)

	userStats := make(map[string][]time.Time)
	actionStats := make(map[string]int)
	minuteStats := make(map[string]map[string]bool)
	minuteOps := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 4 {
			continue
		}

		t, err := time.Parse(layout, parts[0])
		if err != nil {
			fmt.Printf("时间解析失败: %v\n", err)
			continue
		}
		minuteKey := t.Format("2006-01-02 15:04")

		userID := parts[1]
		action := parts[2]

		userStats[userID] = append(userStats[userID], t)
		actionStats[action]++
		if _, ok := minuteStats[minuteKey]; !ok {
			minuteStats[minuteKey] = make(map[string]bool)
		}
		minuteStats[minuteKey][userID] = true
		minuteOps[minuteKey]++
	}

	writeUserStatistics(userStats)
	writeActionStatistics(actionStats)
	writeMinuteStatistics(minuteStats, minuteOps)
}

func writeUserStatistics(userStats map[string][]time.Time) {
	file, err := os.Create("user_statistics.csv")
	if err != nil {
		panic(fmt.Sprintf("无法创建文件 user_statistics.csv: %v", err))
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"用户ID", "操作次数", "首次操作时间", "最后操作时间"}); err != nil {
		panic(fmt.Sprintf("写入表头失败: %v", err))
	}

	for uid, times := range userStats {
		if len(times) == 0 {
			continue
		}
		earliest, latest := times[0], times[0]
		for _, t := range times {
			if t.Before(earliest) {
				earliest = t
			}
			if t.After(latest) {
				latest = t
			}
		}
		if err := writer.Write([]string{
			uid,
			fmt.Sprintf("%d", len(times)),
			earliest.Format("2006-01-02 15:04:05"),
			latest.Format("2006-01-02 15:04:05"),
		}); err != nil {
			panic(fmt.Sprintf("写入用户统计数据失败: %v", err))
		}
	}
}

func writeActionStatistics(actionStats map[string]int) {
	file, err := os.Create("action_statistics.csv")
	if err != nil {
		panic(fmt.Sprintf("无法创建文件 action_statistics.csv: %v", err))
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"行为类型", "总次数"}); err != nil {
		panic(fmt.Sprintf("写入表头失败: %v", err))
	}

	for action, count := range actionStats {
		if err := writer.Write([]string{action, fmt.Sprintf("%d", count)}); err != nil {
			panic(fmt.Sprintf("写入行为统计数据失败: %v", err))
		}
	}
}

func writeMinuteStatistics(minuteStats map[string]map[string]bool, minuteOps map[string]int) {
	file, err := os.Create("minute_statistics.csv")
	if err != nil {
		panic(fmt.Sprintf("无法创建文件 minute_statistics.csv: %v", err))
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"时间段", "活跃用户数", "操作总数"}); err != nil {
		panic(fmt.Sprintf("写入表头失败: %v", err))
	}

	for minute, users := range minuteStats {
		if err := writer.Write([]string{
			minute,
			fmt.Sprintf("%d", len(users)),
			fmt.Sprintf("%d", minuteOps[minute]),
		}); err != nil {
			panic(fmt.Sprintf("写入分钟统计数据失败: %v", err))
		}
	}
}
