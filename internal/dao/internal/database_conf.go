// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DatabaseConfDao is the data access object for the table database_conf.
type DatabaseConfDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  DatabaseConfColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// DatabaseConfColumns defines and stores column names for the table database_conf.
type DatabaseConfColumns struct {
	DatabaseId string //
	DbName     string //
	UserName   string //
	Password   string //
	Host       string //
	Port       string //
	CreateTime string //
	UpdateTime string //
	DbType     string //
}

// databaseConfColumns holds the columns for the table database_conf.
var databaseConfColumns = DatabaseConfColumns{
	DatabaseId: "database_id",
	DbName:     "db_name",
	UserName:   "user_name",
	Password:   "password",
	Host:       "host",
	Port:       "port",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DbType:     "db_type",
}

// NewDatabaseConfDao creates and returns a new DAO object for table data access.
func NewDatabaseConfDao(handlers ...gdb.ModelHandler) *DatabaseConfDao {
	return &DatabaseConfDao{
		group:    "master",
		table:    "database_conf",
		columns:  databaseConfColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DatabaseConfDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DatabaseConfDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DatabaseConfDao) Columns() DatabaseConfColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DatabaseConfDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DatabaseConfDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *DatabaseConfDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
