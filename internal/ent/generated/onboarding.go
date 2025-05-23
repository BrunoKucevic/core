// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/onboarding"
	"github.com/theopenlane/core/internal/ent/generated/organization"
)

// Onboarding is the model entity for the Onboarding schema.
type Onboarding struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID string `json:"organization_id,omitempty"`
	// name of the company
	CompanyName string `json:"company_name,omitempty"`
	// domains associated with the company
	Domains []string `json:"domains,omitempty"`
	// details given about the company during the onboarding process, including things such as company size, sector, etc
	CompanyDetails map[string]interface{} `json:"company_details,omitempty"`
	// details given about the user during the onboarding process, including things such as name, job title, department, etc
	UserDetails map[string]interface{} `json:"user_details,omitempty"`
	// details given about the compliance requirements during the onboarding process, such as coming with existing policies, controls, risk assessments, etc
	Compliance map[string]interface{} `json:"compliance,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OnboardingQuery when eager-loading is set.
	Edges        OnboardingEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OnboardingEdges holds the relations/edges for other nodes in the graph.
type OnboardingEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OnboardingEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Onboarding) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case onboarding.FieldDomains, onboarding.FieldCompanyDetails, onboarding.FieldUserDetails, onboarding.FieldCompliance:
			values[i] = new([]byte)
		case onboarding.FieldID, onboarding.FieldDeletedBy, onboarding.FieldOrganizationID, onboarding.FieldCompanyName:
			values[i] = new(sql.NullString)
		case onboarding.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Onboarding fields.
func (o *Onboarding) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case onboarding.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				o.ID = value.String
			}
		case onboarding.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				o.DeletedAt = value.Time
			}
		case onboarding.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				o.DeletedBy = value.String
			}
		case onboarding.FieldOrganizationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				o.OrganizationID = value.String
			}
		case onboarding.FieldCompanyName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field company_name", values[i])
			} else if value.Valid {
				o.CompanyName = value.String
			}
		case onboarding.FieldDomains:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field domains", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.Domains); err != nil {
					return fmt.Errorf("unmarshal field domains: %w", err)
				}
			}
		case onboarding.FieldCompanyDetails:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field company_details", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.CompanyDetails); err != nil {
					return fmt.Errorf("unmarshal field company_details: %w", err)
				}
			}
		case onboarding.FieldUserDetails:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field user_details", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.UserDetails); err != nil {
					return fmt.Errorf("unmarshal field user_details: %w", err)
				}
			}
		case onboarding.FieldCompliance:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field compliance", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.Compliance); err != nil {
					return fmt.Errorf("unmarshal field compliance: %w", err)
				}
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Onboarding.
// This includes values selected through modifiers, order, etc.
func (o *Onboarding) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryOrganization queries the "organization" edge of the Onboarding entity.
func (o *Onboarding) QueryOrganization() *OrganizationQuery {
	return NewOnboardingClient(o.config).QueryOrganization(o)
}

// Update returns a builder for updating this Onboarding.
// Note that you need to call Onboarding.Unwrap() before calling this method if this Onboarding
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Onboarding) Update() *OnboardingUpdateOne {
	return NewOnboardingClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Onboarding entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Onboarding) Unwrap() *Onboarding {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("generated: Onboarding is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Onboarding) String() string {
	var builder strings.Builder
	builder.WriteString("Onboarding(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("deleted_at=")
	builder.WriteString(o.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(o.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(o.OrganizationID)
	builder.WriteString(", ")
	builder.WriteString("company_name=")
	builder.WriteString(o.CompanyName)
	builder.WriteString(", ")
	builder.WriteString("domains=")
	builder.WriteString(fmt.Sprintf("%v", o.Domains))
	builder.WriteString(", ")
	builder.WriteString("company_details=")
	builder.WriteString(fmt.Sprintf("%v", o.CompanyDetails))
	builder.WriteString(", ")
	builder.WriteString("user_details=")
	builder.WriteString(fmt.Sprintf("%v", o.UserDetails))
	builder.WriteString(", ")
	builder.WriteString("compliance=")
	builder.WriteString(fmt.Sprintf("%v", o.Compliance))
	builder.WriteByte(')')
	return builder.String()
}

// Onboardings is a parsable slice of Onboarding.
type Onboardings []*Onboarding
