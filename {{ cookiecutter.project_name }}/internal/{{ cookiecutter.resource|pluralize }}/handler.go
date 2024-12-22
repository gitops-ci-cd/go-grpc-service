package {{ cookiecutter.resource|pluralize }}

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/{{ cookiecutter.project_owner }}/{{ cookiecutter.project_name }}/internal/_gen/pb/v1"
)

// Handler implements the {{ cookiecutter.proto_service }}Server interface
type Handler struct {
	// Embedding for forward compatibility
	pb.Unimplemented{{ cookiecutter.proto_service }}Server
	Service service
}

// Register associates the handler with the given gRPC server
func (h *Handler) Register(server *grpc.Server) {
	pb.Register{{ cookiecutter.proto_service }}Server(server, h)
}

// Todo handles an RPC request.
func (h *Handler) Todo()(ctx context.Context, req *pb.TodoRequest) (*pb.TodoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	h.Service.Todo()

	return &pb.TodoResponse{}, nil
}
