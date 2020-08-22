// Package models contains the types for schema 'xodb'.
package models

// Code generated by xo. DO NOT EDIT.

// MyEnumValue represents a row from '[custom my_enum_value]'.
type MyEnumValue struct {
	EnumValues string // enum_values
}

// MyEnumValues runs a custom query, returning results as MyEnumValue.
func MyEnumValues(db XODB, schema string, table, enum string) (*MyEnumValue, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`SUBSTRING(column_type, 6, CHAR_LENGTH(column_type) - 6) AS enum_values ` +
		`FROM information_schema.columns ` +
		`WHERE data_type = 'enum' AND table_schema = ? AND table_name = ? AND column_name = ?`

	// run query
	XOLog(sqlstr, schema, table, enum)
	var mev MyEnumValue
	err = db.QueryRow(sqlstr, schema, table, enum).Scan(&mev.EnumValues)
	if err != nil {
		return nil, err
	}

	return &mev, nil
}
