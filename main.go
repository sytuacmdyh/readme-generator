package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var fileContent = `
# 学习分享

------

`

func main() {
	dir := getCurDir()
	println(dir)
	listFile(dir, "", 0)

	println(fileContent)

	write("README.md", fileContent)
}

func getCurDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd + "/"
}

func listFile(path, mid string, dep int) {
	files, _ := ioutil.ReadDir(path + mid)
	pre := ""
	for t := 0; t < dep; t++ {
		pre += "  "
	}
	for _, file := range files {
		fn := file.Name()
		if fn[0] == '.' {
			continue
		}

		if strings.HasSuffix(fn, ".md") || strings.HasSuffix(fn, ".MD") || file.IsDir() {
			fileContent += pre + fmt.Sprintf("- [%s](%s)\n", fn, mid+fn)
		}

		if file.IsDir() && file.Name()[0] != '.' {
			listFile(path, mid+file.Name()+"/", dep+1)
		}
	}
}

func write(fileName, content string) {
	data := []byte(content)
	if ioutil.WriteFile(fileName, data, 0644) == nil {
		fmt.Println("写入文件成功:", content)
	}
}
