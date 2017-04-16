package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"strings"

	"github.com/wulijun/go-php-serialize/phpserialize"
)

var pathPrefix = "data"

func main() {
	index, err := importIndex()
	if err != nil {
		panic(err)
	}
	fmt.Println(index)

	entries, err := importEntries(index)
	if err != nil {
		panic(err)
	}

	fmt.Println(entries)
}

func importIndex() (map[string]string, error) {
	fmt.Println("Importing indices:")
	index := make(map[string]string)
	end := false

	for i := 0; !end; i++ {
		fname := fmt.Sprintf("%s/page_%d.idx", pathPrefix, i)
		fmt.Printf("Importing %s\n", fname)
		data, err := importFile(fname)

		if err != nil {
			if os.IsNotExist(err) {
				end = true
				continue
			}

			// Otherwise end it...
			return nil, err
		}

		for key, val := range data.(map[interface{}]interface{}) {
			if key == "entries" {
				for idx, name := range val.(map[interface{}]interface{}) {
					index[idx.(string)] = name.(string)
				}
			}
		}
	}

	return index, nil
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

	return phpserialize.Decode(string(buf))
}

func importEntries(index map[string]string) ([]interface{}, error) {
	entries := make([]interface{}, len(index))

	for key := range index {
		entry, err := importEntry(key)
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func importEntry(key string) (interface{}, error) {
	parts := strings.SplitN(key, "-", 2)
	fname := fmt.Sprintf("%s/%s", pathPrefix, parts[1])

	fmt.Printf("Importing entry '%s'\n", fname)

	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return phpserialize.Decode(string(buf))
}
