package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/theopenlane/core/internal/ent/generated"
	"github.com/theopenlane/gqlgen-plugins/graphutils"
	"github.com/theopenlane/iam/auth"
	"github.com/theopenlane/utils/rout"
)

// Note is the resolver for the Note field.
func (r *createEntityInputResolver) Note(ctx context.Context, obj *generated.CreateEntityInput, data *generated.CreateNoteInput) error {
	c := withTransactionalMutation(ctx)

	note, err := c.Note.Create().SetInput(*data).Save(ctx)
	if err != nil {
		return parseRequestError(err, action{action: ActionCreate, object: "entity"})
	}

	obj.NoteIDs = []string{note.ID}

	return nil
}

// Note is the resolver for the Note field.
func (r *updateEntityInputResolver) Note(ctx context.Context, obj *generated.UpdateEntityInput, data *generated.CreateNoteInput) error {
	// get the organization id from the context and if not found, get it from the entity
	// this should only happen when a personal access token is used to authenticate
	ownerID, err := auth.GetOrganizationIDFromContext(ctx)
	if err != nil || ownerID == "" {
		// get the entity id from the context
		id := graphutils.GetStringInputVariableByName(ctx, "id")
		if id == nil {
			return rout.NewMissingRequiredFieldError("entity id")
		}

		// get the entity in order to set the organization in the auth context
		res, err := withTransactionalMutation(ctx).Entity.Get(ctx, *id)
		if err != nil {
			return parseRequestError(err, action{action: ActionUpdate, object: "entity"})
		}

		// set the organization in the auth context if its not done for us
		if err := setOrganizationInAuthContext(ctx, &res.OwnerID); err != nil {
			log.Error().Err(err).Msg("failed to set organization in auth context")
			return rout.ErrPermissionDenied
		}
	}

	c := withTransactionalMutation(ctx)

	note, err := c.Note.Create().SetInput(*data).Save(ctx)
	if err != nil {
		return parseRequestError(err, action{action: ActionUpdate, object: "entity"})
	}

	obj.AddNoteIDs = []string{note.ID}

	return nil
}
