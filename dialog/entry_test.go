package dialog

import (
	"testing"

	"time"

	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

func Test_NewEntry_Album(t *testing.T) {
	val := php_serialize.NewPhpObject("de.thekid.dialog.Album")
	val.SetPublic("name", "The name")
	val.SetPublic("title", "My Dialog Album")
	val.SetPublic("description", "This is the full whole description. With <html/>.")
	val.SetPublic("createdAt", int64(0))

	entry, err := NewEntry(val)

	if err != nil {
		t.Error(err)
		return
	}

	assertDeepEqual(&Album{
		Entry: &Entry{
			Name:        "The name",
			Title:       "My Dialog Album",
			Description: "This is the full whole description. With <html/>.",
			CreatedAt:   time.Unix(0, 0),
		},
		Highlights: nil,
		Chapters:   nil,
	}, entry, t)
}
