// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/yyoshiki41/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// Column represents a row from 'information_schema.columns'.
type Column struct {
	TableCatalog           pgtypes.SQLIdentifier  `json:"table_catalog"`            // table_catalog
	TableSchema            pgtypes.SQLIdentifier  `json:"table_schema"`             // table_schema
	TableName              pgtypes.SQLIdentifier  `json:"table_name"`               // table_name
	ColumnName             pgtypes.SQLIdentifier  `json:"column_name"`              // column_name
	OrdinalPosition        pgtypes.CardinalNumber `json:"ordinal_position"`         // ordinal_position
	ColumnDefault          pgtypes.CharacterData  `json:"column_default"`           // column_default
	IsNullable             pgtypes.YesOrNo        `json:"is_nullable"`              // is_nullable
	DataType               pgtypes.CharacterData  `json:"data_type"`                // data_type
	CharacterMaximumLength pgtypes.CardinalNumber `json:"character_maximum_length"` // character_maximum_length
	CharacterOctetLength   pgtypes.CardinalNumber `json:"character_octet_length"`   // character_octet_length
	NumericPrecision       pgtypes.CardinalNumber `json:"numeric_precision"`        // numeric_precision
	NumericPrecisionRadix  pgtypes.CardinalNumber `json:"numeric_precision_radix"`  // numeric_precision_radix
	NumericScale           pgtypes.CardinalNumber `json:"numeric_scale"`            // numeric_scale
	DatetimePrecision      pgtypes.CardinalNumber `json:"datetime_precision"`       // datetime_precision
	IntervalType           pgtypes.CharacterData  `json:"interval_type"`            // interval_type
	IntervalPrecision      pgtypes.CardinalNumber `json:"interval_precision"`       // interval_precision
	CharacterSetCatalog    pgtypes.SQLIdentifier  `json:"character_set_catalog"`    // character_set_catalog
	CharacterSetSchema     pgtypes.SQLIdentifier  `json:"character_set_schema"`     // character_set_schema
	CharacterSetName       pgtypes.SQLIdentifier  `json:"character_set_name"`       // character_set_name
	CollationCatalog       pgtypes.SQLIdentifier  `json:"collation_catalog"`        // collation_catalog
	CollationSchema        pgtypes.SQLIdentifier  `json:"collation_schema"`         // collation_schema
	CollationName          pgtypes.SQLIdentifier  `json:"collation_name"`           // collation_name
	DomainCatalog          pgtypes.SQLIdentifier  `json:"domain_catalog"`           // domain_catalog
	DomainSchema           pgtypes.SQLIdentifier  `json:"domain_schema"`            // domain_schema
	DomainName             pgtypes.SQLIdentifier  `json:"domain_name"`              // domain_name
	UdtCatalog             pgtypes.SQLIdentifier  `json:"udt_catalog"`              // udt_catalog
	UdtSchema              pgtypes.SQLIdentifier  `json:"udt_schema"`               // udt_schema
	UdtName                pgtypes.SQLIdentifier  `json:"udt_name"`                 // udt_name
	ScopeCatalog           pgtypes.SQLIdentifier  `json:"scope_catalog"`            // scope_catalog
	ScopeSchema            pgtypes.SQLIdentifier  `json:"scope_schema"`             // scope_schema
	ScopeName              pgtypes.SQLIdentifier  `json:"scope_name"`               // scope_name
	MaximumCardinality     pgtypes.CardinalNumber `json:"maximum_cardinality"`      // maximum_cardinality
	DtdIdentifier          pgtypes.SQLIdentifier  `json:"dtd_identifier"`           // dtd_identifier
	IsSelfReferencing      pgtypes.YesOrNo        `json:"is_self_referencing"`      // is_self_referencing
	IsIdentity             pgtypes.YesOrNo        `json:"is_identity"`              // is_identity
	IdentityGeneration     pgtypes.CharacterData  `json:"identity_generation"`      // identity_generation
	IdentityStart          pgtypes.CharacterData  `json:"identity_start"`           // identity_start
	IdentityIncrement      pgtypes.CharacterData  `json:"identity_increment"`       // identity_increment
	IdentityMaximum        pgtypes.CharacterData  `json:"identity_maximum"`         // identity_maximum
	IdentityMinimum        pgtypes.CharacterData  `json:"identity_minimum"`         // identity_minimum
	IdentityCycle          pgtypes.YesOrNo        `json:"identity_cycle"`           // identity_cycle
	IsGenerated            pgtypes.CharacterData  `json:"is_generated"`             // is_generated
	GenerationExpression   pgtypes.CharacterData  `json:"generation_expression"`    // generation_expression
	IsUpdatable            pgtypes.YesOrNo        `json:"is_updatable"`             // is_updatable
}
