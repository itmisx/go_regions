package main

import (
	"embed"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 参考 https://github.com/wecatch/china_regions json数据
// 最新数据为2020的数据

//go:embed sql/*
var embeddedRegionSQL embed.FS

func main() {
	os.Remove("db/go_regions.db")
	db, err := gorm.Open(sqlite.Open("db/go_regions.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect go_regions database:", err)
	}

	initSql, err := embeddedRegionSQL.ReadFile("sql/init.sql")
	if err != nil {
		log.Fatal("failed to read go_regions init_sql:", err)
	}
	err = db.Exec(string(initSql)).Error
	if err != nil {
		panic("faile to init go_regions init sql")
	}

	initSql, err = embeddedRegionSQL.ReadFile("sql/province.sql")
	if err != nil {
		log.Fatal("failed to read go_regions province sql:", err)
	}
	err = db.Exec(string(initSql)).Error
	if err != nil {
		panic("faile to init go_regions province sql")
	}

	initSql, err = embeddedRegionSQL.ReadFile("sql/city.sql")
	if err != nil {
		log.Fatal("failed to read go_regions city sql:", err)
	}
	err = db.Exec(string(initSql)).Error
	if err != nil {
		panic("faile to init go_regions city sql")
	}

	initSql, err = embeddedRegionSQL.ReadFile("sql/county.sql")
	if err != nil {
		log.Fatal("failed to read go_regions county sql:", err)
	}
	err = db.Exec(string(initSql)).Error
	if err != nil {
		panic("faile to init go_regions county sql")
	}

	initSql, err = embeddedRegionSQL.ReadFile("sql/town.sql")
	if err != nil {
		log.Fatal("failed to read go_region town sql:", err)
	}
	err = db.Exec(string(initSql)).Error
	if err != nil {
		panic("faile to init go_regions town sql")
	}
}
