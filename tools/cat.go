package main

import (
	"io"
	"os"
	"fmt"
	"bufio"
	"flag"
)
var numberFlag = flag.Bool("n", false, "number each line")
var helpFlag = flag.Bool("h", false, "useage: [-n] <file path...> \n\tn: show the numberline")

func cat(r *bufio.Reader) {
	i := 1
	for {
		buf, e := r.ReadBytes('\n')
		if e == io.EOF {
			break
		}
		if *numberFlag {
			fmt.Fprintf(os.Stdout, "%5d %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func main() {
	flag.Parse()
	if *helpFlag || flag.NArg() == 0 {
		fmt.Println(flag.Lookup("h").Usage)
		return
	}
	for _, fileName := range flag.Args() {
		fd, err := os.Open(fileName)
		defer fd.Close()
		if err != nil {
			fmt.Printf("%s not exist!", fileName)
			continue
		}
		cat(bufio.NewReader(fd))
	}
}