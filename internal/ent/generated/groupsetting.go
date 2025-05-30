// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/group"
	"github.com/theopenlane/core/internal/ent/generated/groupsetting"
	"github.com/theopenlane/core/pkg/enums"
)

// GroupSetting is the model entity for the GroupSetting schema.
type GroupSetting struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// whether the group is visible to it's members / owners only or if it's searchable by anyone within the organization
	Visibility enums.Visibility `json:"visibility,omitempty"`
	// the policy governing ability to freely join a group, whether it requires an invitation, application, or either
	JoinPolicy enums.JoinPolicy `json:"join_policy,omitempty"`
	// whether to sync group members to slack groups
	SyncToSlack bool `json:"sync_to_slack,omitempty"`
	// whether to sync group members to github groups
	SyncToGithub bool `json:"sync_to_github,omitempty"`
	// the group id associated with the settings
	GroupID string `json:"group_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupSettingQuery when eager-loading is set.
	Edges        GroupSettingEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GroupSettingEdges holds the relations/edges for other nodes in the graph.
type GroupSettingEdges struct {
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupSettingEdges) GroupOrErr() (*Group, error) {
	if e.Group != nil {
		return e.Group, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: group.Label}
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupSetting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupsetting.FieldSyncToSlack, groupsetting.FieldSyncToGithub:
			values[i] = new(sql.NullBool)
		case groupsetting.FieldID, groupsetting.FieldCreatedBy, groupsetting.FieldUpdatedBy, groupsetting.FieldDeletedBy, groupsetting.FieldVisibility, groupsetting.FieldJoinPolicy, groupsetting.FieldGroupID:
			values[i] = new(sql.NullString)
		case groupsetting.FieldCreatedAt, groupsetting.FieldUpdatedAt, groupsetting.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupSetting fields.
func (gs *GroupSetting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupsetting.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				gs.ID = value.String
			}
		case groupsetting.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gs.CreatedAt = value.Time
			}
		case groupsetting.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gs.UpdatedAt = value.Time
			}
		case groupsetting.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				gs.CreatedBy = value.String
			}
		case groupsetting.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				gs.UpdatedBy = value.String
			}
		case groupsetting.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gs.DeletedAt = value.Time
			}
		case groupsetting.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				gs.DeletedBy = value.String
			}
		case groupsetting.FieldVisibility:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field visibility", values[i])
			} else if value.Valid {
				gs.Visibility = enums.Visibility(value.String)
			}
		case groupsetting.FieldJoinPolicy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field join_policy", values[i])
			} else if value.Valid {
				gs.JoinPolicy = enums.JoinPolicy(value.String)
			}
		case groupsetting.FieldSyncToSlack:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field sync_to_slack", values[i])
			} else if value.Valid {
				gs.SyncToSlack = value.Bool
			}
		case groupsetting.FieldSyncToGithub:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field sync_to_github", values[i])
			} else if value.Valid {
				gs.SyncToGithub = value.Bool
			}
		case groupsetting.FieldGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				gs.GroupID = value.String
			}
		default:
			gs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GroupSetting.
// This includes values selected through modifiers, order, etc.
func (gs *GroupSetting) Value(name string) (ent.Value, error) {
	return gs.selectValues.Get(name)
}

// QueryGroup queries the "group" edge of the GroupSetting entity.
func (gs *GroupSetting) QueryGroup() *GroupQuery {
	return NewGroupSettingClient(gs.config).QueryGroup(gs)
}

// Update returns a builder for updating this GroupSetting.
// Note that you need to call GroupSetting.Unwrap() before calling this method if this GroupSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (gs *GroupSetting) Update() *GroupSettingUpdateOne {
	return NewGroupSettingClient(gs.config).UpdateOne(gs)
}

// Unwrap unwraps the GroupSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gs *GroupSetting) Unwrap() *GroupSetting {
	_tx, ok := gs.config.driver.(*txDriver)
	if !ok {
		panic("generated: GroupSetting is not a transactional entity")
	}
	gs.config.driver = _tx.drv
	return gs
}

// String implements the fmt.Stringer.
func (gs *GroupSetting) String() string {
	var builder strings.Builder
	builder.WriteString("GroupSetting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gs.ID))
	builder.WriteString("created_at=")
	builder.WriteString(gs.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gs.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(gs.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(gs.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(gs.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(gs.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("visibility=")
	builder.WriteString(fmt.Sprintf("%v", gs.Visibility))
	builder.WriteString(", ")
	builder.WriteString("join_policy=")
	builder.WriteString(fmt.Sprintf("%v", gs.JoinPolicy))
	builder.WriteString(", ")
	builder.WriteString("sync_to_slack=")
	builder.WriteString(fmt.Sprintf("%v", gs.SyncToSlack))
	builder.WriteString(", ")
	builder.WriteString("sync_to_github=")
	builder.WriteString(fmt.Sprintf("%v", gs.SyncToGithub))
	builder.WriteString(", ")
	builder.WriteString("group_id=")
	builder.WriteString(gs.GroupID)
	builder.WriteByte(')')
	return builder.String()
}

// GroupSettings is a parsable slice of GroupSetting.
type GroupSettings []*GroupSetting
