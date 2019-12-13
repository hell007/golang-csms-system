package common

type UploadFile struct {
	FileData string `json:"fileData"`
	FileType string `json:"fileType"`
	FileSize int    `json:"fileSize"`
}
