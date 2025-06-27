package go_regions

import (
	"embed"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 参考 https://github.com/wecatch/china_regions json数据
// 最新数据为2020的数据

//go:embed db/*
var embeddedRegionDB embed.FS

var _db *gorm.DB

func newDB() *gorm.DB {
	return _db
}

func init() {
	// 读取嵌入的 SQLite 数据库内容
	dbBytes, err := embeddedRegionDB.ReadFile("db/go_regions.db")
	if err != nil {
		log.Fatal("failed to read embedded db:", err)
	}

	// 创建临时文件并写入内容
	tmpFile, err := os.CreateTemp("", "go_regions.db")
	if err != nil {
		log.Fatal("failed to create temp file:", err)
	}
	defer os.Remove(tmpFile.Name()) // 退出时删除（可选）

	if _, err := tmpFile.Write(dbBytes); err != nil {
		log.Fatal("failed to write db to temp file:", err)
	}
	tmpFile.Close()

	// 用 GORM 打开这个临时文件
	_db, err = gorm.Open(sqlite.Open(tmpFile.Name()), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to open db with gorm:", err)
	}
}

type region struct {
	ID    int64  `json:"id" gorm:"column:id"`
	Name  string `json:"name" gorm:"column:name"`
	Level int    `json:"level" gorm:"column:level"` // 0-省 1-市 2-区 3-街道
}

// RegionList 地区列表
// pid为0时，获取的是省份列表
func RegionList(pid int) []region {
	var records []region
	newDB().Table("regions").Where("pid = ?", pid).Find(&records)
	return records
}

// RegionName 获取地区名称
func RegionName(id int) string {
	var record region
	newDB().Table("regions").Where("id = ?", id).Take(&record)
	return record.Name
}

func RegionInfo(id int) *region {
	var record region
	newDB().Table("regions").Where("id = ?", id).Take(&record)
	if record.ID == 0 {
		return nil
	}
	return &record
}
