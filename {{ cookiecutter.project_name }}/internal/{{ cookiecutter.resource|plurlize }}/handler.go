package {{ cookiecutter.resource|plurlize }}

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/{{ cookiecutter.project_owner }}/{{ cookiecutter.project_name }}/internal/gen/pb/v1"
)

// {{ cookiecutter.resource }}Handler implements the {{ cookiecutter.proto_service }}Server interface.
type {{ cookiecutter.resource }}Handler struct {
	pb.Unimplemented{{ cookiecutter.proto_service }}Server // Embedding for forward compatibility
}

// New{{ cookiecutter.proto_service }}Handler creates a new instance of {{ cookiecutter.resource }}Handler.
func New{{ cookiecutter.proto_service }}Handler() pb.{{ cookiecutter.proto_service }}Server {
	return &{{ cookiecutter.resource }}Handler{}
}

// {{ cookiecutter.proto_rpc }} handles an RPC request.
func (h *{{ cookiecutter.resource }}Handler) {{ cookiecutter.proto_rpc }}(ctx context.Context, req *pb.{{ cookiecutter.proto_message_request }}) (*pb.{{ cookiecutter.proto_message_response }}, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	// TODO: Implement the handler logic

	return &pb.{{ cookiecutter.proto_message_response }}{}, nil
}
