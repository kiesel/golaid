package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/juju/loggo"
	"github.com/kiesel/golaid/dialog"
	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

var pathPrefix = "data"
var file string
var logger loggo.Logger

func init() {
	flag.StringVar(&pathPrefix, "path", "data", "Path to the files")
	flag.StringVar(&file, "filename", "", "Specific file to import")
}

func main() {
	flag.Parse()

	loggo.ConfigureLoggers("<root> = DEBUG")
	logger = loggo.GetLogger("")

	spew.Config = spew.ConfigState{
		MaxDepth: 3,
		Indent:   "  ",
	}

	if file != "" {
		logger.Infof("Importing entries from '%s'", file)
		data, err := importEntryFromFile(file)
		if err != nil {
			panic(err)
		}

		spew.Dump(data)
	} else {
		logger.Infof("Importing default index ...")
		pages, err := importIndex()
		if err != nil {
			panic(err)
		}

		entries, err := importEntries(pages)
		if err != nil {
			panic(err)
		}

		spew.Dump(entries)
	}
}

func importIndex() ([]dialog.Page, error) {
	logger.Infof("Importing indices:")
	pages := make([]dialog.Page, 0)
	end := false

	for i := 0; !end; i++ {
		fname := fmt.Sprintf("%s/page_%d.idx", pathPrefix, i)
		logger.Infof("Importing %s\n", fname)
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

func importEntries(pages []dialog.Page) ([]dialog.IEntry, error) {
	entries := make([]dialog.IEntry, pages[0].Total)

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

func importEntry(er dialog.EntryRef) (dialog.IEntry, error) {
	fname := fmt.Sprintf("%s/%s", pathPrefix, er.Filename)
	logger.Infof("Importing entry '%s'\n", fname)

	return importEntryFromFile(fname)
}

func importEntryFromFile(fname string) (dialog.IEntry, error) {

	data, err := importFile(fname)
	if err != nil {
		return nil, err
	}

	ptr := data.(*php_serialize.PhpObject)
	return dialog.NewEntry(ptr)
}
