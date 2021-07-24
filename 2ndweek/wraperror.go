package main

import (
	"database/sql"
	"fmt"
	"wraperror/dao"

	"github.com/pkg/errors"
)

func main() {
	dao := dao.NewDao()
	if err := queryHelper(dao); err != nil {
		// 查询失败，可能是包内代码的错误，也可能第三方包的错误
		fmt.Printf("fail to query: %v\n\n", err)
		// 根因
		fmt.Printf("original error: %T %v\n\n", errors.Cause(err), errors.Cause(err))
		// 堆栈信息
		fmt.Printf("stack trace:\n%+v\n", err)

		// 调用Is方法判断是否wrap错误sql.ErrNoRows
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("no rows return\n")
		}
	}
}

func queryHelper(d dao.DaoInterface) error {
	if err := d.Query("query something 1"); err != nil {
		return errors.Wrap(err, "fail to Query")
	}
	return nil
}
