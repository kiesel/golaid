package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"strings"

	"github.com/kiesel/golaid/dialog"
	"github.com/wulijun/go-php-serialize/phpserialize"
)

var pathPrefix = "data"

func main() {
	pages, err := importIndex()
	if err != nil {
		panic(err)
	}
	fmt.Println(pages)

	entries, err := importEntries(pages)
	if err != nil {
		panic(err)
	}

	fmt.Println(entries)
}

func importIndex() ([]dialog.Page, error) {
	fmt.Println("Importing indices:")
	pages := make([]dialog.Page, 1)
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

		page, err := dialog.NewPage(data.(map[interface{}]interface{}))
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)
	}

	return pages, nil
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

func importEntries(pages []dialog.Page) ([]interface{}, error) {
	entries := make([]interface{}, 1)

	for _, page := range pages {
		for ref := range page.Entries {
			fmt.Println(ref)
		}
		// entry, err := importEntry(key)
		// if err != nil {
		// 	return nil, err
		// }

		// entries = append(entries, entry)
	}

	return entries, nil
}

func importEntry(key string) (interface{}, error) {
	parts := strings.SplitN(key, "-", 2)
	fname := fmt.Sprintf("%s/%s", pathPrefix, parts[1])

	fmt.Printf("Importing entry '%s'\n", fname)
	return importFile(fname)
}
