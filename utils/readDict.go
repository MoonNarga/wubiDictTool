package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadDict(path string) [][]string {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Error", err)
		return nil
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	size := stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	dicts := make([][]string, 0)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		// fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return nil
			}
		}
		strs := strings.Split(line, "	")
		dicts = append(dicts, strs)
	}
	return dicts
}