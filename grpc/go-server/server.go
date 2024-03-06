package main

import (
	"fmt"
	"io"
	"leetCode/grpc/proto"
	"sync"
	"time"

	"log"
	"net"

	"google.golang.org/grpc"
)

type MessageHandler struct {
	proto.UnimplementedPingPongServiceServer
	activeConnections map[proto.PingPongService_HandleMessageServer]struct{}
	mu                sync.Mutex
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		activeConnections: make(map[proto.PingPongService_HandleMessageServer]struct{}),
	}
}

func (s *MessageHandler) HandleMessage(stream proto.PingPongService_HandleMessageServer) error {
	s.mu.Lock()
	s.activeConnections[stream] = struct{}{}
	s.mu.Unlock()

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("EOF: %v", err)
			break
		}
		if err != nil {
			fmt.Printf("Error! %v", err)
			break
		}

		fmt.Printf("Received message: %s\n", msg.Body)
		response := &proto.PongMessage{Body: "Response from GO", Code: 100}
		if err := stream.Send(response); err != nil {
			log.Printf("Error sending response: %v", err)
			break
		}
		if err = stream.SendMsg(&proto.PongMessage{Body: "One more response from GO", Code: 999}); err != nil {
			log.Printf("Error sending active message: %v", err)
		}
	}
	s.mu.Lock()
	delete(s.activeConnections, stream)
	s.mu.Unlock()

	return nil
}

func (s *MessageHandler) broadcastMessage(msg *proto.PongMessage) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for conn := range s.activeConnections {
		err := conn.Send(msg)
		if err != nil {
			log.Printf("Error sending message to client: %v", err)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	messageHandler := NewMessageHandler()
	proto.RegisterPingPongServiceServer(srv, messageHandler)

	fmt.Println("Server is listening on port 50051...")

	go func() {
		for {
			time.Sleep(3 * time.Second)
			fmt.Printf("Pushing a message to clients: %d \n", len(messageHandler.activeConnections))
			messageHandler.broadcastMessage(&proto.PongMessage{Body: "Push message", Code: 98})
		}
	}()

	// Serve the gRPC messageHandler
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
