package main
import (
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	"regexp"
	"strings"
)
var helpFlag = flag.Bool("h", false, "useage: [-d] <dir path...> <grep file> ")
var deepFlag = flag.Bool("d", false, "find in child dir")

func listFiles(dirPath, grepPattern string) {
	fi, err := os.Stat(dirPath)
	if err != nil || !fi.IsDir() {
		//fmt.Printf("1: %x\n", err)
		return;
	}
	fis, err :=ioutil.ReadDir(dirPath)
	if err != nil || len(fis) == 0 {
		//fmt.Printf("2: %x\n", err)
		return
	}
	grepReg := regexp.MustCompile(grepPattern)
	//grepReg, err := regexp.Compile(grepPattern)
	if err != nil {
		//fmt.Printf("3: %x\n", err)
		return
	}
	for _, fi := range fis {
		if fi.IsDir() {
			if(*deepFlag) {
				listFiles(dirPath + "/" + fi.Name(), grepPattern)
			}
			continue
		}
		if grepReg.MatchString(fi.Name()) {
			fmt.Println(dirPath + "/" + fi.Name())
		}
	}
}

func main() {
	flag.Parse()
	if *helpFlag{
		fmt.Println(flag.Lookup("h").Usage)
		return
	}
	var args []string
	if flag.NArg() == 0 {
		args = []string{".", "*"}
	} else if flag.NArg() == 1 {
		args = []string{".", flag.Args()[0]}
	} else {
		args = flag.Args()
	}
	for i, s := range args {
		if !strings.Contains(s, "-") {
			args = args[i:len(args)]
			break;
		}
	}
	dirs := args[0:len(args)-1]
	grepPattern := args[len(args)-1]
	grepPattern = strings.Replace(grepPattern, ".", "\\.", 256)
	grepPattern = "^" + strings.Replace(grepPattern, "*", ".*", 256) + "$"
	for _, dirPath := range dirs {
		listFiles(dirPath, grepPattern)
	}
}