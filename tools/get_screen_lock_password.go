package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/fighterlyt/permutation"
	"github.com/noverguo/util"
	"io/ioutil"
	"os"
	"strings"
)
// 破解图案解锁的密码
func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: <gesture.key path>\ngesture.key: you can found it in /data/system/gesture.key from the android device.")
		os.Exit(0)
	}
	gestureKeyPath := os.Args[1]
	fd, err := os.Open(gestureKeyPath)
	util.CheckErr(err)
	data, err := ioutil.ReadAll(fd)
	util.CheckErr(err)
	originKey := hex.EncodeToString(data)
	//fmt.Println(originKey)
	fmt.Println(strings.Join(find(originKey), " "))
}

func sum(val string) string {
	data, err := hex.DecodeString(val)
	util.CheckErr(err)
	h := sha1.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func find(originKey string) []string {
	arr := []string{"00", "01", "02", "03", "04", "05", "06", "07", "08"}
	p, err := permutation.NewPerm(arr, nil) //generate a Permutator
	util.CheckErr(err)
	for i, err := p.Next(); err == nil; i, err = p.Next() {
		vals := i.([]string)
		for i := 4; i <= len(arr); i++ {
			info := strings.Join(vals[0:i], "")
			//fmt.Println(info)
			tmpKey := sum(info)
			//fmt.Println(vals, tmpKey)
			if tmpKey == originKey {
				return vals[0:i]
			}
		}
	}

	return nil
}
