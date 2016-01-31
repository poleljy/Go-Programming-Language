// PostgreSQL驱动pg, https://github.com/lib/pq 支持database/sql驱动，纯Go写

/* 建表语句：

CREATE TABLE userinfo
(
    uid serial NOT NULL,
    username character varying(100) NOT NULL,
    departname character varying(500) NOT NULL,
    Created date,
    CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
)
WITH (OIDS=FALSE);

*/

```go
package main

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/lib/pq"
)

func main() {
	//db, err := sql.Open("postgres", "user=postgres password=postgres dbname=test sslmode=disable")
	
	db, err := sql.Open("postgres", "postgres://postgres:postgres@192.168.1.189/test?sslmode=disable")
	
	checkErr(err)
	
	// 插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)
	
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)
	
	//pg不支持这个函数，因为他没有类似MySQL的自增ID
	//id, err := res.LastInsertId()
	//checkErr(err)
	//fmt.Println(id)
	
	// 更新数据
	stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", 6)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("affect rows:", affect)
	
	// 查询数据
	
	rows, err := db.Query("SELECT * FROM userinfo order by uid")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid:",uid, "username:", username, "department:", department, "created", created)
		
		// 测试插入
		stmt, err = db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
		checkErr(err)
	
		res, err = stmt.Exec("pole", "研发部门", "2015-12-09")
		checkErr(err)
	}
	
	// 删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	res, err = stmt.Exec(6)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println("affect rows:", affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
```