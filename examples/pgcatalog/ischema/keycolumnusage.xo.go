// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/yyoshiki41/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// KeyColumnUsage represents a row from 'information_schema.key_column_usage'.
type KeyColumnUsage struct {
	ConstraintCatalog          pgtypes.SQLIdentifier  `json:"constraint_catalog"`            // constraint_catalog
	ConstraintSchema           pgtypes.SQLIdentifier  `json:"constraint_schema"`             // constraint_schema
	ConstraintName             pgtypes.SQLIdentifier  `json:"constraint_name"`               // constraint_name
	TableCatalog               pgtypes.SQLIdentifier  `json:"table_catalog"`                 // table_catalog
	TableSchema                pgtypes.SQLIdentifier  `json:"table_schema"`                  // table_schema
	TableName                  pgtypes.SQLIdentifier  `json:"table_name"`                    // table_name
	ColumnName                 pgtypes.SQLIdentifier  `json:"column_name"`                   // column_name
	OrdinalPosition            pgtypes.CardinalNumber `json:"ordinal_position"`              // ordinal_position
	PositionInUniqueConstraint pgtypes.CardinalNumber `json:"position_in_unique_constraint"` // position_in_unique_constraint
}
