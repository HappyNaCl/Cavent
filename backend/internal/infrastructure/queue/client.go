package queue

import "github.com/hibiken/asynq"

func InitClient(redisAddr string) (*asynq.Client, error) {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: redisAddr,
		DB: 1,
	})

	if err := client.Ping(); err != nil {
		return nil, err
	}

	return client, nil
}