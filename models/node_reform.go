// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type nodeTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *nodeTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("nodes").
func (v *nodeTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *nodeTableType) Columns() []string {
	return []string{"id", "type", "name"}
}

// NewStruct makes a new struct for that view or table.
func (v *nodeTableType) NewStruct() reform.Struct {
	return new(Node)
}

// NewRecord makes a new record for that table.
func (v *nodeTableType) NewRecord() reform.Record {
	return new(Node)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *nodeTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// NodeTable represents nodes view or table in SQL database.
var NodeTable = &nodeTableType{
	s: parse.StructInfo{Type: "Node", SQLSchema: "", SQLName: "nodes", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}, {Name: "Type", Type: "NodeType", Column: "type"}, {Name: "Name", Type: "string", Column: "name"}}, PKFieldIndex: 0},
	z: new(Node).Values(),
}

// String returns a string representation of this struct or record.
func (s Node) String() string {
	res := make([]string, 3)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Type: " + reform.Inspect(s.Type, true)
	res[2] = "Name: " + reform.Inspect(s.Name, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Node) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Type,
		s.Name,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Node) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Type,
		&s.Name,
	}
}

// View returns View object for that struct.
func (s *Node) View() reform.View {
	return NodeTable
}

// Table returns Table object for that record.
func (s *Node) Table() reform.Table {
	return NodeTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Node) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Node) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Node) HasPK() bool {
	return s.ID != NodeTable.z[NodeTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Node) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = NodeTable
	_ reform.Struct = (*Node)(nil)
	_ reform.Table  = NodeTable
	_ reform.Record = (*Node)(nil)
	_ fmt.Stringer  = (*Node)(nil)
)

type rDSNodeTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *rDSNodeTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("nodes").
func (v *rDSNodeTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *rDSNodeTableType) Columns() []string {
	return []string{"id", "type", "name", "region"}
}

// NewStruct makes a new struct for that view or table.
func (v *rDSNodeTableType) NewStruct() reform.Struct {
	return new(RDSNode)
}

// NewRecord makes a new record for that table.
func (v *rDSNodeTableType) NewRecord() reform.Record {
	return new(RDSNode)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *rDSNodeTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// RDSNodeTable represents nodes view or table in SQL database.
var RDSNodeTable = &rDSNodeTableType{
	s: parse.StructInfo{Type: "RDSNode", SQLSchema: "", SQLName: "nodes", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}, {Name: "Type", Type: "NodeType", Column: "type"}, {Name: "Name", Type: "string", Column: "name"}, {Name: "Region", Type: "string", Column: "region"}}, PKFieldIndex: 0},
	z: new(RDSNode).Values(),
}

// String returns a string representation of this struct or record.
func (s RDSNode) String() string {
	res := make([]string, 4)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Type: " + reform.Inspect(s.Type, true)
	res[2] = "Name: " + reform.Inspect(s.Name, true)
	res[3] = "Region: " + reform.Inspect(s.Region, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *RDSNode) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Type,
		s.Name,
		s.Region,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *RDSNode) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Type,
		&s.Name,
		&s.Region,
	}
}

// View returns View object for that struct.
func (s *RDSNode) View() reform.View {
	return RDSNodeTable
}

// Table returns Table object for that record.
func (s *RDSNode) Table() reform.Table {
	return RDSNodeTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *RDSNode) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *RDSNode) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *RDSNode) HasPK() bool {
	return s.ID != RDSNodeTable.z[RDSNodeTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *RDSNode) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = RDSNodeTable
	_ reform.Struct = (*RDSNode)(nil)
	_ reform.Table  = RDSNodeTable
	_ reform.Record = (*RDSNode)(nil)
	_ fmt.Stringer  = (*RDSNode)(nil)
)

type postgreSQLNodeTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *postgreSQLNodeTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("nodes").
func (v *postgreSQLNodeTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *postgreSQLNodeTableType) Columns() []string {
	return []string{"id", "type", "name"}
}

// NewStruct makes a new struct for that view or table.
func (v *postgreSQLNodeTableType) NewStruct() reform.Struct {
	return new(PostgreSQLNode)
}

// NewRecord makes a new record for that table.
func (v *postgreSQLNodeTableType) NewRecord() reform.Record {
	return new(PostgreSQLNode)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *postgreSQLNodeTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// PostgreSQLNodeTable represents nodes view or table in SQL database.
var PostgreSQLNodeTable = &postgreSQLNodeTableType{
	s: parse.StructInfo{Type: "PostgreSQLNode", SQLSchema: "", SQLName: "nodes", Fields: []parse.FieldInfo{{Name: "ID", Type: "int32", Column: "id"}, {Name: "Type", Type: "NodeType", Column: "type"}, {Name: "Name", Type: "string", Column: "name"}}, PKFieldIndex: 0},
	z: new(PostgreSQLNode).Values(),
}

// String returns a string representation of this struct or record.
func (s PostgreSQLNode) String() string {
	res := make([]string, 3)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Type: " + reform.Inspect(s.Type, true)
	res[2] = "Name: " + reform.Inspect(s.Name, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *PostgreSQLNode) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Type,
		s.Name,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *PostgreSQLNode) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Type,
		&s.Name,
	}
}

// View returns View object for that struct.
func (s *PostgreSQLNode) View() reform.View {
	return PostgreSQLNodeTable
}

// Table returns Table object for that record.
func (s *PostgreSQLNode) Table() reform.Table {
	return PostgreSQLNodeTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *PostgreSQLNode) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *PostgreSQLNode) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *PostgreSQLNode) HasPK() bool {
	return s.ID != PostgreSQLNodeTable.z[PostgreSQLNodeTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *PostgreSQLNode) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = PostgreSQLNodeTable
	_ reform.Struct = (*PostgreSQLNode)(nil)
	_ reform.Table  = PostgreSQLNodeTable
	_ reform.Record = (*PostgreSQLNode)(nil)
	_ fmt.Stringer  = (*PostgreSQLNode)(nil)
)

func init() {
	parse.AssertUpToDate(&NodeTable.s, new(Node))
	parse.AssertUpToDate(&RDSNodeTable.s, new(RDSNode))
	parse.AssertUpToDate(&PostgreSQLNodeTable.s, new(PostgreSQLNode))
}
