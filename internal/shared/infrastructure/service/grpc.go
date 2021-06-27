package service

import (
	"context"
	"crypto/tls"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func NewGRPCServer(lg *logrus.Entry) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime: time.Minute * 1,
		}),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
				err := handler(srv, ss)
				if err != nil {
					//sentry.CaptureException(err)
				}
				return err
			},
			grpc_logrus.StreamServerInterceptor(lg),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				res, err := handler(ctx, req)
				if err != nil {
					//sentry.CaptureException(err)
				}
				return res, err
			},
			grpc_logrus.UnaryServerInterceptor(lg),
		)),
	}
	if os.Getenv("APP_ENV") == "production" {
		backendCert, _ := ioutil.ReadFile("./tls/tls.cert")
		backendKey, _ := ioutil.ReadFile("./tls/tls.key")

		// Generate Certificate struct
		cert, err := tls.X509KeyPair(backendCert, backendKey)
		if err != nil {
			log.Fatalf("failed to parse certificate: %v", err)
		}

		// Create credentials
		creds := credentials.NewServerTLSFromCert(&cert)

		// Use Credentials in gRPC server options
		serverOption := grpc.Creds(creds)
		opts = append(opts, serverOption)
	}
	grpcServer := grpc.NewServer(opts...)
	return grpcServer
}
