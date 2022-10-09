package utils

import (
	"os"
	"sort"
	"strconv"
	"strings"
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

func SortDictWithKeyword(dicts [][]string, key string) {
	sort.Slice(dicts, func(i, j int) bool {
		if strings.Contains(dicts[i][0], key) && strings.Contains(dicts[j][0], key) {
			if dicts[i][0] == dicts[j][0] {
				return dicts[i][1] < dicts[j][1]
			}
			return dicts[i][0] < dicts[j][0]
		}
		if strings.Contains(dicts[i][0], key) {
			return true
		}
		if strings.Contains(dicts[j][0], key) {
			return false
		}
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

func OutputNewDict(dicts [][]string, path string) {
	content := []byte{}
	for _, v := range dicts {
		for _, vv := range v {
			content = append(content, []byte(vv)...)
			content = append(content, '	')
		}
		content[len(content)-1] = '\n'
	}
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		panic(err)
	}
}
