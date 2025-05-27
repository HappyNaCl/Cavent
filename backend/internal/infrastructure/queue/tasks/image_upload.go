package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

type ImageUploadTask struct {
	FileBytes []byte `json:"fileBytes"`
	FileExt   string `json:"fileExt"`
	Path  	  string `json:"path"`
}

func NewImageUploadTask(fileBytes []byte, fileExt, path string) (*asynq.Task , error){
	payload := ImageUploadTask{
		FileBytes: fileBytes,
		FileExt:   fileExt,
		Path:  path,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeImageUpload, data), nil
}