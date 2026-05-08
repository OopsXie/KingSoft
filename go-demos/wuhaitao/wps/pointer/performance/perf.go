package performance

// BigStruct 模拟一个包含大量数据的用户信息结构体
type BigStruct struct {
	ID        int64
	Name      string
	Email     string
	Address   string
	Phone     string
	Company   string
	Position  string
	Skills    [50]string    // 使用大数组来增加结构体大小
	Projects  [20]Project   // 嵌套结构体
	MetaData  [100]float64  // 更多数据来增加结构体大小
}

type Project struct {
	Name        string
	Description string
	Duration    int
	Team        [10]string
}

// 不使用指针的函数
func ProcessStructByValue(data BigStruct) BigStruct {
	data.Position = "Updated Position"
	data.Email = "updated@example.com"
	return data
}

// 使用指针的函数
func ProcessStructByPointer(data *BigStruct) {
	data.Position = "Updated Position"
	data.Email = "updated@example.com"
}

// 创建测试数据
func CreateTestData() BigStruct {
	var data BigStruct
	data.ID = 1
	data.Name = "Test User"
	data.Email = "test@example.com"
	data.Address = "Test Address, City, Country"
	data.Phone = "+1234567890"
	data.Company = "Test Company"
	data.Position = "Software Engineer"
	
	// 填充Skills数组
	for i := range data.Skills {
		data.Skills[i] = "Skill " + string(rune(i))
	}
	
	// 填充Projects数组
	for i := range data.Projects {
		data.Projects[i] = Project{
			Name:        "Project " + string(rune(i)),
			Description: "Description " + string(rune(i)),
			Duration:    i,
		}
		for j := range data.Projects[i].Team {
			data.Projects[i].Team[j] = "Member " + string(rune(j))
		}
	}
	
	// 填充MetaData
	for i := range data.MetaData {
		data.MetaData[i] = float64(i)
	}
	
	return data
}