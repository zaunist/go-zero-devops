package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ USERModel = (*customUSERModel)(nil)

var (
	userFieldNames    = builder.RawFieldNames(&USER{})
	userRows          = strings.Join(userFieldNames, ",")
	cacheUserIdPrefix = "cache#user#id#"
)

type (
	// USERModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUSERModel.
	USERModel interface {
		uSERModel
		FindByName(ctx context.Context, name string) (*USER, error)
		FindAll() ([]*USER, error)
	}

	customUSERModel struct {
		*defaultUSERModel
		sqlc.CachedConn
	}
)

// NewUSERModel returns a model for the database table.
func NewUSERModel(conn sqlx.SqlConn, c cache.CacheConf) USERModel {
	return &customUSERModel{
		defaultUSERModel: newUSERModel(conn),
		CachedConn:       sqlc.NewConn(conn, c),
	}
}

func (m *customUSERModel) FindByName(ctx context.Context, name string) (*USER, error) {
	var resp USER
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, name)
	err := m.QueryRow(&resp, userIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where username = ? limit 1", userRows, m.table)
		return conn.QueryRow(v, query, name)
	})

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUSERModel) FindAll() ([]*USER, error) {
	var resp []*USER

	query := fmt.Sprintf("select * from %s", m.table)
	err := m.QueryRowsNoCache(&resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
