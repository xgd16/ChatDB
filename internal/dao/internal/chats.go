// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChatsDao is the data access object for the table chats.
type ChatsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ChatsColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ChatsColumns defines and stores column names for the table chats.
type ChatsColumns struct {
	ChatId     string //
	Name       string //
	Remark     string //
	Sort       string //
	CreateTime string //
	UpdateTime string //
}

// chatsColumns holds the columns for the table chats.
var chatsColumns = ChatsColumns{
	ChatId:     "chat_id",
	Name:       "name",
	Remark:     "remark",
	Sort:       "sort",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewChatsDao creates and returns a new DAO object for table data access.
func NewChatsDao(handlers ...gdb.ModelHandler) *ChatsDao {
	return &ChatsDao{
		group:    "master",
		table:    "chats",
		columns:  chatsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ChatsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ChatsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ChatsDao) Columns() ChatsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ChatsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ChatsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ChatsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
