// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/templatehistory"
	"github.com/theopenlane/core/pkg/enums"
	"github.com/theopenlane/entx/history"
)

// TemplateHistory is the model entity for the TemplateHistory schema.
type TemplateHistory struct {
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
	// the organization id that owns the object
	OwnerID string `json:"owner_id,omitempty"`
	// the name of the template
	Name string `json:"name,omitempty"`
	// the type of the template, either a provided template or an implementation (document)
	TemplateType enums.DocumentType `json:"template_type,omitempty"`
	// the description of the template
	Description string `json:"description,omitempty"`
	// the jsonschema object of the template
	Jsonconfig map[string]interface{} `json:"jsonconfig,omitempty"`
	// the uischema for the template to render in the UI
	Uischema     map[string]interface{} `json:"uischema,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TemplateHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case templatehistory.FieldTags, templatehistory.FieldJsonconfig, templatehistory.FieldUischema:
			values[i] = new([]byte)
		case templatehistory.FieldOperation:
			values[i] = new(history.OpType)
		case templatehistory.FieldID, templatehistory.FieldRef, templatehistory.FieldCreatedBy, templatehistory.FieldUpdatedBy, templatehistory.FieldDeletedBy, templatehistory.FieldOwnerID, templatehistory.FieldName, templatehistory.FieldTemplateType, templatehistory.FieldDescription:
			values[i] = new(sql.NullString)
		case templatehistory.FieldHistoryTime, templatehistory.FieldCreatedAt, templatehistory.FieldUpdatedAt, templatehistory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TemplateHistory fields.
func (th *TemplateHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case templatehistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				th.ID = value.String
			}
		case templatehistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				th.HistoryTime = value.Time
			}
		case templatehistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				th.Ref = value.String
			}
		case templatehistory.FieldOperation:
			if value, ok := values[i].(*history.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				th.Operation = *value
			}
		case templatehistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				th.CreatedAt = value.Time
			}
		case templatehistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				th.UpdatedAt = value.Time
			}
		case templatehistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				th.CreatedBy = value.String
			}
		case templatehistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				th.UpdatedBy = value.String
			}
		case templatehistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				th.DeletedAt = value.Time
			}
		case templatehistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				th.DeletedBy = value.String
			}
		case templatehistory.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &th.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case templatehistory.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				th.OwnerID = value.String
			}
		case templatehistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				th.Name = value.String
			}
		case templatehistory.FieldTemplateType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field template_type", values[i])
			} else if value.Valid {
				th.TemplateType = enums.DocumentType(value.String)
			}
		case templatehistory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				th.Description = value.String
			}
		case templatehistory.FieldJsonconfig:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field jsonconfig", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &th.Jsonconfig); err != nil {
					return fmt.Errorf("unmarshal field jsonconfig: %w", err)
				}
			}
		case templatehistory.FieldUischema:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field uischema", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &th.Uischema); err != nil {
					return fmt.Errorf("unmarshal field uischema: %w", err)
				}
			}
		default:
			th.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TemplateHistory.
// This includes values selected through modifiers, order, etc.
func (th *TemplateHistory) Value(name string) (ent.Value, error) {
	return th.selectValues.Get(name)
}

// Update returns a builder for updating this TemplateHistory.
// Note that you need to call TemplateHistory.Unwrap() before calling this method if this TemplateHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (th *TemplateHistory) Update() *TemplateHistoryUpdateOne {
	return NewTemplateHistoryClient(th.config).UpdateOne(th)
}

// Unwrap unwraps the TemplateHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (th *TemplateHistory) Unwrap() *TemplateHistory {
	_tx, ok := th.config.driver.(*txDriver)
	if !ok {
		panic("generated: TemplateHistory is not a transactional entity")
	}
	th.config.driver = _tx.drv
	return th
}

// String implements the fmt.Stringer.
func (th *TemplateHistory) String() string {
	var builder strings.Builder
	builder.WriteString("TemplateHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", th.ID))
	builder.WriteString("history_time=")
	builder.WriteString(th.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(th.Ref)
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", th.Operation))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(th.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(th.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(th.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(th.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(th.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(th.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", th.Tags))
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(th.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(th.Name)
	builder.WriteString(", ")
	builder.WriteString("template_type=")
	builder.WriteString(fmt.Sprintf("%v", th.TemplateType))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(th.Description)
	builder.WriteString(", ")
	builder.WriteString("jsonconfig=")
	builder.WriteString(fmt.Sprintf("%v", th.Jsonconfig))
	builder.WriteString(", ")
	builder.WriteString("uischema=")
	builder.WriteString(fmt.Sprintf("%v", th.Uischema))
	builder.WriteByte(')')
	return builder.String()
}

// TemplateHistories is a parsable slice of TemplateHistory.
type TemplateHistories []*TemplateHistory
