package utils

import (
	"os"
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

	for k, v := range originMap {
		vNew, ok := newMap[k]
		if ok {
			if vNew[0] != v[0] {
				// fmt.Println(vNew[0], v[1])
				content = append(content, []byte(k)...)
				for _, vv := range v {
					content = append(content, '	')
					content = append(content, []byte(vv)...)
				}
				content = append(content, '	')
				content = append(content, []byte("->")...)
				for _, vv := range vNew {
					content = append(content, '	')
					content = append(content, []byte(vv)...)
				}
				content = append(content, '\n')
			}
		} else {
			content = append(content, []byte(k)...)
			content = append(content, []byte("	delete\n")...)
		}
	}

	err := os.WriteFile("./yamls/diff.yaml", content, 0644)
	if err != nil {
		panic(err)
	}
}
