package performance

import (
	"testing"
)

// 基准测试：通过值传递
func BenchmarkProcessByValue(b *testing.B) {
	data := CreateTestData()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_ = ProcessStructByValue(data)
	}
}

// 基准测试：通过指针传递
func BenchmarkProcessByPointer(b *testing.B) {
	data := CreateTestData()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		ProcessStructByPointer(&data)
	}
}

// 测试函数的正确性
func TestStructProcessing(t *testing.T) {
	// 测试值传递
	data1 := CreateTestData()
	result := ProcessStructByValue(data1)
	if result.Position != "Updated Position" || result.Email != "updated@example.com" {
		t.Error("ProcessStructByValue failed to update fields")
	}
	// 原始数据应该保持不变
	if data1.Position == "Updated Position" || data1.Email == "updated@example.com" {
		t.Error("ProcessStructByValue modified original data")
	}

	// 测试指针传递
	data2 := CreateTestData()
	ProcessStructByPointer(&data2)
	if data2.Position != "Updated Position" || data2.Email != "updated@example.com" {
		t.Error("ProcessStructByPointer failed to update fields")
	}
}
