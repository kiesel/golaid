package dialog

import (
	"strings"

	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

// Page represents a Dialog page
type Page struct {
	Total   int64
	Perpage int64
	Entries []EntryRef
}

// EntryRef represents a reference to an entry
type EntryRef struct {
	Utime    string
	Filename string
	Symbol   string
}

// NewPage creates a new Page instance
func NewPage(in php_serialize.PhpArray) (Page, error) {
	entries := make([]EntryRef, 0)

	for key, val := range in["entries"].(php_serialize.PhpArray) {
		entries = append(entries, NewEntryRef(php_serialize.PhpValueString(key), php_serialize.PhpValueString(val)))
	}

	return Page{
		Total:   php_serialize.PhpValueInt64(in["total"]),
		Perpage: php_serialize.PhpValueInt64(in["perpage"]),
		Entries: entries,
	}, nil
}

// NewEntryRef creates a new EntryRef instance
func NewEntryRef(key, name string) EntryRef {
	parts := strings.SplitN(key, "-", 2)
	return EntryRef{
		Utime:    parts[0],
		Filename: parts[1],
		Symbol:   name,
	}
}

func findKey(in php_serialize.PhpArray, needle php_serialize.PhpValue) php_serialize.PhpValue {
	out, ok := in[needle]
	if !ok {
		return nil
	}

	return out
}
