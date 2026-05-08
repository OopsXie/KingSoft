package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"homework/model"
)

func AppendToJSONFile(data model.APIResult) error {
	t := time.Now().Format("2006_01_02")
	filePath := filepath.Join("data", t+".json")

	var records []model.APIResult
	_ = os.MkdirAll("data", os.ModePerm)

	if b, err := os.ReadFile(filePath); err == nil && len(b) > 0 {
		_ = json.Unmarshal(b, &records)
	}

	records = append(records, data)
	newData, _ := json.MarshalIndent(records, "", "  ")
	return os.WriteFile(filePath, newData, 0644)
}
