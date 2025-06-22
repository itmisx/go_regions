package go_regions

import (
	"fmt"
	"testing"
)

func TestXxx(*testing.T) {
	provinceList := ProvinceList("")
	cityList := CityList("410000000000")
	countyList := CountyList("410400000000")
	townList := TownList("410423000000")
	provinceName := ProvinceName("410000000000")
	cityName := CityName("410400000000")
	countyName := CountyName("410423000000")
	townName := TownName("410423202000")
	fmt.Println("provinceList", provinceList)
	fmt.Println("cityList", cityList)
	fmt.Println("countyList", countyList)
	fmt.Println("townList", townList)
	fmt.Println("provinceName", provinceName)
	fmt.Println("cityName", cityName)
	fmt.Println("countyName", countyName)
	fmt.Println("townName", townName)
}
