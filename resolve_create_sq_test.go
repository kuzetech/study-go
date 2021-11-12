package main

import (
	"reflect"
	"testing"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

type Column struct {
	Name               string
	Type               string
	Default_Expression string
	Compression_Codec  string
	Comment            string
}

func Test_Resolve_Create_Sql(t *testing.T) {
	dsn := "tcp://localhost:9000?database=default&username=default&read_timeout=10&write_timeout=20"
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

	if err != nil {
		t.Fatal(err)
	}

	tx := db.Raw("SELECT name, type, default_expression, compression_codec, comment FROM system.columns where database='default' and table='codec_example';")

	if tx.Error != nil {
		t.Fatal(tx.Error)
	}

	result := []Column{}
	tx.Scan(&result)

	t.Log(result)

	t.Log(reflect.DeepEqual(result[0], result[1]))

	Column1 := Column{Name: "1"}
	Column2 := Column{Name: "1"}

	t.Log(reflect.DeepEqual(Column1, Column2))

}
