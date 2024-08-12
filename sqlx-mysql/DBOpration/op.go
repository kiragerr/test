package DBOpration

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var (
	now  = time.Now()
	date = now.Format("2006-01-02")
)

type TodoItem struct {
	Id       int    `db:"id"`
	Title    string `db:"title"`
	Content  string `db:"content"`
	Editdate string `db:"editdate"`
}

func AddTodo(db *sqlx.DB) (err error) {
	var (
		title   string
		content string
	)
	fmt.Println("输入标题:")
	fmt.Scanln(&title)
	fmt.Println("输入内容:")
	fmt.Scanln(&content)

	// 插入到表
	res, err := db.Exec("INSERT INTO todolist (title, content, editdate) VALUES(?,?,?)", title, content, date)
	if err != nil {
		return err
	}

	// 获取插入的行的id
	nID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	fmt.Println("添加成功,ID:", nID)
	return nil
}

func DeleteTodo(db *sqlx.DB) error {
	var id int

	fmt.Println("输入待删除的ID:")
	fmt.Scanln(&id)

	_, err := db.Exec("DELETE FROM todolist WHERE id =?", id)
	if err != nil {
		fmt.Println("删除失败")
		return err
	}

	fmt.Println("删除成功")
	return nil
}

func ViewTodo(db *sqlx.DB) error {
	var (
		data []TodoItem
	)

	err := db.Select(&data, "SELECT * FROM todolist")
	if err != nil {
		fmt.Println("查询失败")
		return err
	}

	for _, item := range data {
		fmt.Println("ID:", item.Id, "标题:", item.Title, "内容:", item.Content, "编辑日期:", item.Editdate)
	}
	fmt.Println("查询完毕")

	return nil
}

func ModifyTodo(db *sqlx.DB) error {

	var (
		id      int
		title   string
		content string
	)

	fmt.Println("输入待修改的ID:")
	fmt.Scanln(&id)

	fmt.Println("输入新标题")
	fmt.Scanln(&title)

	fmt.Println("输入新内容")
	fmt.Scanln(&content)

	_, err := db.Exec("UPDATE todolist SET title = ?, content = ?, editdate = ? WHERE id = ?", title, content, date, id)
	if err != nil {
		fmt.Println("修改失败")
		return err
	}

	fmt.Println("修改成功")
	return nil
}
