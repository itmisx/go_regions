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

type provinceT struct {
	Name       string `json:"name" gorm:"column:name"`
	ProvinceID string `json:"id" gorm:"column:province_id"`
}

// ProvinceList 省份列表
func ProvinceList(provinceName string) []provinceT {
	var records []provinceT
	newDB().Table("province").Find(&records)
	return records
}

type cityT struct {
	Name   string `json:"name" gorm:"column:name"`
	CityID string `json:"id" gorm:"column:city_id"`
}

// CityList 城市列表
func CityList(provinceID string) []cityT {
	var records []cityT
	tx := newDB().Table("city")
	if provinceID != "" {
		tx = tx.Where("province_id = ?", provinceID)
	}
	tx.Find(&records)
	return records
}

type countyT struct {
	Name     string `json:"name" gorm:"column:name"`
	CountyID string `json:"id" gorm:"column:county_id"`
}

// CountyList 县区列表
func CountyList(cityID string) []countyT {
	var records []countyT
	newDB().Table("county").
		Where("city_id = ?", cityID).
		Find(&records)
	return records
}

type townT struct {
	Name   string `json:"name" gorm:"column:name"`
	TownID string `json:"id" gorm:"column:town_id"`
}

// TownList 乡镇街道
func TownList(countyID string) []townT {
	var records []townT
	newDB().Table("town").
		Where("county_id = ?", countyID).
		Find(&records)
	return records
}

// 获取id名称
func ProvinceName(id string) string {
	var record provinceT
	newDB().Table("province").
		Where("province_id = ?", id).
		Take(&record)
	return record.Name
}
func CityName(id string) string {
	var record cityT
	newDB().Table("city").
		Where("city_id = ?", id).
		Take(&record)
	return record.Name
}
func CountyName(id string) string {
	var record countyT
	newDB().Table("county").
		Where("county_id = ?", id).
		Take(&record)
	return record.Name
}
func TownName(id string) string {
	var record townT
	newDB().Table("town").
		Where("town_id = ?", id).
		Take(&record)
	return record.Name
}
