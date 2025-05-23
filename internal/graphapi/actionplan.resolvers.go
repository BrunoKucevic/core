package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rs/zerolog/log"
	"github.com/theopenlane/core/internal/ent/generated"
	"github.com/theopenlane/core/internal/ent/generated/actionplan"
	gqlgenerated "github.com/theopenlane/core/internal/graphapi/generated"
	"github.com/theopenlane/core/internal/graphapi/model"
	"github.com/theopenlane/utils/rout"
)

// CreateActionPlan is the resolver for the createActionPlan field.
func (r *mutationResolver) CreateActionPlan(ctx context.Context, input generated.CreateActionPlanInput) (*model.ActionPlanCreatePayload, error) {
	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, input.OwnerID); err != nil {
		log.Error().Err(err).Msg("failed to set organization in auth context")

		return nil, rout.NewMissingRequiredFieldError("owner_id")
	}

	res, err := withTransactionalMutation(ctx).ActionPlan.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "actionplan"})
	}

	return &model.ActionPlanCreatePayload{
		ActionPlan: res,
	}, nil
}

// CreateBulkActionPlan is the resolver for the createBulkActionPlan field.
func (r *mutationResolver) CreateBulkActionPlan(ctx context.Context, input []*generated.CreateActionPlanInput) (*model.ActionPlanBulkCreatePayload, error) {
	if len(input) == 0 {
		return nil, rout.NewMissingRequiredFieldError("input")
	}

	// set the organization in the auth context if its not done for us
	// this will choose the first input OwnerID when using a personal access token
	if err := setOrganizationInAuthContextBulkRequest(ctx, input); err != nil {
		log.Error().Err(err).Msg("failed to set organization in auth context")

		return nil, rout.NewMissingRequiredFieldError("owner_id")
	}

	return r.bulkCreateActionPlan(ctx, input)
}

// CreateBulkCSVActionPlan is the resolver for the createBulkCSVActionPlan field.
func (r *mutationResolver) CreateBulkCSVActionPlan(ctx context.Context, input graphql.Upload) (*model.ActionPlanBulkCreatePayload, error) {
	data, err := unmarshalBulkData[generated.CreateActionPlanInput](input)
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

	return r.bulkCreateActionPlan(ctx, data)
}

// UpdateActionPlan is the resolver for the updateActionPlan field.
func (r *mutationResolver) UpdateActionPlan(ctx context.Context, id string, input generated.UpdateActionPlanInput) (*model.ActionPlanUpdatePayload, error) {
	res, err := withTransactionalMutation(ctx).ActionPlan.Get(ctx, id)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "actionplan"})
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
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "actionplan"})
	}

	return &model.ActionPlanUpdatePayload{
		ActionPlan: res,
	}, nil
}

// DeleteActionPlan is the resolver for the deleteActionPlan field.
func (r *mutationResolver) DeleteActionPlan(ctx context.Context, id string) (*model.ActionPlanDeletePayload, error) {
	if err := withTransactionalMutation(ctx).ActionPlan.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, parseRequestError(err, action{action: ActionDelete, object: "actionplan"})
	}

	if err := generated.ActionPlanEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &model.ActionPlanDeletePayload{
		DeletedID: id,
	}, nil
}

// ActionPlan is the resolver for the actionPlan field.
func (r *queryResolver) ActionPlan(ctx context.Context, id string) (*generated.ActionPlan, error) {
	query, err := withTransactionalMutation(ctx).ActionPlan.Query().Where(actionplan.ID(id)).CollectFields(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionGet, object: "actionplan"})
	}

	res, err := query.Only(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionGet, object: "actionplan"})
	}

	return res, nil
}

// Mutation returns gqlgenerated.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgenerated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
