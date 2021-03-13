package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/aeramu/graphql-federation-go/order-service/graph/generated"
	"github.com/aeramu/graphql-federation-go/order-service/graph/model"
)

func (r *mutationResolver) CreateOrder(ctx context.Context, name string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetOrder(ctx context.Context, id string) (*model.Order, error) {
	return &model.Order{
		ID:          id,
		Products:    []string{},
		TotalAmount: float64(10000),
		Status:      "finished",
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
