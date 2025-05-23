// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/note"
	"github.com/theopenlane/core/internal/ent/generated/organization"
	"github.com/theopenlane/core/internal/ent/generated/task"
)

// Note is the model entity for the Note schema.
type Note struct {
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
	// the ID of the organization owner of the object
	OwnerID string `json:"owner_id,omitempty"`
	// the text of the note
	Text string `json:"text,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NoteQuery when eager-loading is set.
	Edges         NoteEdges `json:"edges"`
	entity_notes  *string
	program_notes *string
	task_comments *string
	selectValues  sql.SelectValues
}

// NoteEdges holds the relations/edges for other nodes in the graph.
type NoteEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Organization `json:"owner,omitempty"`
	// Task holds the value of the task edge.
	Task *Task `json:"task,omitempty"`
	// Files holds the value of the files edge.
	Files []*File `json:"files,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedFiles map[string][]*File
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NoteEdges) OwnerOrErr() (*Organization, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NoteEdges) TaskOrErr() (*Task, error) {
	if e.Task != nil {
		return e.Task, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: task.Label}
	}
	return nil, &NotLoadedError{edge: "task"}
}

// FilesOrErr returns the Files value or an error if the edge
// was not loaded in eager-loading.
func (e NoteEdges) FilesOrErr() ([]*File, error) {
	if e.loadedTypes[2] {
		return e.Files, nil
	}
	return nil, &NotLoadedError{edge: "files"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Note) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case note.FieldID, note.FieldCreatedBy, note.FieldUpdatedBy, note.FieldDeletedBy, note.FieldDisplayID, note.FieldOwnerID, note.FieldText:
			values[i] = new(sql.NullString)
		case note.FieldCreatedAt, note.FieldUpdatedAt, note.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case note.ForeignKeys[0]: // entity_notes
			values[i] = new(sql.NullString)
		case note.ForeignKeys[1]: // program_notes
			values[i] = new(sql.NullString)
		case note.ForeignKeys[2]: // task_comments
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Note fields.
func (n *Note) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case note.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				n.ID = value.String
			}
		case note.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				n.CreatedAt = value.Time
			}
		case note.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				n.UpdatedAt = value.Time
			}
		case note.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				n.CreatedBy = value.String
			}
		case note.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				n.UpdatedBy = value.String
			}
		case note.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				n.DeletedAt = value.Time
			}
		case note.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				n.DeletedBy = value.String
			}
		case note.FieldDisplayID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_id", values[i])
			} else if value.Valid {
				n.DisplayID = value.String
			}
		case note.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				n.OwnerID = value.String
			}
		case note.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				n.Text = value.String
			}
		case note.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field entity_notes", values[i])
			} else if value.Valid {
				n.entity_notes = new(string)
				*n.entity_notes = value.String
			}
		case note.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field program_notes", values[i])
			} else if value.Valid {
				n.program_notes = new(string)
				*n.program_notes = value.String
			}
		case note.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field task_comments", values[i])
			} else if value.Valid {
				n.task_comments = new(string)
				*n.task_comments = value.String
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Note.
// This includes values selected through modifiers, order, etc.
func (n *Note) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Note entity.
func (n *Note) QueryOwner() *OrganizationQuery {
	return NewNoteClient(n.config).QueryOwner(n)
}

// QueryTask queries the "task" edge of the Note entity.
func (n *Note) QueryTask() *TaskQuery {
	return NewNoteClient(n.config).QueryTask(n)
}

// QueryFiles queries the "files" edge of the Note entity.
func (n *Note) QueryFiles() *FileQuery {
	return NewNoteClient(n.config).QueryFiles(n)
}

// Update returns a builder for updating this Note.
// Note that you need to call Note.Unwrap() before calling this method if this Note
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Note) Update() *NoteUpdateOne {
	return NewNoteClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Note entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Note) Unwrap() *Note {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("generated: Note is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Note) String() string {
	var builder strings.Builder
	builder.WriteString("Note(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("created_at=")
	builder.WriteString(n.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(n.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(n.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(n.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(n.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(n.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("display_id=")
	builder.WriteString(n.DisplayID)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(n.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(n.Text)
	builder.WriteByte(')')
	return builder.String()
}

// NamedFiles returns the Files named value or an error if the edge was not
// loaded in eager-loading with this name.
func (n *Note) NamedFiles(name string) ([]*File, error) {
	if n.Edges.namedFiles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := n.Edges.namedFiles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (n *Note) appendNamedFiles(name string, edges ...*File) {
	if n.Edges.namedFiles == nil {
		n.Edges.namedFiles = make(map[string][]*File)
	}
	if len(edges) == 0 {
		n.Edges.namedFiles[name] = []*File{}
	} else {
		n.Edges.namedFiles[name] = append(n.Edges.namedFiles[name], edges...)
	}
}

// Notes is a parsable slice of Note.
type Notes []*Note
