package service

import (
	"context"
	"fmt"
	"kratos-m/internal/biz"

	pb "kratos-m/api/mykratos/v1"
)

type MyService struct {
	pb.UnimplementedMyServer

	uc *biz.MyUseCase
}

func NewMyService(uc *biz.MyUseCase) *MyService {
	return &MyService{uc: uc}
}

func (s *MyService) CreateMy(ctx context.Context, req *pb.CreateMyRequest) (*pb.CreateMyReply, error) {
	m, err := s.uc.SaveM(ctx, biz.MyReq{Name: req.Name})
	if err != nil {
		return nil, err
	}
	fmt.Printf("CreateM %+v\n", m)

	return &pb.CreateMyReply{}, nil
}
func (s *MyService) UpdateMy(ctx context.Context, req *pb.UpdateMyRequest) (*pb.UpdateMyReply, error) {
	return &pb.UpdateMyReply{}, nil
}
func (s *MyService) DeleteMy(ctx context.Context, req *pb.DeleteMyRequest) (*pb.DeleteMyReply, error) {
	return &pb.DeleteMyReply{}, nil
}
func (s *MyService) GetMy(ctx context.Context, req *pb.GetMyRequest) (*pb.GetMyReply, error) {
	fmt.Printf("GetMy %+v\n", req)
	m, err := s.uc.SaveM(ctx, biz.MyReq{Name: string(req.Id)})
	if err != nil {
		return nil, err
	}
	fmt.Printf("m %+v\n", m)
	return &pb.GetMyReply{}, nil
}
func (s *MyService) ListMy(ctx context.Context, req *pb.ListMyRequest) (*pb.ListMyReply, error) {
	return &pb.ListMyReply{}, nil
}
