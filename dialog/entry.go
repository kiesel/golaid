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

// EntryCollection is a collection of entries
type EntryCollection struct {
	*Entry
}

// SingleShot is a single shot
type SingleShot struct {
	*Entry
}

// Chapter represents a chapter
type Chapter struct {
}

// AlbumImage represents Dialog Image
type AlbumImage struct {
	Name     string
	Width    int64
	Height   int64
	ExifData ExifData
	IptcData IptcData
}

// ExifData represents exif data
type ExifData struct {
}

// IptcData represents iptc data
type IptcData struct {
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

func (a *Album) String() string {
	return fmt.Sprintf("{%T}", a)
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

func newAlbum(in *php_serialize.PhpObject) (*Album, error) {
	highlights, err := newHighlights(phpArrayFrom(in.GetPublic("highlights")))
	if err != nil {
		return nil, err
	}

	chapters, err := newChapter(phpArrayFrom(in.GetPublic("chapters")))
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
		Chapters:   chapters,
	}, nil
}

func newEntryCollection(in *php_serialize.PhpObject) (*EntryCollection, error) {
	fmt.Println(in)
	return nil, nil
}

func newSingleShot(in *php_serialize.PhpObject) (*SingleShot, error) {
	fmt.Println(in)
	return nil, nil
}

func phpArrayFrom(in php_serialize.PhpValue, ok bool) *php_serialize.PhpArray {
	if !ok {
		return nil
	}

	out := in.(php_serialize.PhpArray)
	return &out
}

func newHighlights(in *php_serialize.PhpArray) ([]AlbumImage, error) {
	if in == nil {
		return []AlbumImage{}, nil
	}

	out := make([]AlbumImage, len(*in))

	for _, item := range *in {
		image, err := newAlbumImage(item.(*php_serialize.PhpObject))
		if err != nil {
			return nil, err
		}

		out = append(out, *image)
	}

	return out, nil
}

func newAlbumImage(in *php_serialize.PhpObject) (*AlbumImage, error) {
	fmt.Println(in)
	return &AlbumImage{
		Name:   getFieldString(in, "name"),
		Width:  getFieldInt64(in, "width"),
		Height: getFieldInt64(in, "height"),
		// TODO ExifData: ...
		// TODO IptcData: ...
	}, nil
}

func newChapter(in *php_serialize.PhpArray) (*Chapter, error) {
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
