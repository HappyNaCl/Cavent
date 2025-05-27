package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/errors"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/queue/tasks"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type ImageUploadHandler struct {}

func NewImageUploadHandler() *ImageUploadHandler {
	return &ImageUploadHandler{}
}

func (h *ImageUploadHandler) Handle(ctx context.Context ,task *asynq.Task) error {
	zap.L().Sugar().Infof("Processing task: %s", task.Type())
	var payload tasks.ImageUploadTask

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	supabaseUrl := os.Getenv("SUPABASE_PROJECT_URL")
	if supabaseUrl == "" {
		return errors.ErrSupabaseUrlMissing
	}

	supabaseKey := os.Getenv("SUPABASE_API_KEY")
	if supabaseKey == "" {
		return errors.ErrSupabaseKeyMissing
	}

	supabaseBucket := os.Getenv("SUPABASE_BUCKET_NAME")
	if supabaseBucket == "" {
		return errors.ErrSupabaseBucketMissing
	}

	url := fmt.Sprintf("%s/storage/v1/object/%s/%s", supabaseUrl, supabaseBucket, payload.Path)

	supabaseReq, err := http.NewRequest("POST", url, bytes.NewReader(payload.FileBytes))
	if err != nil {
		return err
	}
	supabaseReq.Header.Set("Authorization", "Bearer " + supabaseKey)
	supabaseReq.Header.Set("Content-Type", "image/" + payload.FileExt[1:])


	client := &http.Client{}
	resp, err := client.Do(supabaseReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			zap.L().Sugar().Errorf("Failed to read response body: %v", err)
		}
		bodyString := string(bodyBytes)
		zap.L().Sugar().Errorf("Error: %s\n", bodyString)
		return errors.ErrSupabaseRequestFailed
	}

	return nil
}