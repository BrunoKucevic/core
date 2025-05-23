package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rs/zerolog/log"
	"github.com/theopenlane/core/internal/ent/generated"
	"github.com/theopenlane/core/internal/ent/generated/entitytype"
	"github.com/theopenlane/core/internal/graphapi/model"
	"github.com/theopenlane/utils/rout"
)

// CreateEntityType is the resolver for the createEntityType field.
func (r *mutationResolver) CreateEntityType(ctx context.Context, input generated.CreateEntityTypeInput) (*model.EntityTypeCreatePayload, error) {
	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, input.OwnerID); err != nil {
		log.Error().Err(err).Msg("failed to set organization in auth context")

		return nil, rout.NewMissingRequiredFieldError("owner_id")
	}

	res, err := withTransactionalMutation(ctx).EntityType.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "entitytype"})
	}

	return &model.EntityTypeCreatePayload{
		EntityType: res,
	}, nil
}

// CreateBulkEntityType is the resolver for the createBulkEntityType field.
func (r *mutationResolver) CreateBulkEntityType(ctx context.Context, input []*generated.CreateEntityTypeInput) (*model.EntityTypeBulkCreatePayload, error) {
	if len(input) == 0 {
		return nil, rout.NewMissingRequiredFieldError("input")
	}

	// set the organization in the auth context if its not done for us
	// this will choose the first input OwnerID when using a personal access token
	if err := setOrganizationInAuthContextBulkRequest(ctx, input); err != nil {
		log.Error().Err(err).Msg("failed to set organization in auth context")

		return nil, rout.NewMissingRequiredFieldError("owner_id")
	}

	return r.bulkCreateEntityType(ctx, input)
}

// CreateBulkCSVEntityType is the resolver for the createBulkCSVEntityType field.
func (r *mutationResolver) CreateBulkCSVEntityType(ctx context.Context, input graphql.Upload) (*model.EntityTypeBulkCreatePayload, error) {
	data, err := unmarshalBulkData[generated.CreateEntityTypeInput](input)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal bulk data")

		return nil, err
	}

	if len(data) == 0 {
		return nil, rout.NewMissingRequiredFieldError("input")
	}

	// set the organization in the auth context if its not done for us
	// this will choose the first input OwnerID when using a personal access token
	if err := setOrganizationInAuthContextBulkRequest(ctx, data); err != nil {
		log.Error().Err(err).Msg("failed to set organization in auth context")

		return nil, rout.NewMissingRequiredFieldError("owner_id")
	}

	return r.bulkCreateEntityType(ctx, data)
}

// UpdateEntityType is the resolver for the updateEntityType field.
func (r *mutationResolver) UpdateEntityType(ctx context.Context, id string, input generated.UpdateEntityTypeInput) (*model.EntityTypeUpdatePayload, error) {
	res, err := withTransactionalMutation(ctx).EntityType.Get(ctx, id)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "entitytype"})
	}

	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, &res.OwnerID); err != nil {
		log.Error().Err(err).Msg("failed to set organization in auth context")

		return nil, rout.ErrPermissionDenied
	}

	// setup update request
	req := res.Update().SetInput(input).AppendTags(input.AppendTags)

	res, err = req.Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "entitytype"})
	}

	return &model.EntityTypeUpdatePayload{
		EntityType: res,
	}, nil
}

// DeleteEntityType is the resolver for the deleteEntityType field.
func (r *mutationResolver) DeleteEntityType(ctx context.Context, id string) (*model.EntityTypeDeletePayload, error) {
	if err := withTransactionalMutation(ctx).EntityType.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, parseRequestError(err, action{action: ActionDelete, object: "entitytype"})
	}

	if err := generated.EntityTypeEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &model.EntityTypeDeletePayload{
		DeletedID: id,
	}, nil
}

// EntityType is the resolver for the entityType field.
func (r *queryResolver) EntityType(ctx context.Context, id string) (*generated.EntityType, error) {
	query, err := withTransactionalMutation(ctx).EntityType.Query().Where(entitytype.ID(id)).CollectFields(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionGet, object: "entitytype"})
	}

	res, err := query.Only(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionGet, object: "entitytype"})
	}

	return res, nil
}
