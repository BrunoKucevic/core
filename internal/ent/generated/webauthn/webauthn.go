// Code generated by ent, DO NOT EDIT.

package webauthn

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the webauthn type in the database.
	Label = "webauthn"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldOwnerID holds the string denoting the owner_id field in the database.
	FieldOwnerID = "owner_id"
	// FieldCredentialID holds the string denoting the credential_id field in the database.
	FieldCredentialID = "credential_id"
	// FieldPublicKey holds the string denoting the public_key field in the database.
	FieldPublicKey = "public_key"
	// FieldAttestationType holds the string denoting the attestation_type field in the database.
	FieldAttestationType = "attestation_type"
	// FieldAaguid holds the string denoting the aaguid field in the database.
	FieldAaguid = "aaguid"
	// FieldSignCount holds the string denoting the sign_count field in the database.
	FieldSignCount = "sign_count"
	// FieldTransports holds the string denoting the transports field in the database.
	FieldTransports = "transports"
	// FieldBackupEligible holds the string denoting the backup_eligible field in the database.
	FieldBackupEligible = "backup_eligible"
	// FieldBackupState holds the string denoting the backup_state field in the database.
	FieldBackupState = "backup_state"
	// FieldUserPresent holds the string denoting the user_present field in the database.
	FieldUserPresent = "user_present"
	// FieldUserVerified holds the string denoting the user_verified field in the database.
	FieldUserVerified = "user_verified"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the webauthn in the database.
	Table = "webauthns"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "webauthns"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"
)

// Columns holds all SQL columns for webauthn fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldTags,
	FieldOwnerID,
	FieldCredentialID,
	FieldPublicKey,
	FieldAttestationType,
	FieldAaguid,
	FieldSignCount,
	FieldTransports,
	FieldBackupEligible,
	FieldBackupState,
	FieldUserPresent,
	FieldUserVerified,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/theopenlane/core/internal/ent/generated/runtime"
var (
	Hooks        [3]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// DefaultBackupEligible holds the default value on creation for the "backup_eligible" field.
	DefaultBackupEligible bool
	// DefaultBackupState holds the default value on creation for the "backup_state" field.
	DefaultBackupState bool
	// DefaultUserPresent holds the default value on creation for the "user_present" field.
	DefaultUserPresent bool
	// DefaultUserVerified holds the default value on creation for the "user_verified" field.
	DefaultUserVerified bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the Webauthn queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByOwnerID orders the results by the owner_id field.
func ByOwnerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerID, opts...).ToFunc()
}

// ByAttestationType orders the results by the attestation_type field.
func ByAttestationType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAttestationType, opts...).ToFunc()
}

// ByAaguid orders the results by the aaguid field.
func ByAaguid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAaguid, opts...).ToFunc()
}

// BySignCount orders the results by the sign_count field.
func BySignCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSignCount, opts...).ToFunc()
}

// ByBackupEligible orders the results by the backup_eligible field.
func ByBackupEligible(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBackupEligible, opts...).ToFunc()
}

// ByBackupState orders the results by the backup_state field.
func ByBackupState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBackupState, opts...).ToFunc()
}

// ByUserPresent orders the results by the user_present field.
func ByUserPresent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserPresent, opts...).ToFunc()
}

// ByUserVerified orders the results by the user_verified field.
func ByUserVerified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserVerified, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
