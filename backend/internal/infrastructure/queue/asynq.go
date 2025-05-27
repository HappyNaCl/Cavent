package queue

import (
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/infrastructure/queue/handler"
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
				"default": 10,
				"critical": 3,
				"low": 1,
			},
			RetryDelayFunc: func(n int, err error, t *asynq.Task) time.Duration {
				return time.Duration(n) * time.Second
			},
		},
	)
}

func StartWorker(redisAddr string, eventView *handler.EventViewedHandler){
	srv := NewServer(redisAddr)
	mux := asynq.NewServeMux()
	mux.HandleFunc("event_viewed", eventView.Handle)

	if err := srv.Run(mux); err != nil {
		panic(err)
	}
}