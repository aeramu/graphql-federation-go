package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/aeramu/graphql-federation-go/sales-service/graph/generated"
	"github.com/aeramu/graphql-federation-go/sales-service/graph/model"
)

func (r *mutationResolver) CreateSales(ctx context.Context, name string, phone string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetSales(ctx context.Context, id string) (*model.Sales, error) {
	return &model.Sales{
		ID:    id,
		Name:  "Sulthan Alam",
		Code:  "YOG012",
		Phone: "081288340024",
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
