package actor

import (
	"context"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend"
)

// DBColumns represents the fields necessary to store an actor in a DB table.
//
// TODO(sqs): Add a field for organization ID.
type DBColumns struct {
	UserID                int32
	ExternalActorUsername string
	ExternalActorURL      string
}

// GQL returns the Actor GraphQL type for this actor.
func (c *DBColumns) GQL(ctx context.Context) (*graphqlbackend.Actor, error) {
	switch {
	case c.UserID != 0:
		user, err := graphqlbackend.UserByIDInt32(ctx, c.UserID)
		if err != nil {
			return nil, err
		}
		return &graphqlbackend.Actor{User: user}, nil
	case c.ExternalActorUsername != "" || c.ExternalActorURL != "":
		return &graphqlbackend.Actor{
			ExternalActor: &graphqlbackend.ExternalActor{
				Username_: c.ExternalActorUsername,
				URL_:      c.ExternalActorURL,
			},
		}, nil
	default:
		return nil, nil
	}
}