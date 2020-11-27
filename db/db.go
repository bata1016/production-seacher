package db

import (
	"fmt"

	"github.com/bata1016/production-seacher/models/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Initはデータベースの起動を行うファンクション
func Init() {
	GetGormConnect()
}

// var db gorm.DB

// GetGormConnect データベースへの接続を行う
func GetGormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	// PASS := "root"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "production_seacher"
	CONNECT := USER + ":" + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err)
	}
	// defer db.Close()
	// DBエンジンを「InnoDB」に設定
	// DBエンジンはDBがCRUDを行うための部品
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	// 登録するテーブル名を単数形にする
	db.SingularTable(true)

	// マイグレーション（テーブルがない時は自動生成）
	db.AutoMigrate(&entity.Employee{})
	db.AutoMigrate(&entity.Production{})

	fmt.Println("db connected: ", db)
	return db
}

// func autoMigration() {
// 	db.AutoMigrate(&entity.Employee{})
// }
