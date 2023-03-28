package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-m/internal/biz"
)

type myRepo struct {
	data *Data
	log  *log.Helper
}

func NewMyRepo(data *Data, logger log.Logger) biz.MyRepo {
	return &myRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m *myRepo) Save(ctx context.Context, req *biz.MyReq) (bool, error) {
	fmt.Printf("data %+v\n", req)
	return true, nil
}
