package graph

import (
	pb "quesurifn/simple-auth/grpc/gen/proto/auth/v1"
	"quesurifn/simple-auth/pkg/jwt"
	"quesurifn/simple-auth/pkg/logger"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	jwt        *jwt.JWT
	logger     *logger.Logger
	grpcClient pb.AuthServiceClient
}
