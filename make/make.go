package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//go:embed data/*
var embeddedRegionData embed.FS

func main() {
	os.Remove("db/go_regions.db")
	db, err := gorm.Open(sqlite.Open("db/go_regions.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect go_regions database:", err)
	}
	db = db.Debug()
	initSql, err := embeddedRegionData.ReadFile("data/init.sql")
	if err != nil {
		log.Fatal("failed to read go_regions init_sql:", err)
	}

	err = db.Exec(string(initSql)).Error
	if err != nil {
		panic("faile to init go_regions province sql")
	}

	regionsText, err := embeddedRegionData.ReadFile("data/regions.txt")
	if err != nil {

	}
	regionsTextRows := strings.Split(string(regionsText), "\n")
	for index, row := range regionsTextRows {
		if index < 1 {
			continue
		}
		cols := strings.Split(row, ",")
		// 不满足六位长度的补0
		if len(cols[0]) < 6 {
			cols[0] = cols[0] + strings.Repeat("0", 6-len(cols[0]))
		}
		if len(cols[1]) < 6 && cols[1] != "0" {
			cols[1] = cols[1] + strings.Repeat("0", 6-len(cols[1]))
		}
		insertSql := fmt.Sprintf("INSERT INTO regions ('id', 'pid', 'level','pinyin_prefix','pinyin','name') VALUES (%s,%s,%s,%s,%s,%s)",
			cols[0], cols[1], cols[2], cols[4], cols[5], cols[7])
		db.Exec(insertSql)
	}
}
