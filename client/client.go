package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hawa130/metalx/client/config"
	"github.com/hawa130/metalx/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

var (
	sessionID string
)

func main() {
	cfg := config.Config()

	var opts []grpc.DialOption

	if cfg.Auth.EnableTLS {
		keypair, err := tls.X509KeyPair([]byte(cfg.Auth.PublicKey), []byte(cfg.Auth.PrivateKey))
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM([]byte(cfg.Auth.CACert))
		creds := credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{keypair},
			RootCAs:      caCertPool,
		})
		if err != nil {
			log.Fatalf("Failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	opts = append(opts, grpc.WithUnaryInterceptor(func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		p, ok := peer.FromContext(ctx)
		if !ok {
			return status.Errorf(codes.Unauthenticated, "could not find p info")
		}

		tlsInfo, ok := p.AuthInfo.(credentials.TLSInfo)
		if !ok {
			return status.Errorf(codes.Unauthenticated, "unexpected p transport credentials")
		}

		// Check controller's certificate CN
		for _, chain := range tlsInfo.State.VerifiedChains {
			for _, cert := range chain {
				cn := cert.Subject.CommonName
				if cn == "Controller" {
					return invoker(ctx, method, req, reply, cc, opts...)
				}
			}
		}
		return status.Errorf(codes.Unauthenticated, "invalid client certificate CN")
	}))

	opts = append(opts, grpc.WithUnaryInterceptor(func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		if sessionID != "" {
			ctx = metadata.AppendToOutgoingContext(ctx, "session-id", sessionID)
		}
		var header metadata.MD
		err := invoker(ctx, method, req, reply, cc, append(opts, grpc.Header(&header))...)
		if err != nil {
			log.Printf("Error during handle rpc invoke: %v", err)
		}
		if ids := header["set-session-id"]; len(ids) > 0 {
			sessionID = ids[0]
		}
		return err
	}))

	opts = append(opts, grpc.WithStreamInterceptor(func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		if sessionID != "" {
			ctx = metadata.AppendToOutgoingContext(ctx, "session-id", sessionID)
		}
		var header metadata.MD
		clientStream, err := streamer(ctx, desc, cc, method, append(opts, grpc.Header(&header))...)
		if err != nil {
			return nil, err
		}
		if ids := header["session-id"]; len(ids) > 0 {
			sessionID = ids[0]
		}
		return clientStream, nil
	}))

	conn, err := grpc.NewClient(cfg.ControllerAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("failed to close grpc client: %v", err)
		}
	}(conn)

	client := pb.NewControllerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Hour)
	defer cancel()

	fakeID := uuid.New()
	fakeIDBytes, _ := fakeID.MarshalBinary()

	_, err = client.RegisterAgent(ctx, &pb.AgentId{Id: fakeIDBytes})

	if err != nil {
		println("err: ", err.Error())
	}
	println("session id", sessionID)
}
