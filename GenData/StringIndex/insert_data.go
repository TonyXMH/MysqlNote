package main

import (
	"fmt"
	"github.com/TonyXMH/MysqlNote/GenData/Client"
	"strconv"
	"time"
)

//为了逻辑的清晰去掉了所有的错误处理。
//为SUser表插入指定数量的数据
func genUserInfo(db, table string, num int) {
	nameList := []string{"chao", "qian", "sun", "li", "zhou", "wu", "zhen", "wang"}

	client, _ := Client.NewClient(db)
	defer client.Close()
	query := "insert into " + table + "(ID,email,name,passwd) values (?,?,?,?)"
	start := time.Now()
	tx, _ := client.Begin()
	stmt, _ := tx.Prepare(query)
	for i := 0; i < num; i++ {
		name := nameList[i%8]
		email := name + strconv.Itoa(i) + "@qq.com"
		passwd := "123456"
		_, _ = stmt.Exec(i, email, name, passwd)
	}
	_ = stmt.Close()
	_ = tx.Commit()
	end := time.Since(start)
	fmt.Println(end)
}

func main() {
	genUserInfo("ex11", "SUser", 100000)
}
