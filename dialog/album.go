package dialog

import (
	"time"

	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

// Album represents Dialog Album
type Album struct {
	*Entry
	Highlights []Image
	Chapters   []Chapter
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
