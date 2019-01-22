package resolver

import (
	"context"
	"log"

	"github.com/SnapshotApp/model"
	"github.com/graph-gophers/graphql-go"
)

// User resolves queries for a User
func (r *Resolver) User(ctx context.Context, args struct{ ID graphql.ID }) (*UserResolver, error) {
	// Get User by ID from user service
	user := &model.User{}
	log.Printf("Resolved user: %v", user)
	return &UserResolver{
		user,
	}, nil
}

// UserResolver resolves User objects
type UserResolver struct {
	user *model.User
}

// ID resolves the ID field of a User
func (r *UserResolver) ID() graphql.ID {
	return r.user.ID
}

// Name resolves the Name field of a User
func (r *UserResolver) Name() string {
	return r.user.Name
}

// ProfilePhotoURL resolves the ProfilePhotoURL field of a User
func (r *UserResolver) ProfilePhotoURL() string {
	return r.user.ProfilePhotoURL
}
