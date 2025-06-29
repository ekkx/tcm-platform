package commands

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/ekkx/tcmrsv-web/server/internal/config"
	userv1 "github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/pb/user/v1/userv1connect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type UserServer struct {}

func (s *UserServer) GetUser(ctx context.Context, req *connect.Request[userv1.GetUserRequest]) (*connect.Response[userv1.GetUserResponse], error) {
    res := connect.NewResponse(&userv1.GetUserResponse{
        User: &userv1.User{

        },
    })
    return res, nil
}

func (s *UserServer) CreateSlaveUser(ctx context.Context, req *connect.Request[userv1.CreateSlaveUserRequest]) (*connect.Response[userv1.CreateSlaveUserResponse], error) {
    return nil, nil
}

func (s *UserServer) DeleteSlaveUser(ctx context.Context, req *connect.Request[userv1.DeleteSlaveUserRequest]) (*connect.Response[userv1.DeleteSlaveUserResponse], error) {
    return nil, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *connect.Request[userv1.DeleteUserRequest]) (*connect.Response[userv1.DeleteUserResponse], error) {
    return nil, nil
}

func (s *UserServer) ListSlaveUsers(ctx context.Context, req *connect.Request[userv1.ListSlaveUsersRequest]) (*connect.Response[userv1.ListSlaveUsersResponse], error) {
    return nil, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *connect.Request[userv1.UpdateUserRequest]) (*connect.Response[userv1.UpdateUserResponse], error) {
    return nil, nil
}

func Run(cfg *config.Config) error {
    userServer := &UserServer{}
    mux := http.NewServeMux()
    path, handler := userv1connect.NewUserServiceHandler(userServer)
    mux.Handle(path, handler)
    corsHandler := cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{}))
    http.ListenAndServe(fmt.Sprintf(":%d", 50051), corsHandler)
    return nil
}
