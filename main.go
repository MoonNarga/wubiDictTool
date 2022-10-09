package main

import (
	"fmt"
	"os"
	"wubiDictTool/utils"
)

func main() {
	argc := len(os.Args)
	if argc > 2 {
		fmt.Println("too many args")
		return
	}
	dicts := utils.ReadDict("./yamls/dict.yaml")
	dictsOrigin := utils.ReadDict("./yamls/origin.yaml")
	switch argc {
	case 1:
		{
			utils.SortDict(dicts)
			utils.SortDict(dictsOrigin)
			utils.MergeHead(dicts)
		}
	case 2:
		{
			utils.SortDictWithKeyword(dicts, os.Args[1])
			utils.SortDictWithKeyword(dictsOrigin, os.Args[1])
		}
	}
	utils.OutputNewDict(dicts, "./yamls/dict.yaml")
	utils.OutputNewDict(dictsOrigin, "./yamls/origin.yaml")
	utils.GetDiff(dicts)
}
