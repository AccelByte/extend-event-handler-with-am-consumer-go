package service

import (
    "context"
    "net"
    "testing"
    "time"

    pb "extend-async-messaging/pkg/pb/async_messaging"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func TestAsyncMessagingHandler_OnMessage_viaGRPC(t *testing.T) {
    t.Parallel()

    lis, err := net.Listen("tcp", "127.0.0.1:0")
    if err != nil {
        t.Fatalf("failed to listen: %v", err)
    }
    defer func() { _ = lis.Close() }()

    srv := grpc.NewServer()
    defer srv.Stop()

    handler := NewAsyncMessagingHandler()
    pb.RegisterAsyncMessagingConsumerServiceServer(srv, handler)

    go func() {
        // nolint: errcheck
        _ = srv.Serve(lis)
    }()

    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    conn, err := grpc.DialContext(ctx, lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        t.Fatalf("failed to dial: %v", err)
    }
    defer func() { _ = conn.Close() }()

    client := pb.NewAsyncMessagingConsumerServiceClient(conn)

    _, err = client.OnMessage(ctx, &pb.ReceivedMessage{Topic: "PlayerSaved", Body: "Saved."})
    if err != nil {
        t.Fatalf("OnMessage returned error: %v", err)
    }
}


