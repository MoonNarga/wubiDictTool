package utils

import (
	"os"
	"sort"
)

func GetDiff(dicts [][]string) {
	content := []byte{}

	originDict := ReadDict("./yamls/origin.yaml")
	SortDict(originDict)
	originMap := map[string][]string{}
	for _, v := range originDict {
		originMap[v[0]] = append(originMap[v[0]], v[1])
	}
	newMap := map[string][]string{}
	for _, v := range dicts {
		newMap[v[0]] = append(newMap[v[0]], v[1])
	}

	diffDict := make([][]string, 0)
	for k, v := range originMap {
		vNew, ok := newMap[k]
		str := []string{}
		if ok {
			if vNew[0] != v[0] {
				str = append(str, k)
				str = append(str, v...)
				str = append(str, "->")
				str = append(str, vNew...)
			}
			delete(newMap, k)
		} else {
			str = append(str, k, "delete")
		}
		if len(str) != 0 {
			diffDict = append(diffDict, str)
		}
	}
	for k, _ := range newMap {
		diffDict = append(diffDict, []string{k, "add"})
	}

	sort.Slice(diffDict, func(i, j int) bool {
		if len(diffDict[i]) > 2 && len(diffDict[j]) > 2 {
			return diffDict[i][len(diffDict[i])-1] < diffDict[j][len(diffDict[j])-1]
		}
		if len(diffDict[i]) == len(diffDict[j]) {
			if diffDict[i][1] == diffDict[j][1] {
				return diffDict[i][0] < diffDict[j][0]
			} else {
				return diffDict[i][1] < diffDict[j][1]
			}
		}
		return len(diffDict[i]) > len(diffDict[j])
	})
	for _, v := range diffDict {
		content = append(content, []byte(v[0])...)
		for i := 1; i < len(v); i++ {
			content = append(content, '	')
			content = append(content, []byte(v[i])...)
		}
		content = append(content, '\n')
	}

	err := os.WriteFile("./yamls/diff.yaml", content, 0644)
	if err != nil {
		panic(err)
	}
}
