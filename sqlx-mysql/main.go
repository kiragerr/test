package main

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 为sql包提供mysql驱动支持
	"github.com/jmoiron/sqlx"          // sql包的扩展包，需要引入sql包

	. "sqlx-mysql/DBOpration"
)

// 创建CLI程序，要求进入todo代办数据库中增删改查待办事项

func main() {
	var (
		db *sqlx.DB

		tmp int
	)

	// 连接数据库
	connet, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/todo")
	if err != nil {
		fmt.Println("连接数据库失败：", err)
		return
	}

	// 执行完毕后关闭数据库
	db = connet
	defer db.Close()

	// 处理错误
	fmt.Println("数据库连接成功")
	fmt.Println("输入操作类型: 1.添加待办事项 2.删除 3.查看 4.修改 5.退出")
	fmt.Scanln(&tmp)

	switch tmp {
	case 1:
		err = AddTodo(db)
	case 2:
		err = DeleteTodo(db)
	case 3:
		err = ViewTodo(db)
	case 4:
		err = ModifyTodo(db)
	case 5:
		fmt.Println("退出程序")
		return
	default:
		fmt.Println("输入有误，结束")
		return
	}

	if err != nil {
		fmt.Println("操作失败：", err)
		return
	}
}
