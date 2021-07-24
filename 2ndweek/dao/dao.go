package dao

import (
	"database/sql"

	"github.com/pkg/errors"
)

var _ DaoInterface = (*Dao)(nil)

type Dao struct {
}

func NewDao() DaoInterface {
	return &Dao{}
}

type DaoInterface interface {
	Query(s string) error
}

// query 可能是业务代码返回的错误，也可能是第三方返回sql.ErrNoRows错误
// 此处只返回sql.ErrNoRows
func (d *Dao) Query(s string) error {
	return errors.Wrap(sql.ErrNoRows, s)
}
