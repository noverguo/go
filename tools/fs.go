package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"bufio"
	"strings"
)

var suffixes []string = []string{"c++", "java", "c", "go", "py", "s", "cpp", "h", "js", "html", "jsp", "smali", "xml", "txt", "css"}

func findString(fileName string, str string) {
	stat, err := os.Stat(fileName)
	if err != nil {
		return
	}
	if stat.IsDir() {
		fileList ,err := ioutil.ReadDir(fileName)

		if err != nil {
			fmt.Println("read dir error")
			return
		}
		for _, name := range fileList {
			//fmt.Printf("find: %s\n", fileName + "\\" + name.Name())
			findString(fileName + "/" + name.Name(), str)
		}
		return;
	}
	txtFile := false
	tmpName := strings.ToLower(fileName)
	for _, suffix := range suffixes {
		if strings.HasSuffix(tmpName, suffix) {
			txtFile = true
			break
		}
	}
	if !txtFile {
		return
	}
	file, err := os.Open(fileName);
	defer file.Close()
	if err != nil {
		return;
	}
	reader := bufio.NewReader(file);
	i := 0
	printFile := false;
	for {
		currentLine, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		i++
		if strings.Contains(currentLine, str) {
			if !printFile {
				printFile = true
				fmt.Println(fileName)
			}
			fmt.Printf("\t%5d%s", i, currentLine)
		}
	}
	if printFile {
		fmt.Println()
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: <dir> <find string> [subfix...]")
		os.Exit(0)
	}
	dirName := os.Args[1]
	str := os.Args[2]
	if len(os.Args) > 3 {
		suffixes = os.Args[3:]
	}
	dirStat, err := os.Stat(dirName)
	if err != nil || !dirStat.IsDir() {
		fmt.Printf("%s not a dir\n", dirName)
		os.Exit(-1)
	}
	findString(dirName, str);
}
