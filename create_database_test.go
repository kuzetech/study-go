package main

import (
	"testing"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

func Test_Create_Database(t *testing.T) {
	dsn := "tcp://localhost:9000?database=default&username=default&read_timeout=10&write_timeout=20"
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

	if err != nil {
		t.Fatal(err)
	}

	tx := db.Exec("CREATE DATABASE IF NOT EXISTS test")

	if tx.Error != nil {
		t.Fatal(tx.Error)
	}

}
