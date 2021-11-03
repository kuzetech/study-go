package main

import (
	"testing"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

func Test_Create_Table(t *testing.T) {
	dsn := "tcp://localhost:9000?database=default&username=default&read_timeout=10&write_timeout=20"
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

	if err != nil {
		t.Fatal(err)
	}

	tx := db.Exec(`CREATE TABLE test.user_local
	(
		uid         UInt32                      COMMENT '用户ID',
		time        Date                        COMMENT '事件时间戳'
	)
	ENGINE = ReplicatedReplacingMergeTree('/clickhouse/tables/{shard}/user_local', '{replica}')
	PARTITION BY (time)
	ORDER BY (uid);`)

	if tx.Error != nil {
		t.Fatal(tx.Error)
	}

}
