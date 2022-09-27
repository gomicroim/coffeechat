package service

import (
	"context"
	"pushjob/internal/mq"
	"strconv"

	pb "pushjob/api/pushjob"
)

type PushJobService struct {
	pb.UnimplementedPushJobServer

	producer mq.PushMsgProducer
}

func NewPushJobService(producer mq.PushMsgProducer) *PushJobService {
	return &PushJobService{
		producer: producer,
	}
}

func (s *PushJobService) PushToSingleUser(ctx context.Context, req *pb.SingleUserRequest) (*pb.SingleUserReply, error) {
	internalWsMsg := s.producer.BuildBatchUserReceiver(ctx, []string{strconv.FormatInt(req.UserId, 10)}, req.WsMsgType, req.Data)
	p, offset, err := s.producer.PushWsMsg(ctx, internalWsMsg)
	if err != nil {
		return nil, err
	}
	return &pb.SingleUserReply{
		Partition: p,
		Offset:    offset,
	}, nil
}

func (s *PushJobService) PushToBatchUser(ctx context.Context, req *pb.BatchUserRequest) (*pb.BatchUserReply, error) {
	ids := make([]string, len(req.UserId))
	for k, id := range req.UserId {
		ids[k] = strconv.FormatInt(id, 10)
	}
	internalWsMsg := s.producer.BuildBatchUserReceiver(ctx, ids, req.WsMsgType, req.Data)
	p, offset, err := s.producer.PushWsMsg(ctx, internalWsMsg)
	if err != nil {
		return nil, err
	}
	return &pb.BatchUserReply{
		Partition: p,
		Offset:    offset,
	}, nil
}

func (s *PushJobService) PushToAllOnlineUser(ctx context.Context, req *pb.AllOnlineUserReq) (*pb.AllOnlineUserReply, error) {
	internalWsMsg := s.producer.BuildAllUserReceiver(ctx, req.WsMsgType, req.Data)
	p, offset, err := s.producer.PushWsMsg(ctx, internalWsMsg)
	if err != nil {
		return nil, err
	}
	return &pb.AllOnlineUserReply{
		Partition: p,
		Offset:    offset,
	}, nil
}
