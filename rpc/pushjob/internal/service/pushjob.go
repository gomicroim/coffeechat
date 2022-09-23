package service

import (
	"context"

	pb "pushjob/api/pushjob"
)

type PushJobService struct {
	pb.UnimplementedPushJobServer
}

func NewPushJobService() *PushJobService {
	return &PushJobService{}
}

func (s *PushJobService) PushToSingleUser(ctx context.Context, req *pb.SingleUserRequest) (*pb.SingleUserReply, error) {
	return &pb.SingleUserReply{}, nil
}

func (s *PushJobService) PushToBatchUser(ctx context.Context, req *pb.BatchUserRequest) (*pb.BatchUserReply, error) {
	return &pb.BatchUserReply{}, nil
}

func (s *PushJobService) PushToAllOnlineUser(ctx context.Context, req *pb.AllOnlineUserReq) (*pb.AllOnlineUserReply, error) {
	return &pb.AllOnlineUserReply{}, nil
}
