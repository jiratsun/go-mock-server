//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var Host = newHostTable("go_mock_server", "host", "")

type hostTable struct {
	mysql.Table

	// Columns
	ID          mysql.ColumnInteger
	DomainName  mysql.ColumnString
	Alias_      mysql.ColumnString
	Description mysql.ColumnString
	IsActive    mysql.ColumnBool
	CreatedAt   mysql.ColumnTimestamp
	UpdatedAt   mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type HostTable struct {
	hostTable

	NEW hostTable
}

// AS creates new HostTable with assigned alias
func (a HostTable) AS(alias string) *HostTable {
	return newHostTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new HostTable with assigned schema name
func (a HostTable) FromSchema(schemaName string) *HostTable {
	return newHostTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new HostTable with assigned table prefix
func (a HostTable) WithPrefix(prefix string) *HostTable {
	return newHostTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new HostTable with assigned table suffix
func (a HostTable) WithSuffix(suffix string) *HostTable {
	return newHostTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newHostTable(schemaName, tableName, alias string) *HostTable {
	return &HostTable{
		hostTable: newHostTableImpl(schemaName, tableName, alias),
		NEW:       newHostTableImpl("", "new", ""),
	}
}

func newHostTableImpl(schemaName, tableName, alias string) hostTable {
	var (
		IDColumn          = mysql.IntegerColumn("id")
		DomainNameColumn  = mysql.StringColumn("domain_name")
		Alias_Column      = mysql.StringColumn("alias")
		DescriptionColumn = mysql.StringColumn("description")
		IsActiveColumn    = mysql.BoolColumn("is_active")
		CreatedAtColumn   = mysql.TimestampColumn("created_at")
		UpdatedAtColumn   = mysql.TimestampColumn("updated_at")
		allColumns        = mysql.ColumnList{IDColumn, DomainNameColumn, Alias_Column, DescriptionColumn, IsActiveColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns    = mysql.ColumnList{DomainNameColumn, Alias_Column, DescriptionColumn, IsActiveColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return hostTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		DomainName:  DomainNameColumn,
		Alias_:      Alias_Column,
		Description: DescriptionColumn,
		IsActive:    IsActiveColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
