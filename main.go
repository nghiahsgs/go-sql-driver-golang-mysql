package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //khoi tao, su dung side effect
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func excuteStatement(db *sql.DB, sql string) int {
	insert, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer insert.Close()
	return 1
}
func main() {
	db, err := sql.Open("mysql", "nghiahsgs4:261997@tcp(165.22.58.219:3306)/test")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Test1 : INSERT TO DB
	// for i := 0; i < 1000000; i++ {
	// 	sql := "INSERT INTO `user`(`name`, `age`) VALUES ('nghiaz" + strconv.Itoa(i) + "',20)"
	// 	kq := excuteStatement(db, sql)
	// 	fmt.Println("kq", kq)
	// }

	// Test2 : SELECT MANY
	// sql := "SELECT * FROM user"
	// results, err := db.Query(sql)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var listUser []User
	// for results.Next() {
	// 	var user User

	// 	err := results.Scan(&user.ID, &user.Name, &user.Age)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	listUser = append(listUser, user)
	// }
	// fmt.Println(len(listUser))

	// Test3 : SELECT ONE
	var user User
	// db.QueryRow("SELECT * FROM user WHERE id =100").Scan(&user.ID, &user.Name, &user.Age)
	db.QueryRow("SELECT * FROM user WHERE id =?", 100).Scan(&user.ID, &user.Name, &user.Age)
	fmt.Println(user)
}
