package main

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/spf13/viper"
)

type Hash struct {
	Id         int64
	Encodedval string
}

var globalDb *pg.DB = nil

func FetchHash(id int64) (string, error) {
	hash := Hash{Id: id}
	err := globalDb.Model(&hash).WherePK().Select()
	return hash.Encodedval, err
}

func SaveHash(encodedString string) (int64, error) {
	hash := Hash{Encodedval: encodedString}
	_, err := globalDb.Model(&hash).Returning("id").Insert()
	return hash.Id, err
}

func CloseDb() {
	log.Println("Closing db connection ...")
	globalDb.Close()
}

// Connecting to db
func InitDb() *pg.DB {
	//log.Println("Propertty value " + viper.GetString("db.user"))
	opts := &pg.Options{
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Addr:     viper.GetString("db.address"),
		Database: viper.GetString("db.database"),
	}
	var db *pg.DB = pg.Connect(opts)
	//db = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to database XXXXXXXXXXXXXXXX")
	} else {
		log.Printf("Connected to postgres with user " + viper.GetString("db.user"))
	}

	globalDb = db
	//create table
	db.Model(&Hash{}).CreateTable(&orm.CreateTableOptions{})
	return db
}
