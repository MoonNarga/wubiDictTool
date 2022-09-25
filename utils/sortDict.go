package utils

import (
	"sort"
	"strconv"
)

func SortDict(dicts [][]string) {
	sort.Slice(dicts, func(i, j int) bool {
		if dicts[i][1] == dicts[j][1] {
			if len(dicts[i]) == 3 {
				v1, _ := strconv.Atoi(dicts[i][2])
				v2, _ := strconv.Atoi(dicts[j][2])
				return v1 > v2
			}
			return dicts[i][0] < dicts[j][0]
		}
		return dicts[i][1] < dicts[j][1]
	})
}
