package main

// 定义main 包
// 导入必须的几个包
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// github.com/go-sql-driver/mysql  这个包 使用go get 的方式安装到了本地

func main() {
	// 主程序
	db, err := sql.Open("mysql", "root:123456@tcp(10.10.203.69:6606)/test")
	// 打开连接  方法是 sql.Open 第一个参数是 数据库类型. 第二个是 用户名:密码@网络协议(ip:port)/需要查询的数据库名
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// 好像都有一个这样的处理.
	rows, _ := db.Query("select id,name from zhaobsh")
	//rows 查询 表里面所有的数据 结果应该是一个数组 方式db.Query
	id := 0
	name := ""
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
	//遍历数组里面的内容. 并且打印出来.  Scan 和 Next 的函数
	dbinsert, _ := db.Exec("insert into zhaobsh(name) values('zhaobsh01-')")
	//执行插入的数据, db.Exec 的函数
	fmt.Println(dbinsert)
	// 避免编译出错的处理.
}
