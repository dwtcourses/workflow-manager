package store

import (
	"context"
	"fmt"

	"github.com/Clever/workflow-manager/gen-go/models"
)

// Store defines the interface for persistence of Workflow Manager resources.
type Store interface {
	SaveWorkflowDefinition(ctx context.Context, wfd models.WorkflowDefinition) error
	UpdateWorkflowDefinition(ctx context.Context, wfd models.WorkflowDefinition) (models.WorkflowDefinition, error)
	GetWorkflowDefinitions(ctx context.Context) ([]models.WorkflowDefinition, error)
	GetWorkflowDefinitionVersions(ctx context.Context, name string) ([]models.WorkflowDefinition, error)
	GetWorkflowDefinition(ctx context.Context, name string, version int) (models.WorkflowDefinition, error)
	LatestWorkflowDefinition(ctx context.Context, name string) (models.WorkflowDefinition, error)

	SaveStateResource(ctx context.Context, res models.StateResource) error
	GetStateResource(ctx context.Context, name, namespace string) (models.StateResource, error)
	DeleteStateResource(ctx context.Context, name, namespace string) error

	SaveWorkflow(ctx context.Context, workflow models.Workflow) error
	DeleteWorkflowByID(ctx context.Context, workflowID string) error
	UpdateWorkflow(ctx context.Context, workflow models.Workflow) error
	GetWorkflowByID(ctx context.Context, id string) (models.Workflow, error)
}

type ConflictError struct {
	name string
}

func (e ConflictError) Error() string {
	return fmt.Sprintf("Already Exists: %s", e.name)
}

func NewConflict(name string) ConflictError {
	return ConflictError{name}
}

func NewNotFound(name string) models.NotFound {
	return models.NotFound{Message: name}
}

// InvalidPageTokenError is returned for workflow queries that contain a malformed or invalid page
// token.
type InvalidPageTokenError struct {
	cause error
}

// NewInvalidPageTokenError returns a new InvalidPageTokenError.
func NewInvalidPageTokenError(cause error) InvalidPageTokenError {
	return InvalidPageTokenError{
		cause: cause,
	}
}

// Error implements the error interface.
func (e InvalidPageTokenError) Error() string {
	return fmt.Sprintf("invalid page token: %v", e.cause)
}

// InvalidQueryStructureError is returned for workflow queries that have a disallowed structure,
//  like including both a Status and a ResolvedByUser value.
type InvalidQueryStructureError struct {
	cause string
}

// Error implements the error interface for InvalidQueryStructureError.
func (e InvalidQueryStructureError) Error() string {
	return fmt.Sprintf("Invalid query structure: %v", e.cause)
}

// NewInvalidQueryStructureError creates an InvalidQueryStructureErorr
func NewInvalidQueryStructureError(cause string) InvalidQueryStructureError {
	return InvalidQueryStructureError{cause}
}
