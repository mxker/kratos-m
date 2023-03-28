package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type MyReq struct {
	Name string `json:"name"`
}

type MyRepo interface {
	Save(context.Context, *MyReq) (bool, error)
}

type MyUseCase struct {
	repo MyRepo
	log  log.Logger
}

func NewMyUseCase(repo MyRepo, log log.Logger) *MyUseCase {
	return &MyUseCase{
		repo: repo,
		log:  log,
	}
}

func (m *MyUseCase) SaveM(ctx context.Context, req MyReq) (bool, error) {
	fmt.Printf(" SaveM req %+v\n", req)
	save, err := m.repo.Save(ctx, &req)
	fmt.Printf(" SaveM save %+v\n", save)
	if err != nil {
		return false, err
	}
	return save, err
}
