package persistence

import (
	"fmt"
	"github.com/kailash-bhanushali/backend-golang/pkg/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testStruct api.TestStruct
var testStruct1 api.TestStruct

func DBConnect() api.TestStruct {
	//dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5433 sslmode=disable"
	dsn := "host=hippo-primary.postgres-operator.svc.cluster.local user=hippo password=password " +
		"dbname=hippo port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Getting Error while Connecting to the DB")
	}
	fmt.Println("DB CONN OPEN:", db)
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Getting Error while extracting DB from db conn")
	}
	fmt.Println("DB CONN OPEN:", sqlDB.Ping())
	defer sqlDB.Close()
	schema := "test"
	sqlDB.Exec(fmt.Sprintf("SET search_path TO %s, public", schema))
	sqlDB.QueryRow("SELECT * FROM account").Scan(&testStruct)
	//db.Raw("SELECT * FROM account").Scan(&testStruct)
	db.Raw("SELECT * FROM test.account").Scan(&testStruct1)
	//db.First(&testStruct)
	fmt.Println("Printing Test Account Result Data", testStruct)

	return testStruct
}
