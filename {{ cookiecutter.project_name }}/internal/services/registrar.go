package services

import (
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/{{ cookiecutter.project_owner }}/{{ cookiecutter.project_name }}/internal/_gen/pb/v1"
	"github.com/{{ cookiecutter.project_owner }}/{{ cookiecutter.project_name }}/internal/{{ cookiecutter.resource|pluralize }}"
)

// Register registers all gRPC service handlers and debugging capabilities with the server
func Register(server *grpc.Server) {
	pb.Register{{ cookiecutter.proto_service }}Server(server, {{ cookiecutter.resource|pluralize }}.New{{ cookiecutter.proto_service }}Handler())

	// Register reflection service for debugging
	reflection.Register(server)

	for key, value := range server.GetServiceInfo() {
		slog.Info("Service registered", "service", key, "methods", value.Methods, "metadata", value.Metadata)
	}
}
