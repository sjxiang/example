package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:my-secret-pw@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=true&loc=Local"
)


func main() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot connect to db")
	}

	// 迁移 schema
	if err := db.Debug().AutoMigrate(&Shelf{}, &Book{}); err != nil {
		panic("failed to create table")
	}
	
	fmt.Println("over")
}

// Shelf 书架
type Shelf struct {
	ID       int64 `gorm:"primaryKey"`
	Theme    string
	Size     int64
	CreateAt time.Time
	UpdateAt time.Time
}

func (s Shelf) TableName() string {
	return "shelf"
}


// Book 图书
type Book struct {
	ID       int64 `gorm:"primaryKey"`
	Author   string
	Title    string
	ShelfID  int64
	CreateAt time.Time
	UpdateAt time.Time
}

func (b Book) TableName() string {
	return "book"
}
