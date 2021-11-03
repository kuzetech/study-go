package main

import (
	"testing"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func Test_DBResolver(t *testing.T) {
	dsn := "tcp://localhost:9000?database=default&username=default&read_timeout=10&write_timeout=20"
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

	if err != nil {
		t.Fatal(err)
	}

	err = db.Use(
		dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{clickhouse.Open("tcp://localhost:9000?database=system&username=default&read_timeout=10&write_timeout=20")},
			Replicas: []gorm.Dialector{},
			Policy:   dbresolver.RandomPolicy{},
		}, "settings").Register(dbresolver.Config{
			Sources:  []gorm.Dialector{clickhouse.Open("tcp://localhost:9000?database=test&username=default&read_timeout=10&write_timeout=20")},
			Replicas: []gorm.Dialector{},
			Policy:   dbresolver.RandomPolicy{},
		}, "user_local"))

	if err != nil {
		t.Fatal(err)
	}

	result := map[string]interface{}{}
	db.Table("user_local").Take(&result)

	t.Log(result)
}
