package main

import (
	"Week02/service"
	"database/sql"
	"errors"
	"fmt"
	pkgerror "github.com/pkg/errors"
)

func main() {
	_, err := service.Service()
	fmt.Println(err)
	if err != nil {
		// 在最顶层处理error，记录日志
		if errors.Is(pkgerror.Cause(err), sql.ErrNoRows) {
			// 打印堆栈
			fmt.Printf("根错误发生信息：%T %v\n", pkgerror.Cause(err), pkgerror.Cause(err))
			fmt.Printf("堆栈信息：\n%+v\n", err)
		}

		return
	}

	fmt.Println("返回结果或处理业务")
}
