package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/aeramu/graphql-federation-go/sales-service/graph/generated"
	"github.com/aeramu/graphql-federation-go/sales-service/graph/model"
)

func (r *entityResolver) FindSalesByID(ctx context.Context, id string) (*model.Sales, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
