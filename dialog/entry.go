package dialog

import (
	"errors"
	"time"

	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

// Entry represents a base entry
type Entry struct {
	Name        string
	Title       string
	Description string
	CreatedAt   time.Time
}

// IEntry is the entries interface
type IEntry interface {
	GetName() string
	// GetTitle() string
	// GetDescription() string
	// GetCreatedAt() time.Time
}

// GetName retrieves the name of the entry
func (e *Entry) GetName() string {
	return e.Name
}

// NewEntry constructs a new dialog entry structure
func NewEntry(in *php_serialize.PhpObject) (IEntry, error) {
	switch in.GetClassName() {
	case "Album", "de.thekid.dialog.Album":
		return newAlbum(in)

	case "EntryCollection", "de.thekid.dialog.EntryCollection":
		return newEntryCollection(in)

	case "SingleShot", "de.thekid.dialog.SingleShot":
		return newSingleShot(in)

	default:
		return nil, errors.New("Cannot convert class " + in.GetClassName())
	}
}

func phpArrayFrom(in php_serialize.PhpValue, ok bool) *php_serialize.PhpArray {
	if !ok {
		return nil
	}

	out := in.(php_serialize.PhpArray)
	return &out
}

func getFieldString(p *php_serialize.PhpObject, field string) string {
	val, ok := p.GetPublic(field)
	if !ok {
		return ""
	}

	return php_serialize.PhpValueString(val)
}

func getFieldInt64(p *php_serialize.PhpObject, field string) int64 {
	val, ok := p.GetPublic(field)
	if !ok {
		return 0
	}

	return php_serialize.PhpValueInt64(val)
}
