package dialect

import "database/sql"

var _ ColumnSchema = &sqlite3ColumnSchema{}

type sqlite3ColumnSchema struct {
	tableName     string
	columnName    string
	columnDefault sql.NullString
	isNotNullable bool
	columnType    string
}

func (schema *sqlite3ColumnSchema) TableName() string {
	return schema.tableName
}

func (schema *sqlite3ColumnSchema) ColumnName() string {
	return schema.columnName
}

func (schema *sqlite3ColumnSchema) DataType() string {
	return ""
}

func (schema *sqlite3ColumnSchema) GoType() string {
	return ""
}

func (schema *sqlite3ColumnSchema) IsDatetime() bool {
	return false
}

func (schema *sqlite3ColumnSchema) IsPrimaryKey() bool {
	return false
}

func (schema *sqlite3ColumnSchema) IsAutoIncrement() bool {
	return false
}

func (schema *sqlite3ColumnSchema) Index() (string, bool, bool) {
	return "", false, false
}

func (schema *sqlite3ColumnSchema) Default() (string, bool) {
	return "", false
}

func (schema *sqlite3ColumnSchema) Size() (int64, bool) {
	return 0, false
}

func (schema *sqlite3ColumnSchema) Precision() (int64, bool) {
	return 0, false
}

func (schema *sqlite3ColumnSchema) Scale() (int64, bool) {
	return 0, false
}

func (schema *sqlite3ColumnSchema) IsNullable() bool {
	return false
}

func (schema *sqlite3ColumnSchema) Extra() (string, bool) {
	return "", false
}

func (schema *sqlite3ColumnSchema) Comment() (string, bool) {
	return "", false
}
