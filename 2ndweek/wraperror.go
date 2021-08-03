package main

import (
	"fmt"
	"wraperror/dao"

	"github.com/pkg/errors"
)

var (
	ErrNoUserFound = errors.New("no user found")
)

func main() {
	dao := dao.NewDao()
	if err := queryUser(dao); err != nil {
		// 查询失败，可能是包内代码的错误，也可能第三方包的错误
		fmt.Printf("fail to query user: %v\n\n", err)
		// 根因，这里的根因是业务错误的根因
		// fmt.Printf("original error: %T %v\n\n", errors.Cause(err), errors.Cause(err))
		// 堆栈信息
		fmt.Printf("stack trace:\n%+v\n", err)

		// 调用Is方法判断是否wrap错误sql.ErrNoRows
		if errors.Is(err, ErrNoUserFound) {
			fmt.Printf("no user found\n")
		}
	}
}

//这里是具体的业务，返回wrap具体的错误
func queryUser(d dao.DaoInterface) error {
	if err := d.Query("SQL"); err != nil {
		// 记录日志，dao层错误
		fmt.Printf("query error: %v\n\n", err)
		// 堆栈信息
		fmt.Printf("stack trace:\n%+v\n", err)
		// 返回业务错误
		return errors.Wrap(ErrNoUserFound, "fail query")
	}
	return nil
}
