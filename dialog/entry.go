package dialog

import (
	"errors"
	"fmt"
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

// Album represents Dialog Album
type Album struct {
	*Entry
	Highlights interface{}
	Chapters   interface{}
}

// Image represents Dialog Image
type Image struct {
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

	default:
		return nil, errors.New("Cannot convert class " + in.GetClassName())
	}
}

func newAlbum(in *php_serialize.PhpObject) (*Album, error) {
	highlights, err := newHighlights(phpArrayFrom(in.GetPublic("highlights")))
	if err != nil {
		return nil, err
	}

	return &Album{
		Entry: &Entry{
			Name:        getFieldString(in, "name"),
			Title:       getFieldString(in, "title"),
			Description: getFieldString(in, "description"),
			CreatedAt:   time.Unix(getFieldInt64(in, "createdAt"), 0),
		},
		Highlights: highlights,
		Chapters:   nil,
	}, nil
}

func phpArrayFrom(in php_serialize.PhpValue, ok bool) *php_serialize.PhpArray {
	if !ok {
		return nil
	}

	out := in.(php_serialize.PhpArray)
	return &out
}

func newHighlights(in *php_serialize.PhpArray) ([]Image, error) {
	if in == nil {
		return []Image{}, nil
	}

	out := make([]Image, len(*in))

	for _, item := range *in {
		image, err := newImage(item.(*php_serialize.PhpObject))
		if err != nil {
			return nil, err
		}

		out = append(out, *image)
	}

	return out, nil
}

func newImage(in *php_serialize.PhpObject) (*Image, error) {
	fmt.Println(in)
	return nil, nil
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
