package client

import (
	"context"

	"github.com/docker/stacks/pkg/types"
)

// StackAPIClient defines the client interface for managing Stacks
type StackAPIClient interface {
	StackCreate(ctx context.Context, stack types.StackSpec, options types.StackCreateOptions) (types.StackCreateResponse, error)
	StackInspect(ctx context.Context, id string) (types.Stack, error)
	StackList(ctx context.Context, options types.StackListOptions) ([]types.Stack, error)
	StackUpdate(ctx context.Context, id string, version types.Version, spec types.StackSpec, options types.StackUpdateOptions) error
	StackDelete(ctx context.Context, id string) error
}
