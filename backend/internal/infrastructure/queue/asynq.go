package queue

import (
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/queue/handler"
	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/queue/tasks"
	"github.com/hibiken/asynq"
)

func NewServer(redisAddr string) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: redisAddr,
			DB: 1,
		},
		asynq.Config{
			Concurrency: 10,
			StrictPriority: true,
			Queues: map[string]int{
				tasks.TypeEventView: 10,
				tasks.TypeImageUpload: 10,
			},
			RetryDelayFunc: func(n int, err error, t *asynq.Task) time.Duration {
				return time.Duration(n) * time.Second
			},
		},
	)
}

func StartWorker(redisAddr string, eventView *handler.EventViewedHandler, imageUpload *handler.ImageUploadHandler){
	srv := NewServer(redisAddr)
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeEventView, eventView.Handle)
	mux.HandleFunc(tasks.TypeImageUpload, imageUpload.Handle)

	if err := srv.Run(mux); err != nil {
		panic(err)
	}
}