package dialog

import (
	"reflect"
	"testing"

	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/yvasiyarov/php_session_decoder/php_serialize"
)

func Test_NewObject_Album(t *testing.T) {
	actual, err := fromPHPToObject(Album{}, `O:5:"Album":7:{s:4:"name";s:9:"albumname";s:5:"title";s:13:"Album's Title";s:11:"description";s:23:"The album's description";s:9:"createdAt";O:4:"Date":2:{s:5:"value";s:24:"2012-05-28 00:00:00+0200";s:4:"__id";N;}s:4:"__id";N;s:10:"highlights";a:0:{}s:8:"chapters";a:0:{}}`)

	if err != nil {
		t.Error(err)
		return
	}

	var noHighlights []Image
	var noChapter []Chapter

	expected := Album{
		Entry: &Entry{
			Name:        "albumname",
			Title:       "Album's Title",
			Description: "The album's description",
			CreatedAt:   time.Unix(1338156000, 0),
		},
		Highlights: noHighlights,
		Chapters:   noChapter,
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("Actual value did not meet expected", spew.Sdump(actual), spew.Sdump(expected))
	}
}

func fromPHPToObject(inType interface{}, in string) (interface{}, error) {
	data, err := php_serialize.UnSerialize(in)
	if err != nil {
		return nil, err
	}

	albumData := data.(*php_serialize.PhpObject)
	return newObject(inType, albumData)
}
