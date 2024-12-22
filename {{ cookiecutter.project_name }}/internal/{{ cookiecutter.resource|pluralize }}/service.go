package {{ cookiecutter.resource|pluralize }}

// service defines the interface for business logic
type service interface {}

// Service provides the concrete implementation of the service interface
type Service struct{}

// Todo does something
func (s *Service) Todo() (string) {
	// TODO: Implement business logic

	return "todo"
}
