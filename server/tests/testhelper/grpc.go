package testhelper

import (
	"context"
	"net"
	"testing"

	"github.com/ekkx/tcmrsv-web/server/internal/config"
	"github.com/ekkx/tcmrsv-web/server/internal/shared/ctxhelper"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

// GRPCTestServer represents a test gRPC server
type GRPCTestServer struct {
	lis    *bufconn.Listener
	server *grpc.Server
}

// NewGRPCTestServer creates a new test gRPC server with optional interceptors
func NewGRPCTestServer(unaryInterceptors ...grpc.UnaryServerInterceptor) *GRPCTestServer {
	lis := bufconn.Listen(bufSize)
	
	// コンテキストにConfigを設定するインターセプター
	configInterceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		cfg, err := config.New()
		if err != nil {
			return nil, err
		}
		ctx = ctxhelper.SetConfig(ctx, cfg)
		return handler(ctx, req)
	}
	
	// インターセプターを結合
	allInterceptors := append([]grpc.UnaryServerInterceptor{configInterceptor}, unaryInterceptors...)
	
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(allInterceptors...),
	}
	
	server := grpc.NewServer(opts...)
	
	return &GRPCTestServer{
		lis:    lis,
		server: server,
	}
}

// GetServer returns the underlying gRPC server
func (s *GRPCTestServer) GetServer() *grpc.Server {
	return s.server
}

// Start starts the test server
func (s *GRPCTestServer) Start() {
	go func() {
		if err := s.server.Serve(s.lis); err != nil && err != grpc.ErrServerStopped {
			panic(err)
		}
	}()
}

// Stop stops the test server
func (s *GRPCTestServer) Stop() {
	s.server.Stop()
}

// GetDialer returns a dialer for the test server
func (s *GRPCTestServer) GetDialer() func(context.Context, string) (net.Conn, error) {
	return func(context.Context, string) (net.Conn, error) {
		return s.lis.Dial()
	}
}

// CreateTestClient creates a test client for the gRPC server
func CreateTestClient(t *testing.T, ctx context.Context, dialer func(context.Context, string) (net.Conn, error)) *grpc.ClientConn {
	conn, err := grpc.NewClient("passthrough:///bufnet",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	return conn
}

// SetAuthorizationHeader sets the authorization header in the context
func SetAuthorizationHeader(ctx context.Context, token string) context.Context {
	md := metadata.New(map[string]string{
		"authorization": "Bearer " + token,
	})
	return metadata.NewOutgoingContext(ctx, md)
}

// GetTestJWTToken generates a test JWT token
func GetTestJWTToken(userID string, jwtSecret string) string {
	token, err := GenerateTestJWT(userID, jwtSecret)
	if err != nil {
		panic(err)
	}
	return token
}

// AssertGRPCError asserts that the error is a gRPC error with the expected code and message
func AssertGRPCError(t *testing.T, err error, expectedCode codes.Code, expectedMessage string) {
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, expectedCode, st.Code())
	if expectedMessage != "" {
		require.Contains(t, st.Message(), expectedMessage)
	}
}

// CreateTestContext creates a test context with necessary configurations
func CreateTestContext(t *testing.T) context.Context {
	// Use existing GetContextWithConfig function which already sets up the config properly
	return GetContextWithConfig(t)
}