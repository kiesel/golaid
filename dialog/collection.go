package dialog

import "github.com/yvasiyarov/php_session_decoder/php_serialize"

// EntryCollection is a collection of entries
type EntryCollection struct {
	*Entry
}

func newEntryCollection(in *php_serialize.PhpObject) (EntryCollection, error) {
	return EntryCollection{}, nil
}
