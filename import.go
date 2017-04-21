package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/kiesel/golaid/dialog"
	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

var pathPrefix = "data"

func main() {
	spew.Config = spew.ConfigState{
		MaxDepth: 3,
		Indent:   "  ",
	}
	pages, err := importIndex()
	if err != nil {
		panic(err)
	}

	entries, err := importEntries(pages)
	if err != nil {
		panic(err)
	}

	fmt.Println(entries)
}

func importIndex() ([]dialog.Page, error) {
	fmt.Println("Importing indices:")
	pages := make([]dialog.Page, 0)
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

		array := data.(php_serialize.PhpArray)
		page, err := dialog.NewPage(array)
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)
	}

	return pages, nil
}

func importFile(fname string) (php_serialize.PhpValue, error) {
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

func importEntries(pages []dialog.Page) ([]php_serialize.PhpValue, error) {
	entries := make([]php_serialize.PhpValue, 1)

	for _, page := range pages {
		for _, ref := range page.Entries {
			entry, err := importEntry(ref)
			if err != nil {
				return nil, err
			}

			entries = append(entries, entry)
		}
	}

	return entries, nil
}

func importEntry(er dialog.EntryRef) (php_serialize.PhpValue, error) {
	fname := fmt.Sprintf("%s/%s", pathPrefix, er.Filename)
	fmt.Printf("Importing entry '%s'\n", fname)

	data, err := importFile(fname)
	if err != nil {
		return nil, err
	}

	ptr := data.(*php_serialize.PhpObject)
	return dialog.NewEntry(ptr)
}
