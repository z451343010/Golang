/*
	GoWeb
	Gin框架
	Gin 连接 MySQL
*/
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {

	connStr := "root:zZ18960313139@tcp(127.0.0.1:3306)/ginsql"

	// "mysql" 指定连接数据库的驱动程序
	// connStr 指定连接数据库的字符串
	database, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		fmt.Println("MySQL 连接成功")
		fmt.Println("db = ", database)
	}

	// 创建数据库表
	_, err = database.Exec("create table person(id int auto_increment primary key," +
		"name varchar(12) not null," +
		"age int default 1);")

	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		fmt.Println("数据库建表成功")
	}

	// 插入数据到数据库表
	_, err = database.Exec("insert into person(name, age)"+
		"values(?,?);", "zjune", 18)

	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		fmt.Println("插入数据到数据库表成功")
	}

	// 查询数据库
	rows, err := database.Query("select id,name,age from person")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	for rows.Next() {
		person := new(Person)
		err := rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(person.Id, person.Name, person.Age)
	}

}
