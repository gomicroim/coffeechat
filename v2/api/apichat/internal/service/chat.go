package service

import (
	"chat/api/chat"
	"context"

	pb "apichat/api/chat/v1"
)

type ChatService struct {
	pb.UnimplementedChatServer

	chatClient chat.ChatClient
}

func NewChatService(client chat.ChatClient) *ChatService {
	return &ChatService{chatClient: client}
}

func (s *ChatService) SendMsg(ctx context.Context, req *pb.SendMsgRequest) (*pb.SendMsgReply, error) {
	return &pb.SendMsgReply{}, nil
}
func (s *ChatService) GetSyncMessage(ctx context.Context, req *pb.SyncMessageRequest) (*pb.SyncMessageReply, error) {
	return &pb.SyncMessageReply{}, nil
}
func (s *ChatService) GetRecentContactSession(ctx context.Context, req *pb.GetRecentSessionRequest) (*pb.GetRecentSessionReply, error) {
	return &pb.GetRecentSessionReply{}, nil
}
func (s *ChatService) GetMsgList(ctx context.Context, req *pb.GetMsgListRequest) (*pb.GetMsgListReply, error) {
	return &pb.GetMsgListReply{}, nil
}
func (s *ChatService) MsgReadAck(ctx context.Context, req *pb.MsgReadAckRequest) (*pb.MsgReadAckReply, error) {
	return &pb.MsgReadAckReply{}, nil
}
