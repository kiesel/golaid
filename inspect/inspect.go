package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

var filename string

func init() {
	flag.StringVar(&filename, "filename", "", "Filename to inspect")
}

func main() {
	flag.Parse()

	if filename == "" {
		fmt.Println("Must give -filename param.")
		return
	}

	data, err := importFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}

func importFile(fname string) (interface{}, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return php_serialize.UnSerialize(string(buf))
}
