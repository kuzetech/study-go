package main

import (
	"fmt"
	"testing"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

func Test_Exist_Database(t *testing.T) {
	dsn := "tcp://localhost:9000?database=default&username=default&read_timeout=10&write_timeout=20"
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

	if err != nil {
		t.Fatal(err)
	}

	data := db.Raw("EXISTS DATABASE defaultaaaa;")

	if data.Error != nil {
		t.Fatal(data.Error)
	}

	result := 0
	data.Scan(&result)

	fmt.Println(result)

}
