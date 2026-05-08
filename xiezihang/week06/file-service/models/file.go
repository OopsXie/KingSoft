package models

type FileMeta struct {
	ID         string `gorm:"primaryKey" json:"id"`
	Filename   string `json:"filename"`
	Size       int64  `json:"size"`
	Type       string `json:"type"`
	UploadTime int64  `json:"upload_time"`
}
