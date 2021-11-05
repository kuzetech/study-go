package main

import (
	"fmt"
	"testing"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

func Test_Show_Database(t *testing.T) {
	dsn := "tcp://localhost:9000?database=default&username=default&read_timeout=10&write_timeout=20"
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

	if err != nil {
		t.Fatal(err)
	}

	tx := db.Raw("show databases;")

	if tx.Error != nil {
		t.Fatal(tx.Error)
	}

	result := []string{}
	tx.Scan(&result)

	fmt.Println(result)

}
