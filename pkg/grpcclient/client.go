/*
Copyright © 2022 Chief Of State

*/

package grpcclient

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

// ConnectionBuilder is a builder to create GRPC connection to the GRPC Server
type ConnectionBuilder interface {
	WithOptions(opts ...grpc.DialOption)
	WithInsecure()
	WithKeepAliveParams(params keepalive.ClientParameters)
	GetConn(ctx context.Context, addr string) (*grpc.ClientConn, error)
	GetTLSConn(ctx context.Context, addr string) (*grpc.ClientConn, error)
}

// Builder is grpc client client builder
type Builder struct {
	options              []grpc.DialOption
	transportCredentials credentials.TransportCredentials
}

// NewBuilder creates an instance of Builder
func NewBuilder() *Builder {
	return &Builder{}
}

// WithOptions set dial options
func (b *Builder) WithOptions(opts ...grpc.DialOption) *Builder {
	b.options = append(b.options, opts...)
	return b
}

// WithInsecure set the connection as insecure
func (b *Builder) WithInsecure() *Builder {
	b.options = append(b.options, grpc.WithInsecure()) // nolint
	return b
}

// WithBlock the dialing blocks until the  underlying connection is up.
// Without this, Dial returns immediately and connecting the server happens in background.
func (b *Builder) WithBlock() *Builder {
	b.options = append(b.options, grpc.WithBlock())
	return b
}

// WithKeepAliveParams set the keep alive params
// ClientParameters is used to set keepalive parameters on the client-side.
// These configure how the client will actively probe to notice when a
// connection is broken and send pings so intermediaries will be aware of the
// liveness of the connection. Make sure these parameters are set in
// coordination with the keepalive policy on the server, as incompatible
// settings can result in closing of connection.
func (b *Builder) WithKeepAliveParams(params keepalive.ClientParameters) *Builder {
	keepAlive := grpc.WithKeepaliveParams(params)
	b.options = append(b.options, keepAlive)
	return b
}

// WithClientTransportCredentials builds transport credentials for a gRPC client using the given properties.
func (b *Builder) WithClientTransportCredentials(insecureSkipVerify bool, certPool *x509.CertPool) *Builder {
	var tlsConf tls.Config

	if insecureSkipVerify {
		tlsConf.InsecureSkipVerify = true
		b.transportCredentials = credentials.NewTLS(&tlsConf)
		return b
	}

	tlsConf.RootCAs = certPool
	b.transportCredentials = credentials.NewTLS(&tlsConf)
	return b
}

// GetConn returns the client connection to the server
func (b *Builder) GetConn(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	if addr == "" {
		return nil, fmt.Errorf("target connection parameter missing. address = %s", addr)
	}
	cc, err := grpc.DialContext(ctx, addr, b.options...)

	if err != nil {
		return nil, fmt.Errorf("unable to connect to client. address = %s. error = %+v", addr, err)
	}
	return cc, nil
}

// GetTLSConn returns client connection to the server
func (b *Builder) GetTLSConn(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	b.options = append(b.options, grpc.WithTransportCredentials(b.transportCredentials))
	cc, err := grpc.DialContext(
		ctx,
		addr,
		b.options...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get tls conn. Unable to connect to client. address = %s: %w", addr, err)
	}
	return cc, nil
}
