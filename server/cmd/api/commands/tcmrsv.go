package commands

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/ekkx/tcmrsv-web/server/internal/config"

	auth_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/authorization"
	reservation_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/reservation"
	room_v1 "github.com/ekkx/tcmrsv-web/server/internal/api/v1/room"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(cfg *config.Config) error {
	pool, err := cfg.Database.Open()
	if err != nil {
		return err
	}
	defer pool.Close()

	port := 50051 // TODO: make this configurable
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	deps := GenerateServerDeps(pool)

	s := grpc.NewServer(
	// grpc.UnaryInterceptor(interceptor.Intercepter1),
	)

	auth_v1.RegisterAuthorizationServiceServer(s, deps.AuthorizationServiceServer)
	reservation_v1.RegisterReservationServiceServer(s, deps.ReservationServiceServer)
	room_v1.RegisterRoomServiceServer(s, deps.RoomServiceServer)

	reflection.Register(s)

	go func() {
		fmt.Printf("Server is listening on port %d\n", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Shutting down server...")
	s.GracefulStop()

	return nil
}
