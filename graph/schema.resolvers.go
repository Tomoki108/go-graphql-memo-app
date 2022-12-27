package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/Tomoki108/go-graphql-memo-app/graph/model"
)

// CreateMemo is the resolver for the createMemo field.
func (r *mutationResolver) CreateMemo(ctx context.Context, input model.NewMemo) (*model.Memo, error) {
	memo := model.Memo{
		Title: input.Title,
		Body:  input.Body,
		Label: nil,
	}
	r.DB.Create(&memo)

	return &memo, nil
}

// Labels is the resolver for the labels field.
func (r *queryResolver) Labels(ctx context.Context) ([]*model.Label, error) {
	panic(fmt.Errorf("not implemented: Labels - labels"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
