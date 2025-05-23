// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/controlobjective"
	"github.com/theopenlane/core/internal/ent/generated/organization"
	"github.com/theopenlane/core/pkg/enums"
)

// ControlObjective is the model entity for the ControlObjective schema.
type ControlObjective struct {
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
	// a shortened prefixed id field to use as a human readable identifier
	DisplayID string `json:"display_id,omitempty"`
	// tags associated with the object
	Tags []string `json:"tags,omitempty"`
	// revision of the object as a semver (e.g. v1.0.0), by default any update will bump the patch version, unless the revision_bump field is set
	Revision string `json:"revision,omitempty"`
	// the ID of the organization owner of the object
	OwnerID string `json:"owner_id,omitempty"`
	// the name of the control objective
	Name string `json:"name,omitempty"`
	// the desired outcome or target of the control objective
	DesiredOutcome string `json:"desired_outcome,omitempty"`
	// status of the control objective
	Status enums.ObjectiveStatus `json:"status,omitempty"`
	// source of the control, e.g. framework, template, custom, etc.
	Source enums.ControlSource `json:"source,omitempty"`
	// type of the control objective e.g. compliance, financial, operational, etc.
	ControlObjectiveType string `json:"control_objective_type,omitempty"`
	// category of the control
	Category string `json:"category,omitempty"`
	// subcategory of the control
	Subcategory string `json:"subcategory,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ControlObjectiveQuery when eager-loading is set.
	Edges        ControlObjectiveEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ControlObjectiveEdges holds the relations/edges for other nodes in the graph.
type ControlObjectiveEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Organization `json:"owner,omitempty"`
	// groups that are blocked from viewing or editing the risk
	BlockedGroups []*Group `json:"blocked_groups,omitempty"`
	// provides edit access to the risk to members of the group
	Editors []*Group `json:"editors,omitempty"`
	// provides view access to the risk to members of the group
	Viewers []*Group `json:"viewers,omitempty"`
	// Programs holds the value of the programs edge.
	Programs []*Program `json:"programs,omitempty"`
	// Evidence holds the value of the evidence edge.
	Evidence []*Evidence `json:"evidence,omitempty"`
	// Controls holds the value of the controls edge.
	Controls []*Control `json:"controls,omitempty"`
	// Subcontrols holds the value of the subcontrols edge.
	Subcontrols []*Subcontrol `json:"subcontrols,omitempty"`
	// InternalPolicies holds the value of the internal_policies edge.
	InternalPolicies []*InternalPolicy `json:"internal_policies,omitempty"`
	// Procedures holds the value of the procedures edge.
	Procedures []*Procedure `json:"procedures,omitempty"`
	// Risks holds the value of the risks edge.
	Risks []*Risk `json:"risks,omitempty"`
	// Narratives holds the value of the narratives edge.
	Narratives []*Narrative `json:"narratives,omitempty"`
	// Tasks holds the value of the tasks edge.
	Tasks []*Task `json:"tasks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [13]bool
	// totalCount holds the count of the edges above.
	totalCount [13]map[string]int

	namedBlockedGroups    map[string][]*Group
	namedEditors          map[string][]*Group
	namedViewers          map[string][]*Group
	namedPrograms         map[string][]*Program
	namedEvidence         map[string][]*Evidence
	namedControls         map[string][]*Control
	namedSubcontrols      map[string][]*Subcontrol
	namedInternalPolicies map[string][]*InternalPolicy
	namedProcedures       map[string][]*Procedure
	namedRisks            map[string][]*Risk
	namedNarratives       map[string][]*Narrative
	namedTasks            map[string][]*Task
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ControlObjectiveEdges) OwnerOrErr() (*Organization, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// BlockedGroupsOrErr returns the BlockedGroups value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) BlockedGroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[1] {
		return e.BlockedGroups, nil
	}
	return nil, &NotLoadedError{edge: "blocked_groups"}
}

// EditorsOrErr returns the Editors value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) EditorsOrErr() ([]*Group, error) {
	if e.loadedTypes[2] {
		return e.Editors, nil
	}
	return nil, &NotLoadedError{edge: "editors"}
}

// ViewersOrErr returns the Viewers value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) ViewersOrErr() ([]*Group, error) {
	if e.loadedTypes[3] {
		return e.Viewers, nil
	}
	return nil, &NotLoadedError{edge: "viewers"}
}

// ProgramsOrErr returns the Programs value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) ProgramsOrErr() ([]*Program, error) {
	if e.loadedTypes[4] {
		return e.Programs, nil
	}
	return nil, &NotLoadedError{edge: "programs"}
}

// EvidenceOrErr returns the Evidence value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) EvidenceOrErr() ([]*Evidence, error) {
	if e.loadedTypes[5] {
		return e.Evidence, nil
	}
	return nil, &NotLoadedError{edge: "evidence"}
}

// ControlsOrErr returns the Controls value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) ControlsOrErr() ([]*Control, error) {
	if e.loadedTypes[6] {
		return e.Controls, nil
	}
	return nil, &NotLoadedError{edge: "controls"}
}

// SubcontrolsOrErr returns the Subcontrols value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) SubcontrolsOrErr() ([]*Subcontrol, error) {
	if e.loadedTypes[7] {
		return e.Subcontrols, nil
	}
	return nil, &NotLoadedError{edge: "subcontrols"}
}

// InternalPoliciesOrErr returns the InternalPolicies value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) InternalPoliciesOrErr() ([]*InternalPolicy, error) {
	if e.loadedTypes[8] {
		return e.InternalPolicies, nil
	}
	return nil, &NotLoadedError{edge: "internal_policies"}
}

// ProceduresOrErr returns the Procedures value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) ProceduresOrErr() ([]*Procedure, error) {
	if e.loadedTypes[9] {
		return e.Procedures, nil
	}
	return nil, &NotLoadedError{edge: "procedures"}
}

// RisksOrErr returns the Risks value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) RisksOrErr() ([]*Risk, error) {
	if e.loadedTypes[10] {
		return e.Risks, nil
	}
	return nil, &NotLoadedError{edge: "risks"}
}

// NarrativesOrErr returns the Narratives value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) NarrativesOrErr() ([]*Narrative, error) {
	if e.loadedTypes[11] {
		return e.Narratives, nil
	}
	return nil, &NotLoadedError{edge: "narratives"}
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e ControlObjectiveEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[12] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ControlObjective) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case controlobjective.FieldTags:
			values[i] = new([]byte)
		case controlobjective.FieldID, controlobjective.FieldCreatedBy, controlobjective.FieldUpdatedBy, controlobjective.FieldDeletedBy, controlobjective.FieldDisplayID, controlobjective.FieldRevision, controlobjective.FieldOwnerID, controlobjective.FieldName, controlobjective.FieldDesiredOutcome, controlobjective.FieldStatus, controlobjective.FieldSource, controlobjective.FieldControlObjectiveType, controlobjective.FieldCategory, controlobjective.FieldSubcategory:
			values[i] = new(sql.NullString)
		case controlobjective.FieldCreatedAt, controlobjective.FieldUpdatedAt, controlobjective.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ControlObjective fields.
func (co *ControlObjective) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case controlobjective.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				co.ID = value.String
			}
		case controlobjective.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				co.CreatedAt = value.Time
			}
		case controlobjective.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				co.UpdatedAt = value.Time
			}
		case controlobjective.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				co.CreatedBy = value.String
			}
		case controlobjective.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				co.UpdatedBy = value.String
			}
		case controlobjective.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				co.DeletedAt = value.Time
			}
		case controlobjective.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				co.DeletedBy = value.String
			}
		case controlobjective.FieldDisplayID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_id", values[i])
			} else if value.Valid {
				co.DisplayID = value.String
			}
		case controlobjective.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &co.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case controlobjective.FieldRevision:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field revision", values[i])
			} else if value.Valid {
				co.Revision = value.String
			}
		case controlobjective.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				co.OwnerID = value.String
			}
		case controlobjective.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				co.Name = value.String
			}
		case controlobjective.FieldDesiredOutcome:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desired_outcome", values[i])
			} else if value.Valid {
				co.DesiredOutcome = value.String
			}
		case controlobjective.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				co.Status = enums.ObjectiveStatus(value.String)
			}
		case controlobjective.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				co.Source = enums.ControlSource(value.String)
			}
		case controlobjective.FieldControlObjectiveType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field control_objective_type", values[i])
			} else if value.Valid {
				co.ControlObjectiveType = value.String
			}
		case controlobjective.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				co.Category = value.String
			}
		case controlobjective.FieldSubcategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subcategory", values[i])
			} else if value.Valid {
				co.Subcategory = value.String
			}
		default:
			co.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ControlObjective.
// This includes values selected through modifiers, order, etc.
func (co *ControlObjective) Value(name string) (ent.Value, error) {
	return co.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the ControlObjective entity.
func (co *ControlObjective) QueryOwner() *OrganizationQuery {
	return NewControlObjectiveClient(co.config).QueryOwner(co)
}

// QueryBlockedGroups queries the "blocked_groups" edge of the ControlObjective entity.
func (co *ControlObjective) QueryBlockedGroups() *GroupQuery {
	return NewControlObjectiveClient(co.config).QueryBlockedGroups(co)
}

// QueryEditors queries the "editors" edge of the ControlObjective entity.
func (co *ControlObjective) QueryEditors() *GroupQuery {
	return NewControlObjectiveClient(co.config).QueryEditors(co)
}

// QueryViewers queries the "viewers" edge of the ControlObjective entity.
func (co *ControlObjective) QueryViewers() *GroupQuery {
	return NewControlObjectiveClient(co.config).QueryViewers(co)
}

// QueryPrograms queries the "programs" edge of the ControlObjective entity.
func (co *ControlObjective) QueryPrograms() *ProgramQuery {
	return NewControlObjectiveClient(co.config).QueryPrograms(co)
}

// QueryEvidence queries the "evidence" edge of the ControlObjective entity.
func (co *ControlObjective) QueryEvidence() *EvidenceQuery {
	return NewControlObjectiveClient(co.config).QueryEvidence(co)
}

// QueryControls queries the "controls" edge of the ControlObjective entity.
func (co *ControlObjective) QueryControls() *ControlQuery {
	return NewControlObjectiveClient(co.config).QueryControls(co)
}

// QuerySubcontrols queries the "subcontrols" edge of the ControlObjective entity.
func (co *ControlObjective) QuerySubcontrols() *SubcontrolQuery {
	return NewControlObjectiveClient(co.config).QuerySubcontrols(co)
}

// QueryInternalPolicies queries the "internal_policies" edge of the ControlObjective entity.
func (co *ControlObjective) QueryInternalPolicies() *InternalPolicyQuery {
	return NewControlObjectiveClient(co.config).QueryInternalPolicies(co)
}

// QueryProcedures queries the "procedures" edge of the ControlObjective entity.
func (co *ControlObjective) QueryProcedures() *ProcedureQuery {
	return NewControlObjectiveClient(co.config).QueryProcedures(co)
}

// QueryRisks queries the "risks" edge of the ControlObjective entity.
func (co *ControlObjective) QueryRisks() *RiskQuery {
	return NewControlObjectiveClient(co.config).QueryRisks(co)
}

// QueryNarratives queries the "narratives" edge of the ControlObjective entity.
func (co *ControlObjective) QueryNarratives() *NarrativeQuery {
	return NewControlObjectiveClient(co.config).QueryNarratives(co)
}

// QueryTasks queries the "tasks" edge of the ControlObjective entity.
func (co *ControlObjective) QueryTasks() *TaskQuery {
	return NewControlObjectiveClient(co.config).QueryTasks(co)
}

// Update returns a builder for updating this ControlObjective.
// Note that you need to call ControlObjective.Unwrap() before calling this method if this ControlObjective
// was returned from a transaction, and the transaction was committed or rolled back.
func (co *ControlObjective) Update() *ControlObjectiveUpdateOne {
	return NewControlObjectiveClient(co.config).UpdateOne(co)
}

// Unwrap unwraps the ControlObjective entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (co *ControlObjective) Unwrap() *ControlObjective {
	_tx, ok := co.config.driver.(*txDriver)
	if !ok {
		panic("generated: ControlObjective is not a transactional entity")
	}
	co.config.driver = _tx.drv
	return co
}

// String implements the fmt.Stringer.
func (co *ControlObjective) String() string {
	var builder strings.Builder
	builder.WriteString("ControlObjective(")
	builder.WriteString(fmt.Sprintf("id=%v, ", co.ID))
	builder.WriteString("created_at=")
	builder.WriteString(co.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(co.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(co.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(co.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(co.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(co.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("display_id=")
	builder.WriteString(co.DisplayID)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", co.Tags))
	builder.WriteString(", ")
	builder.WriteString("revision=")
	builder.WriteString(co.Revision)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(co.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(co.Name)
	builder.WriteString(", ")
	builder.WriteString("desired_outcome=")
	builder.WriteString(co.DesiredOutcome)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", co.Status))
	builder.WriteString(", ")
	builder.WriteString("source=")
	builder.WriteString(fmt.Sprintf("%v", co.Source))
	builder.WriteString(", ")
	builder.WriteString("control_objective_type=")
	builder.WriteString(co.ControlObjectiveType)
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(co.Category)
	builder.WriteString(", ")
	builder.WriteString("subcategory=")
	builder.WriteString(co.Subcategory)
	builder.WriteByte(')')
	return builder.String()
}

// NamedBlockedGroups returns the BlockedGroups named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedBlockedGroups(name string) ([]*Group, error) {
	if co.Edges.namedBlockedGroups == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedBlockedGroups[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedBlockedGroups(name string, edges ...*Group) {
	if co.Edges.namedBlockedGroups == nil {
		co.Edges.namedBlockedGroups = make(map[string][]*Group)
	}
	if len(edges) == 0 {
		co.Edges.namedBlockedGroups[name] = []*Group{}
	} else {
		co.Edges.namedBlockedGroups[name] = append(co.Edges.namedBlockedGroups[name], edges...)
	}
}

// NamedEditors returns the Editors named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedEditors(name string) ([]*Group, error) {
	if co.Edges.namedEditors == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedEditors[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedEditors(name string, edges ...*Group) {
	if co.Edges.namedEditors == nil {
		co.Edges.namedEditors = make(map[string][]*Group)
	}
	if len(edges) == 0 {
		co.Edges.namedEditors[name] = []*Group{}
	} else {
		co.Edges.namedEditors[name] = append(co.Edges.namedEditors[name], edges...)
	}
}

// NamedViewers returns the Viewers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedViewers(name string) ([]*Group, error) {
	if co.Edges.namedViewers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedViewers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedViewers(name string, edges ...*Group) {
	if co.Edges.namedViewers == nil {
		co.Edges.namedViewers = make(map[string][]*Group)
	}
	if len(edges) == 0 {
		co.Edges.namedViewers[name] = []*Group{}
	} else {
		co.Edges.namedViewers[name] = append(co.Edges.namedViewers[name], edges...)
	}
}

// NamedPrograms returns the Programs named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedPrograms(name string) ([]*Program, error) {
	if co.Edges.namedPrograms == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedPrograms[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedPrograms(name string, edges ...*Program) {
	if co.Edges.namedPrograms == nil {
		co.Edges.namedPrograms = make(map[string][]*Program)
	}
	if len(edges) == 0 {
		co.Edges.namedPrograms[name] = []*Program{}
	} else {
		co.Edges.namedPrograms[name] = append(co.Edges.namedPrograms[name], edges...)
	}
}

// NamedEvidence returns the Evidence named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedEvidence(name string) ([]*Evidence, error) {
	if co.Edges.namedEvidence == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedEvidence[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedEvidence(name string, edges ...*Evidence) {
	if co.Edges.namedEvidence == nil {
		co.Edges.namedEvidence = make(map[string][]*Evidence)
	}
	if len(edges) == 0 {
		co.Edges.namedEvidence[name] = []*Evidence{}
	} else {
		co.Edges.namedEvidence[name] = append(co.Edges.namedEvidence[name], edges...)
	}
}

// NamedControls returns the Controls named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedControls(name string) ([]*Control, error) {
	if co.Edges.namedControls == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedControls[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedControls(name string, edges ...*Control) {
	if co.Edges.namedControls == nil {
		co.Edges.namedControls = make(map[string][]*Control)
	}
	if len(edges) == 0 {
		co.Edges.namedControls[name] = []*Control{}
	} else {
		co.Edges.namedControls[name] = append(co.Edges.namedControls[name], edges...)
	}
}

// NamedSubcontrols returns the Subcontrols named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedSubcontrols(name string) ([]*Subcontrol, error) {
	if co.Edges.namedSubcontrols == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedSubcontrols[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedSubcontrols(name string, edges ...*Subcontrol) {
	if co.Edges.namedSubcontrols == nil {
		co.Edges.namedSubcontrols = make(map[string][]*Subcontrol)
	}
	if len(edges) == 0 {
		co.Edges.namedSubcontrols[name] = []*Subcontrol{}
	} else {
		co.Edges.namedSubcontrols[name] = append(co.Edges.namedSubcontrols[name], edges...)
	}
}

// NamedInternalPolicies returns the InternalPolicies named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedInternalPolicies(name string) ([]*InternalPolicy, error) {
	if co.Edges.namedInternalPolicies == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedInternalPolicies[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedInternalPolicies(name string, edges ...*InternalPolicy) {
	if co.Edges.namedInternalPolicies == nil {
		co.Edges.namedInternalPolicies = make(map[string][]*InternalPolicy)
	}
	if len(edges) == 0 {
		co.Edges.namedInternalPolicies[name] = []*InternalPolicy{}
	} else {
		co.Edges.namedInternalPolicies[name] = append(co.Edges.namedInternalPolicies[name], edges...)
	}
}

// NamedProcedures returns the Procedures named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedProcedures(name string) ([]*Procedure, error) {
	if co.Edges.namedProcedures == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedProcedures[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedProcedures(name string, edges ...*Procedure) {
	if co.Edges.namedProcedures == nil {
		co.Edges.namedProcedures = make(map[string][]*Procedure)
	}
	if len(edges) == 0 {
		co.Edges.namedProcedures[name] = []*Procedure{}
	} else {
		co.Edges.namedProcedures[name] = append(co.Edges.namedProcedures[name], edges...)
	}
}

// NamedRisks returns the Risks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedRisks(name string) ([]*Risk, error) {
	if co.Edges.namedRisks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedRisks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedRisks(name string, edges ...*Risk) {
	if co.Edges.namedRisks == nil {
		co.Edges.namedRisks = make(map[string][]*Risk)
	}
	if len(edges) == 0 {
		co.Edges.namedRisks[name] = []*Risk{}
	} else {
		co.Edges.namedRisks[name] = append(co.Edges.namedRisks[name], edges...)
	}
}

// NamedNarratives returns the Narratives named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedNarratives(name string) ([]*Narrative, error) {
	if co.Edges.namedNarratives == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedNarratives[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedNarratives(name string, edges ...*Narrative) {
	if co.Edges.namedNarratives == nil {
		co.Edges.namedNarratives = make(map[string][]*Narrative)
	}
	if len(edges) == 0 {
		co.Edges.namedNarratives[name] = []*Narrative{}
	} else {
		co.Edges.namedNarratives[name] = append(co.Edges.namedNarratives[name], edges...)
	}
}

// NamedTasks returns the Tasks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (co *ControlObjective) NamedTasks(name string) ([]*Task, error) {
	if co.Edges.namedTasks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := co.Edges.namedTasks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (co *ControlObjective) appendNamedTasks(name string, edges ...*Task) {
	if co.Edges.namedTasks == nil {
		co.Edges.namedTasks = make(map[string][]*Task)
	}
	if len(edges) == 0 {
		co.Edges.namedTasks[name] = []*Task{}
	} else {
		co.Edges.namedTasks[name] = append(co.Edges.namedTasks[name], edges...)
	}
}

// ControlObjectives is a parsable slice of ControlObjective.
type ControlObjectives []*ControlObjective
