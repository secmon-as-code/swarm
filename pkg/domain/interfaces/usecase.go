package interfaces

import (
	"context"

	"github.com/m-mizutani/swarm/pkg/domain/model"
)

type UseCase interface {
	ObjectToSources(ctx context.Context, obj model.Object) ([]*model.Source, error)
	Load(ctx context.Context, requests []*model.LoadRequest) error
	Enqueue(ctx context.Context, req *model.EnqueueRequest) (*model.EnqueueResponse, error)
	Authorize(ctx context.Context, input *model.AuthPolicyInput) error
}
