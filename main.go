package main

import (
	// "fmt"
	"wubiDictTool/utils"
)

func main() {
	dicts := utils.ReadDict("./yamls/dict.yaml")
	utils.SortDict(dicts)
	utils.MergeHead(dicts)
	utils.GetDiff(dicts)
}
