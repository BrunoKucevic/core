// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/organizationhistory"
	"github.com/theopenlane/entx/history"
)

// OrganizationHistory is the model entity for the OrganizationHistory schema.
type OrganizationHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// HistoryTime holds the value of the "history_time" field.
	HistoryTime time.Time `json:"history_time,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref string `json:"ref,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation history.OpType `json:"operation,omitempty"`
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
	// tags associated with the object
	Tags []string `json:"tags,omitempty"`
	// the name of the organization
	Name string `json:"name,omitempty"`
	// The organization's displayed 'friendly' name
	DisplayName string `json:"display_name,omitempty"`
	// An optional description of the organization
	Description string `json:"description,omitempty"`
	// The ID of the parent organization for the organization.
	ParentOrganizationID string `json:"parent_organization_id,omitempty"`
	// orgs directly associated with a user
	PersonalOrg bool `json:"personal_org,omitempty"`
	// URL of the user's remote avatar
	AvatarRemoteURL *string `json:"avatar_remote_url,omitempty"`
	// The organizations's local avatar file id, takes precedence over the avatar remote URL
	AvatarLocalFileID *string `json:"avatar_local_file_id,omitempty"`
	// The time the user's (local) avatar was last updated
	AvatarUpdatedAt *time.Time `json:"avatar_updated_at,omitempty"`
	// Whether the organization has a dedicated database
	DedicatedDb  bool `json:"dedicated_db,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrganizationHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organizationhistory.FieldTags:
			values[i] = new([]byte)
		case organizationhistory.FieldOperation:
			values[i] = new(history.OpType)
		case organizationhistory.FieldPersonalOrg, organizationhistory.FieldDedicatedDb:
			values[i] = new(sql.NullBool)
		case organizationhistory.FieldID, organizationhistory.FieldRef, organizationhistory.FieldCreatedBy, organizationhistory.FieldUpdatedBy, organizationhistory.FieldDeletedBy, organizationhistory.FieldName, organizationhistory.FieldDisplayName, organizationhistory.FieldDescription, organizationhistory.FieldParentOrganizationID, organizationhistory.FieldAvatarRemoteURL, organizationhistory.FieldAvatarLocalFileID:
			values[i] = new(sql.NullString)
		case organizationhistory.FieldHistoryTime, organizationhistory.FieldCreatedAt, organizationhistory.FieldUpdatedAt, organizationhistory.FieldDeletedAt, organizationhistory.FieldAvatarUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrganizationHistory fields.
func (oh *OrganizationHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organizationhistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				oh.ID = value.String
			}
		case organizationhistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				oh.HistoryTime = value.Time
			}
		case organizationhistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				oh.Ref = value.String
			}
		case organizationhistory.FieldOperation:
			if value, ok := values[i].(*history.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				oh.Operation = *value
			}
		case organizationhistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				oh.CreatedAt = value.Time
			}
		case organizationhistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				oh.UpdatedAt = value.Time
			}
		case organizationhistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				oh.CreatedBy = value.String
			}
		case organizationhistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				oh.UpdatedBy = value.String
			}
		case organizationhistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				oh.DeletedAt = value.Time
			}
		case organizationhistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				oh.DeletedBy = value.String
			}
		case organizationhistory.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &oh.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case organizationhistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				oh.Name = value.String
			}
		case organizationhistory.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				oh.DisplayName = value.String
			}
		case organizationhistory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				oh.Description = value.String
			}
		case organizationhistory.FieldParentOrganizationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_organization_id", values[i])
			} else if value.Valid {
				oh.ParentOrganizationID = value.String
			}
		case organizationhistory.FieldPersonalOrg:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field personal_org", values[i])
			} else if value.Valid {
				oh.PersonalOrg = value.Bool
			}
		case organizationhistory.FieldAvatarRemoteURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_remote_url", values[i])
			} else if value.Valid {
				oh.AvatarRemoteURL = new(string)
				*oh.AvatarRemoteURL = value.String
			}
		case organizationhistory.FieldAvatarLocalFileID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_local_file_id", values[i])
			} else if value.Valid {
				oh.AvatarLocalFileID = new(string)
				*oh.AvatarLocalFileID = value.String
			}
		case organizationhistory.FieldAvatarUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_updated_at", values[i])
			} else if value.Valid {
				oh.AvatarUpdatedAt = new(time.Time)
				*oh.AvatarUpdatedAt = value.Time
			}
		case organizationhistory.FieldDedicatedDb:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field dedicated_db", values[i])
			} else if value.Valid {
				oh.DedicatedDb = value.Bool
			}
		default:
			oh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrganizationHistory.
// This includes values selected through modifiers, order, etc.
func (oh *OrganizationHistory) Value(name string) (ent.Value, error) {
	return oh.selectValues.Get(name)
}

// Update returns a builder for updating this OrganizationHistory.
// Note that you need to call OrganizationHistory.Unwrap() before calling this method if this OrganizationHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (oh *OrganizationHistory) Update() *OrganizationHistoryUpdateOne {
	return NewOrganizationHistoryClient(oh.config).UpdateOne(oh)
}

// Unwrap unwraps the OrganizationHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oh *OrganizationHistory) Unwrap() *OrganizationHistory {
	_tx, ok := oh.config.driver.(*txDriver)
	if !ok {
		panic("generated: OrganizationHistory is not a transactional entity")
	}
	oh.config.driver = _tx.drv
	return oh
}

// String implements the fmt.Stringer.
func (oh *OrganizationHistory) String() string {
	var builder strings.Builder
	builder.WriteString("OrganizationHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(oh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(oh.Ref)
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", oh.Operation))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(oh.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(oh.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(oh.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(oh.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(oh.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(oh.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", oh.Tags))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(oh.Name)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(oh.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(oh.Description)
	builder.WriteString(", ")
	builder.WriteString("parent_organization_id=")
	builder.WriteString(oh.ParentOrganizationID)
	builder.WriteString(", ")
	builder.WriteString("personal_org=")
	builder.WriteString(fmt.Sprintf("%v", oh.PersonalOrg))
	builder.WriteString(", ")
	if v := oh.AvatarRemoteURL; v != nil {
		builder.WriteString("avatar_remote_url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oh.AvatarLocalFileID; v != nil {
		builder.WriteString("avatar_local_file_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := oh.AvatarUpdatedAt; v != nil {
		builder.WriteString("avatar_updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("dedicated_db=")
	builder.WriteString(fmt.Sprintf("%v", oh.DedicatedDb))
	builder.WriteByte(')')
	return builder.String()
}

// OrganizationHistories is a parsable slice of OrganizationHistory.
type OrganizationHistories []*OrganizationHistory
