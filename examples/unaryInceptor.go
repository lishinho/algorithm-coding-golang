package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func orderUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) {
	// golang中元数据以正常map的形式表述，可以通过以下形式创建，相同键的元数据会合并成列表
	metadata.New(map[string]string{"Key1": "val1", "key2": "val2"})
}

// stream handler
