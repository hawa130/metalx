package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/hawa130/metalx/server/config"
	"github.com/hawa130/metalx/server/controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func main() {
	cfg := config.Config()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", cfg.Port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	if cfg.Auth.EnableTLS {
		var keypair tls.Certificate
		keypair, err = tls.X509KeyPair([]byte(cfg.Auth.PublicKey), []byte(cfg.Auth.PrivateKey))
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM([]byte(cfg.Auth.CACert))
		creds := credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{keypair},
			ClientAuth:   tls.RequireAndVerifyClientCert,
			ClientCAs:    caCertPool,
		})
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
		opts = append(opts, grpc.UnaryInterceptor(func(
			ctx context.Context,
			req interface{},
			info *grpc.UnaryServerInfo,
			handler grpc.UnaryHandler,
		) (interface{}, error) {
			p, ok := peer.FromContext(ctx)
			if !ok {
				return nil, status.Errorf(codes.Unauthenticated, "could not find p info")
			}

			tlsInfo, ok := p.AuthInfo.(credentials.TLSInfo)
			if !ok {
				return nil, status.Errorf(codes.Unauthenticated, "unexpected p transport credentials")
			}

			// Check client's certificate CN
			for _, chain := range tlsInfo.State.VerifiedChains {
				for _, cert := range chain {
					cn := cert.Subject.CommonName
					log.Printf("Client CN: %s", cert.Subject.CommonName)
					if strings.HasPrefix(cn, "Agent-") {
						return handler(ctx, req)
					}
				}
			}
			return nil, status.Errorf(codes.Unauthenticated, "invalid client certificate CN")
		}))
	}

	grpcServer := grpc.NewServer(opts...)
	controller.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
