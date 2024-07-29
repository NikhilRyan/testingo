package mocks

import (
	"context"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
)

type RedisMock struct {
	Client  *redis.Client
	Mock    redismock.ClientMock
	Context context.Context
}

func NewRedisMock() *RedisMock {
	db, mock := redismock.NewClientMock()
	return &RedisMock{
		Client:  db,
		Mock:    mock,
		Context: context.Background(),
	}
}
