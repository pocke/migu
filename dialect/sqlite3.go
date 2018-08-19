package dialect

import (
	"database/sql"
	"fmt"
	"strings"
)

type Sqlite3 struct {
	db *sql.DB
}

func NewSqlite3(db *sql.DB) Dialect {
	return &Sqlite3{
		db: db,
	}
}

// TODO: index
func (d *Sqlite3) ColumnSchema(tables ...string) ([]ColumnSchema, error) {
	parts := []string{
		"SELECT",
		"  m.tbl_name,",
		"  t.name,", // column name
		"  t.dflt_value",
		"  t.`notnull`",
		"  t.type",
		"FROM",
		"  sqlite_master as m,",
		"  pragma_table_info(m.name) as t",
	}
	var args []interface{}
	if len(tables) > 0 {
		placeholder := strings.Repeat(",?", len(tables))
		placeholder = placeholder[1:] // truncate the heading comma.
		parts = append(parts, fmt.Sprintf("WHERE m.table_name IN (%s)", placeholder))
		for _, t := range tables {
			args = append(args, t)
		}
	}
	parts = append(parts, "ORDER BY m.tbl_name, t.cid")
	query := strings.Join(parts, "\n")

	rows, err := d.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schemas []ColumnSchema
	for rows.Next() {
		schema := &sqlite3ColumnSchema{}
		if err := rows.Scan(&schema.tableName, &schema.columnName, &schema.columnDefault, &schema.isNotNullable, &schema.columnType); err != nil {
			return nil, err
		}
		schemas = append(schemas, schema)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return schemas, nil
}

func (d *Sqlite3) ColumnType(name string, size uint64, autoIncrement bool) (string, bool, bool) {
	// TODO
	return "", false, false
}

func (d *Sqlite3) DataType(name string, size uint64, unsigned bool, prec, scale int64) string {
	// TODO
	return ""
}

func (d *Sqlite3) Quote(s string) string {
	// TODO
	return ""
}

func (d *Sqlite3) QuoteString(s string) string {
	// TODO
	return ""
}

func (d *Sqlite3) AutoIncrement() string {
	return "AUTOINCREMENT"
}
