package tests

import (
	"favorite/ent/enttest"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDBConnect(t *testing.T) {

	dbDms := "root:abc123@tcp(127.0.0.1:3306)/favorite?charset=utf8&parseTime=True" //db.GetClientDms()
	log.Println("TestDB dbDms:" + dbDms)

	client := enttest.Open(t, "mysql", dbDms)
	defer client.Close()
}
