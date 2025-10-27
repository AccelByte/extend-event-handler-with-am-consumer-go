// Copyright (c) 2025 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package service

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "extend-async-messaging/pkg/pb/async_messaging"
)

type AsyncMessagingHandler struct {
	pb.UnimplementedAsyncMessagingConsumerServiceServer
}

func NewAsyncMessagingHandler() *AsyncMessagingHandler {
	return &AsyncMessagingHandler{}
}

func (h *AsyncMessagingHandler) OnMessage(ctx context.Context, msg *pb.ReceivedMessage) (*emptypb.Empty, error) {
	logrus.Infof("received message from topic: %s, body: %s", msg.Topic, msg.Body)

	return &emptypb.Empty{}, nil
}
