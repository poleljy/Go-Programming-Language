package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func Tutorial_pq() {
	//db, err := sql.Open("postgres", "host=192.168.1.189 port=5432 user=postgres password=123456 dbname=TitanSBO sslmode=disable")
	db, err := sql.Open("postgres", "postgresql://postgres:123456@192.168.1.189:5432/TitanSBO?sslmode=disable")
	checkErr(err)

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	// 插入数据
	stmt, err := db.Prepare("insert into groups(group_name,address,country,join_time) values($1,$2,$3,current_timestamp) RETURNING group_id")
	checkErr(err)
	res, err := stmt.Exec("AeroSpace Titan", "研发部门", "China")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}
