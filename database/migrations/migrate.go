package main

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-framework/app/models"
	"github.com/ArtisanCloud/go-framework/boostrap"
	"github.com/ArtisanCloud/go-framework/database"
	"os"
	"reflect"
)

func init() {
	boostrap.InitProject()
}

var (
	needRefresh bool
	isForce     bool
)

func main() {

	arrayArgs := os.Args[1:]
	needRefresh = object.InArray("refresh", arrayArgs)
	isForce = object.InArray("force", arrayArgs)

	arrayTables := []interface{}{
		&models.User{},
	}
	err := migrate(arrayTables)

	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}

	println("migrate done")

	return
}

func migrate(arrayTables []interface{}) error {
	for _, table := range arrayTables {
		tableName := reflect.TypeOf(table).String()
		hasTable := database.DBConnection.Migrator().HasTable(table)

		// force to drop tables if has this table
		if needRefresh && hasTable {
			err := database.DBConnection.Migrator().DropTable(table)
			if err != nil {
				println(err.Error())
				return err
			} else {
				hasTable = false
				fmt.Printf("drop table:%s \n", tableName)
			}
		}

		if !hasTable {
			err := database.DBConnection.AutoMigrate(table)
			if err != nil {
				println(err.Error())
				return err
			} else {
				fmt.Printf("create table:%s \n", tableName)
			}
		}

	}

	return nil
}
