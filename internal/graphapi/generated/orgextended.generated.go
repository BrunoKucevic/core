// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgenerated

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/theopenlane/core/internal/graphapi/model"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputOrgMembersInput(ctx context.Context, obj any) (model.OrgMembersInput, error) {
	var it model.OrgMembersInput
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"role", "userID"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "role":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("role"))
			data, err := ec.unmarshalOOrgMembershipRole2ᚖgithubᚗcomᚋtheopenlaneᚋcoreᚋpkgᚋenumsᚐRole(ctx, v)
			if err != nil {
				return it, err
			}
			it.Role = data
		case "userID":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("userID"))
			data, err := ec.unmarshalNID2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.UserID = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNOrgMembersInput2ᚖgithubᚗcomᚋtheopenlaneᚋcoreᚋinternalᚋgraphapiᚋmodelᚐOrgMembersInput(ctx context.Context, v any) (*model.OrgMembersInput, error) {
	res, err := ec.unmarshalInputOrgMembersInput(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOOrgMembersInput2ᚕᚖgithubᚗcomᚋtheopenlaneᚋcoreᚋinternalᚋgraphapiᚋmodelᚐOrgMembersInputᚄ(ctx context.Context, v any) ([]*model.OrgMembersInput, error) {
	if v == nil {
		return nil, nil
	}
	var vSlice []any
	vSlice = graphql.CoerceList(v)
	var err error
	res := make([]*model.OrgMembersInput, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNOrgMembersInput2ᚖgithubᚗcomᚋtheopenlaneᚋcoreᚋinternalᚋgraphapiᚋmodelᚐOrgMembersInput(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

// endregion ***************************** type.gotpl *****************************
