package utils

import (
	"sort"
	"strconv"
)

func SortDict(dicts [][]string) {
	sort.Slice(dicts, func(i, j int) bool {
		if dicts[i][1] == dicts[j][1] {
			if len(dicts[i]) != len(dicts[j]) {
				return len(dicts[i]) > len(dicts[j])
			}
			if len(dicts[i]) == 3 {
				v1, _ := strconv.Atoi(dicts[i][2])
				v2, _ := strconv.Atoi(dicts[j][2])
				return v1 > v2
			} else if len(dicts[i][0]) == len(dicts[j][0]) {
				return dicts[i][0] < dicts[j][0]
			}
			return len(dicts[i][0]) < len(dicts[j][0])
		}
		return dicts[i][1] < dicts[j][1]
	})
}
