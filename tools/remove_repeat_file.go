package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/noverguo/util"
	"io"
	"io/ioutil"
	"os"
)
// 删除重复的文件
func main() {
	if len(os.Args) < 2 {
		return
	}
	dir := os.Args[1]
	fis, err := ioutil.ReadDir(dir)
	util.CheckErr(err)
	md5s := make(map[string]bool)
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}
		filePath := dir + "/" + fi.Name()
		fd, err := os.Open(filePath)
		if err != nil {
			continue
		}
		h := md5.New()
		io.Copy(h, fd)
		fd.Close()
		md5Str := hex.EncodeToString(h.Sum(nil))
		if _, exist := md5s[md5Str]; exist == false {
			md5s[md5Str] = true
		} else {
			fmt.Println("remove: ", filePath)
			err := os.Remove(filePath)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
