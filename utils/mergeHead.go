package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func MergeHead(dicts [][]string) {
	content := make([]byte, 0)
	file, err := os.OpenFile("./yamls/head.yaml", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	size := stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
		content = append(content, []byte(line)...)
	}

	for _, v := range dicts {
		for _, vv := range v {
			content = append(content, []byte(vv)...)
			content = append(content, '	')
		}
		content[len(content)-1] = '\n'
	}
	err = os.WriteFile("./yamls/res.yaml", content, 0644)
	if err != nil {
		panic(err)
	}
}
