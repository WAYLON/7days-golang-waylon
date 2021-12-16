package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"waylonorm"
)

type User struct {
	Name string `waylonorm:"PRIMARY KEY"`
	Age  int
}

var (
	user1 = &User{"Tom", 181}
	user2 = &User{"Sam", 251}
	user3 = &User{"Jack", 251}
)

func main() {
	engine, _ := waylonorm.NewEngine("sqlite3", "../waylon.db")
	s := engine.NewSession()
	s.Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("failed init test records")
	}
}
