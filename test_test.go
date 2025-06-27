package go_regions

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestXxx(*testing.T) {
	list := RegionList(0)
	js, _ := json.Marshal(list)
	fmt.Println("list", string(js))
	list = RegionList(int(list[0].ID))
	fmt.Println("list", list)
	list = RegionList(int(list[0].ID))
	fmt.Println("list", list)
	list = RegionList(int(list[0].ID))
	fmt.Println("list", list)
	regionName := RegionName(120000)
	fmt.Println("regionName", regionName)
	regionInfo := RegionInfo(120000)
	js, _ = json.Marshal(regionInfo)
	fmt.Println("regionInfo", string(js))
}
