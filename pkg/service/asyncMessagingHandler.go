// Copyright (c) 2025-2026 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package service

import (
	"context"
	"encoding/json"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "extend-async-messaging/pkg/pb/async_messaging"

	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclient/admin_game_record"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/cloudsave"
)

type PlayerJoinedEvent struct {
	EventType string `json:"eventType"`
	PlayerId  string `json:"playerId"`
	Timestamp string `json:"timestamp"`
}

type AsyncMessagingHandler struct {
	pb.UnimplementedAsyncMessagingConsumerServiceServer
	csStorage    *cloudsave.AdminGameRecordService
	namespace    string
	storeEnabled bool
}

func NewAsyncMessagingHandler(csStorage *cloudsave.AdminGameRecordService, namespace string, storeEnabled bool) *AsyncMessagingHandler {
	return &AsyncMessagingHandler{
		csStorage:    csStorage,
		namespace:    namespace,
		storeEnabled: storeEnabled,
	}
}

func (h *AsyncMessagingHandler) OnMessage(ctx context.Context, msg *pb.ReceivedMessage) (*emptypb.Empty, error) {
	slog.Default().Info("received message", "topic", msg.Topic, "body", msg.Body, "metadata", msg.Metadata)

	playerId, exists := msg.Metadata["PlayerId"]
	if exists {

		if h.storeEnabled {
			var event PlayerJoinedEvent
			err := json.Unmarshal([]byte(msg.Body), &event)
			if err != nil {
				slog.Default().Error("failed to unmarshal PlayerJoinedEvent", "error", err)
				return &emptypb.Empty{}, err
			}

			input := &admin_game_record.AdminPostGameRecordHandlerV1Params{
				Body:      &event,
				Key:       "player_joined_event_" + playerId,
				Namespace: h.namespace,
				Context:   ctx,
			}

			slog.Default().Info("storing PlayerJoinedEvent to CloudSave", "playerId", playerId, "namespace", h.namespace)
			_, err = h.csStorage.AdminPostGameRecordHandlerV1Short(input)
			if err != nil {
				return &emptypb.Empty{}, err
			}
		}

	} else {
		slog.Default().Info("PlayerId not found in metadata")
	}

	return &emptypb.Empty{}, nil
}
